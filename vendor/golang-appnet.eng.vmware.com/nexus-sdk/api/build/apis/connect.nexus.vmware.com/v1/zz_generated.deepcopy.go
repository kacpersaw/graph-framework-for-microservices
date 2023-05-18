//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Child) DeepCopyInto(out *Child) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Child.
func (in *Child) DeepCopy() *Child {
	if in == nil {
		return nil
	}
	out := new(Child)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connect) DeepCopyInto(out *Connect) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connect.
func (in *Connect) DeepCopy() *Connect {
	if in == nil {
		return nil
	}
	out := new(Connect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Connect) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectList) DeepCopyInto(out *ConnectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Connect, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectList.
func (in *ConnectList) DeepCopy() *ConnectList {
	if in == nil {
		return nil
	}
	out := new(ConnectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectNexusStatus) DeepCopyInto(out *ConnectNexusStatus) {
	*out = *in
	out.Nexus = in.Nexus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectNexusStatus.
func (in *ConnectNexusStatus) DeepCopy() *ConnectNexusStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectNexusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectSpec) DeepCopyInto(out *ConnectSpec) {
	*out = *in
	if in.EndpointsGvk != nil {
		in, out := &in.EndpointsGvk, &out.EndpointsGvk
		*out = make(map[string]Child, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReplicationConfigGvk != nil {
		in, out := &in.ReplicationConfigGvk, &out.ReplicationConfigGvk
		*out = make(map[string]Child, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectSpec.
func (in *ConnectSpec) DeepCopy() *ConnectSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hierarchy) DeepCopyInto(out *Hierarchy) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]KVP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hierarchy.
func (in *Hierarchy) DeepCopy() *Hierarchy {
	if in == nil {
		return nil
	}
	out := new(Hierarchy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVP) DeepCopyInto(out *KVP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVP.
func (in *KVP) DeepCopy() *KVP {
	if in == nil {
		return nil
	}
	out := new(KVP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Link) DeepCopyInto(out *Link) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NexusEndpoint) DeepCopyInto(out *NexusEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NexusEndpoint.
func (in *NexusEndpoint) DeepCopy() *NexusEndpoint {
	if in == nil {
		return nil
	}
	out := new(NexusEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NexusEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NexusEndpointList) DeepCopyInto(out *NexusEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NexusEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NexusEndpointList.
func (in *NexusEndpointList) DeepCopy() *NexusEndpointList {
	if in == nil {
		return nil
	}
	out := new(NexusEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NexusEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NexusEndpointNexusStatus) DeepCopyInto(out *NexusEndpointNexusStatus) {
	*out = *in
	out.Nexus = in.Nexus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NexusEndpointNexusStatus.
func (in *NexusEndpointNexusStatus) DeepCopy() *NexusEndpointNexusStatus {
	if in == nil {
		return nil
	}
	out := new(NexusEndpointNexusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NexusEndpointSpec) DeepCopyInto(out *NexusEndpointSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NexusEndpointSpec.
func (in *NexusEndpointSpec) DeepCopy() *NexusEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(NexusEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NexusStatus) DeepCopyInto(out *NexusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NexusStatus.
func (in *NexusStatus) DeepCopy() *NexusStatus {
	if in == nil {
		return nil
	}
	out := new(NexusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectType) DeepCopyInto(out *ObjectType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectType.
func (in *ObjectType) DeepCopy() *ObjectType {
	if in == nil {
		return nil
	}
	out := new(ObjectType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfig) DeepCopyInto(out *ReplicationConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfig.
func (in *ReplicationConfig) DeepCopy() *ReplicationConfig {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigList) DeepCopyInto(out *ReplicationConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigList.
func (in *ReplicationConfigList) DeepCopy() *ReplicationConfigList {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigNexusStatus) DeepCopyInto(out *ReplicationConfigNexusStatus) {
	*out = *in
	out.Nexus = in.Nexus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigNexusStatus.
func (in *ReplicationConfigNexusStatus) DeepCopy() *ReplicationConfigNexusStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigNexusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigSpec) DeepCopyInto(out *ReplicationConfigSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	if in.RemoteEndpointGvk != nil {
		in, out := &in.RemoteEndpointGvk, &out.RemoteEndpointGvk
		*out = new(Link)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigSpec.
func (in *ReplicationConfigSpec) DeepCopy() *ReplicationConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestination) DeepCopyInto(out *ReplicationDestination) {
	*out = *in
	in.Hierarchy.DeepCopyInto(&out.Hierarchy)
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(ObjectType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestination.
func (in *ReplicationDestination) DeepCopy() *ReplicationDestination {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSource) DeepCopyInto(out *ReplicationSource) {
	*out = *in
	out.Type = in.Type
	in.Object.DeepCopyInto(&out.Object)
	in.Filters.DeepCopyInto(&out.Filters)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSource.
func (in *ReplicationSource) DeepCopy() *ReplicationSource {
	if in == nil {
		return nil
	}
	out := new(ReplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceFilters) DeepCopyInto(out *SourceFilters) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]KVP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceFilters.
func (in *SourceFilters) DeepCopy() *SourceFilters {
	if in == nil {
		return nil
	}
	out := new(SourceFilters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceObject) DeepCopyInto(out *SourceObject) {
	*out = *in
	out.ObjectType = in.ObjectType
	in.Hierarchy.DeepCopyInto(&out.Hierarchy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceObject.
func (in *SourceObject) DeepCopy() *SourceObject {
	if in == nil {
		return nil
	}
	out := new(SourceObject)
	in.DeepCopyInto(out)
	return out
}
