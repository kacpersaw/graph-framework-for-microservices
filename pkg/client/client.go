package client

import (
	"api-gw/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/common-library.git/pkg/nexus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"strings"
)

var Client dynamic.Interface
var Host string

func New(config *rest.Config) (err error) {
	Host = config.Host
	Client, err = dynamic.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func CreateObject(gvr schema.GroupVersionResource, kind, hashedName string, labels map[string]string, body map[string]interface{}) error {
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": gvr.GroupVersion().String(),
			"kind":       kind,
			"metadata": map[string]interface{}{
				"name":   hashedName,
				"labels": labels,
			},
			"spec": body,
		},
	}

	// Create resource
	_, err := Client.Resource(gvr).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func GetObject(gvr schema.GroupVersionResource, hashedName string, opts metav1.GetOptions) (*unstructured.Unstructured, error) {
	obj, err := Client.Resource(gvr).Get(context.TODO(), hashedName, opts)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func DeleteObject(gvr schema.GroupVersionResource, crdType string, crdInfo model.NodeInfo, hashedName string) error {
	// Get object
	obj, err := Client.Resource(gvr).Get(context.TODO(), hashedName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	labels := obj.GetLabels()

	// Delete all children
	listOpts := metav1.ListOptions{LabelSelector: fmt.Sprintf("%s=%s", crdType, labels["nexus/display_name"])}
	for k, _ := range crdInfo.Children {
		err = DeleteChildren(k, listOpts)
		if err != nil {
			return err
		}
	}

	if len(crdInfo.ParentHierarchy) > 0 {
		parentCrdName := crdInfo.ParentHierarchy[len(crdInfo.ParentHierarchy)-1]
		parentCrdInfo := model.CrdTypeToNodeInfo[parentCrdName]
		err = UpdateParentWithRemovedChild(parentCrdName, parentCrdInfo, obj.GetLabels(), crdType, labels["nexus/display_name"])
		if err != nil {
			return err
		}
	}

	// Delete object
	err = Client.Resource(gvr).Delete(context.TODO(), hashedName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

func DeleteChildren(crdType string, listOpts metav1.ListOptions) error {
	crdInfo := model.CrdTypeToNodeInfo[crdType]
	for k, _ := range crdInfo.Children {
		err := DeleteChildren(k, listOpts)
		if err != nil {
			return err
		}
	}

	parts := strings.Split(crdType, ".")
	gvr := schema.GroupVersionResource{
		Group:    strings.Join(parts[1:], "."),
		Version:  "v1",
		Resource: parts[0],
	}
	err := Client.Resource(gvr).DeleteCollection(context.TODO(), metav1.DeleteOptions{}, listOpts)
	if err != nil {
		return err
	}

	return nil
}

type PatchOp struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

type Patch []PatchOp

func (p Patch) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// TODO: build PatchOP in common-library
func UpdateParentWithAddedChild(parentCrdType string, parentCrdInfo model.NodeInfo, labels map[string]string, childCrdType string, childName string, childHashedName string) error {
	var (
		patchType types.PatchType
		marshaled []byte
	)

	parentParts := strings.Split(parentCrdType, ".")
	gvr := schema.GroupVersionResource{
		Group:    strings.Join(parentParts[1:], "."),
		Version:  "v1",
		Resource: parentParts[0],
	}

	parentName := labels[parentCrdType]
	hashedParentName := nexus.GetHashedName(parentCrdType, parentCrdInfo.ParentHierarchy, labels, parentName)
	childGvk := parentCrdInfo.Children[childCrdType]

	if childGvk.IsNamed {
		payload := "{\"spec\": {\"" + childGvk.FieldNameGvk + "\": {\"" + childName + "\": {\"name\": \"" + childHashedName + "\"}}}}"
		patchType = types.MergePatchType
		marshaled = []byte(payload)
	} else {
		var patch Patch
		patchOp := PatchOp{
			Op:   "replace",
			Path: "/spec/" + childGvk.FieldNameGvk,
			Value: map[string]interface{}{
				"name": childHashedName,
			},
		}
		patch = append(patch, patchOp)
		patchBytes, err := patch.Marshal()
		if err != nil {
			return err
		}
		marshaled = patchBytes
		patchType = types.JSONPatchType
	}

	_, err := Client.Resource(gvr).Patch(context.TODO(), hashedParentName, patchType, marshaled, metav1.PatchOptions{})
	if err != nil {
		return err
	}

	return nil
}

func UpdateParentWithRemovedChild(parentCrdType string, parentCrdInfo model.NodeInfo, labels map[string]string, childCrdType string, childName string) error {
	parentParts := strings.Split(parentCrdType, ".")
	gvr := schema.GroupVersionResource{
		Group:    strings.Join(parentParts[1:], "."),
		Version:  "v1",
		Resource: parentParts[0],
	}

	parentName := labels[parentCrdType]
	hashedParentName := nexus.GetHashedName(parentCrdType, parentCrdInfo.ParentHierarchy, labels, parentName)
	childGvk := parentCrdInfo.Children[childCrdType]

	var patchOp PatchOp
	if childGvk.IsNamed {
		patchOp = PatchOp{
			Op:   "remove",
			Path: "/spec/" + childGvk.FieldNameGvk + "/" + childName,
		}
	} else {
		patchOp = PatchOp{
			Op:   "remove",
			Path: "/spec/" + childGvk.FieldNameGvk,
		}
	}

	var patch Patch
	patch = append(patch, patchOp)

	marshaled, err := patch.Marshal()
	if err != nil {
		return err
	}

	_, err = Client.Resource(gvr).Patch(context.TODO(), hashedParentName, types.JSONPatchType, marshaled, metav1.PatchOptions{})
	if err != nil {
		return err
	}
	return nil
}
