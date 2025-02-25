/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"connector/pkg/utils"
)

// CustomResourceDefinitionReconciler reconciles a CustomResourceDefinition object
type CustomResourceDefinitionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Cache  *GvrCache
}

//+kubebuilder:rbac:groups=apiextensions.k8s.io.api-gw.com,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apiextensions.k8s.io.api-gw.com,resources=customresourcedefinitions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apiextensions.k8s.io.api-gw.com,resources=customresourcedefinitions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CustomResourceDefinition object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile

type Gvr struct {
	Value schema.GroupVersionResource
}

var GvrCh = make(chan schema.GroupVersionResource, 100)

func (r *CustomResourceDefinitionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var crd apiextensionsv1.CustomResourceDefinition
	eventType := utils.Upsert
	if err := r.Get(ctx, req.NamespacedName, &crd); err != nil {
		if client.IgnoreNotFound(err) != nil {
			return ctrl.Result{}, err
		}
		eventType = utils.Delete
	}

	if utils.NexusDatamodelCRDs(crd.Spec.Group) {
		return ctrl.Result{}, nil
	}
	logrus.Debugf("Received [%s] event for CRD Type %s", eventType, crd.Name)

	var version string
	if eventType != utils.Delete {
		status := &crd.Status
		if status == nil {
			return ctrl.Result{}, fmt.Errorf("status field is empty, hence can't determine the crd version")
		}
		if len(crd.Status.StoredVersions) == 0 {
			return ctrl.Result{}, fmt.Errorf("stored version in status field is empty, hence can't determine the crd version")
		}
		version = crd.Status.StoredVersions[0]
	}

	gvr := r.ConstructGVR(crd.Name, version)
	if err := r.ProcessAnnotation(crd.Name, gvr, crd.Annotations, eventType); err != nil {
		logrus.Errorf("Error Processing CRD Annotation %v\n", err)
		return ctrl.Result{}, err
	}

	if eventType != utils.Delete {
		GvrCh <- gvr
	} else {
		r.Cache.Delete(gvr)
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CustomResourceDefinitionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiextensionsv1.CustomResourceDefinition{}).
		Complete(r)
}
