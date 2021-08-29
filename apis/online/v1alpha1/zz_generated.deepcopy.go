//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Archive) DeepCopyInto(out *Archive) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Archive.
func (in *Archive) DeepCopy() *Archive {
	if in == nil {
		return nil
	}
	out := new(Archive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Archive) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveList) DeepCopyInto(out *ArchiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Archive, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveList.
func (in *ArchiveList) DeepCopy() *ArchiveList {
	if in == nil {
		return nil
	}
	out := new(ArchiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArchiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveSpec) DeepCopyInto(out *ArchiveSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ArchiveSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveSpec.
func (in *ArchiveSpec) DeepCopy() *ArchiveSpec {
	if in == nil {
		return nil
	}
	out := new(ArchiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveSpecCriteria) DeepCopyInto(out *ArchiveSpecCriteria) {
	*out = *in
	if in.DateField != nil {
		in, out := &in.DateField, &out.DateField
		*out = new(string)
		**out = **in
	}
	if in.DateFormat != nil {
		in, out := &in.DateFormat, &out.DateFormat
		*out = new(string)
		**out = **in
	}
	if in.ExpireAfterDays != nil {
		in, out := &in.ExpireAfterDays, &out.ExpireAfterDays
		*out = new(int64)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveSpecCriteria.
func (in *ArchiveSpecCriteria) DeepCopy() *ArchiveSpecCriteria {
	if in == nil {
		return nil
	}
	out := new(ArchiveSpecCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveSpecPartitionFields) DeepCopyInto(out *ArchiveSpecPartitionFields) {
	*out = *in
	if in.FieldName != nil {
		in, out := &in.FieldName, &out.FieldName
		*out = new(string)
		**out = **in
	}
	if in.FieldType != nil {
		in, out := &in.FieldType, &out.FieldType
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveSpecPartitionFields.
func (in *ArchiveSpecPartitionFields) DeepCopy() *ArchiveSpecPartitionFields {
	if in == nil {
		return nil
	}
	out := new(ArchiveSpecPartitionFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveSpecResource) DeepCopyInto(out *ArchiveSpecResource) {
	*out = *in
	if in.ArchiveID != nil {
		in, out := &in.ArchiveID, &out.ArchiveID
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CollName != nil {
		in, out := &in.CollName, &out.CollName
		*out = new(string)
		**out = **in
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = new(ArchiveSpecCriteria)
		(*in).DeepCopyInto(*out)
	}
	if in.DbName != nil {
		in, out := &in.DbName, &out.DbName
		*out = new(string)
		**out = **in
	}
	if in.PartitionFields != nil {
		in, out := &in.PartitionFields, &out.PartitionFields
		*out = make([]ArchiveSpecPartitionFields, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.SyncCreation != nil {
		in, out := &in.SyncCreation, &out.SyncCreation
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveSpecResource.
func (in *ArchiveSpecResource) DeepCopy() *ArchiveSpecResource {
	if in == nil {
		return nil
	}
	out := new(ArchiveSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveStatus) DeepCopyInto(out *ArchiveStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveStatus.
func (in *ArchiveStatus) DeepCopy() *ArchiveStatus {
	if in == nil {
		return nil
	}
	out := new(ArchiveStatus)
	in.DeepCopyInto(out)
	return out
}
