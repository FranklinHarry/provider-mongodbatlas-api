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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ConfigurationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecMatcher) DeepCopyInto(out *ConfigurationSpecMatcher) {
	*out = *in
	if in.FieldName != nil {
		in, out := &in.FieldName, &out.FieldName
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecMatcher.
func (in *ConfigurationSpecMatcher) DeepCopy() *ConfigurationSpecMatcher {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecMetricThresholdConfig) DeepCopyInto(out *ConfigurationSpecMetricThresholdConfig) {
	*out = *in
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecMetricThresholdConfig.
func (in *ConfigurationSpecMetricThresholdConfig) DeepCopy() *ConfigurationSpecMetricThresholdConfig {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecMetricThresholdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecNotification) DeepCopyInto(out *ConfigurationSpecNotification) {
	*out = *in
	if in.ApiToken != nil {
		in, out := &in.ApiToken, &out.ApiToken
		*out = new(string)
		**out = **in
	}
	if in.ChannelName != nil {
		in, out := &in.ChannelName, &out.ChannelName
		*out = new(string)
		**out = **in
	}
	if in.DatadogAPIKey != nil {
		in, out := &in.DatadogAPIKey, &out.DatadogAPIKey
		*out = new(string)
		**out = **in
	}
	if in.DatadogRegion != nil {
		in, out := &in.DatadogRegion, &out.DatadogRegion
		*out = new(string)
		**out = **in
	}
	if in.DelayMin != nil {
		in, out := &in.DelayMin, &out.DelayMin
		*out = new(int64)
		**out = **in
	}
	if in.EmailAddress != nil {
		in, out := &in.EmailAddress, &out.EmailAddress
		*out = new(string)
		**out = **in
	}
	if in.EmailEnabled != nil {
		in, out := &in.EmailEnabled, &out.EmailEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FlowName != nil {
		in, out := &in.FlowName, &out.FlowName
		*out = new(string)
		**out = **in
	}
	if in.FlowdockAPIToken != nil {
		in, out := &in.FlowdockAPIToken, &out.FlowdockAPIToken
		*out = new(string)
		**out = **in
	}
	if in.IntervalMin != nil {
		in, out := &in.IntervalMin, &out.IntervalMin
		*out = new(int64)
		**out = **in
	}
	if in.MobileNumber != nil {
		in, out := &in.MobileNumber, &out.MobileNumber
		*out = new(string)
		**out = **in
	}
	if in.OpsGenieAPIKey != nil {
		in, out := &in.OpsGenieAPIKey, &out.OpsGenieAPIKey
		*out = new(string)
		**out = **in
	}
	if in.OpsGenieRegion != nil {
		in, out := &in.OpsGenieRegion, &out.OpsGenieRegion
		*out = new(string)
		**out = **in
	}
	if in.OrgName != nil {
		in, out := &in.OrgName, &out.OrgName
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceKey != nil {
		in, out := &in.ServiceKey, &out.ServiceKey
		*out = new(string)
		**out = **in
	}
	if in.SmsEnabled != nil {
		in, out := &in.SmsEnabled, &out.SmsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.TypeName != nil {
		in, out := &in.TypeName, &out.TypeName
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.VictorOpsAPIKey != nil {
		in, out := &in.VictorOpsAPIKey, &out.VictorOpsAPIKey
		*out = new(string)
		**out = **in
	}
	if in.VictorOpsRoutingKey != nil {
		in, out := &in.VictorOpsRoutingKey, &out.VictorOpsRoutingKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecNotification.
func (in *ConfigurationSpecNotification) DeepCopy() *ConfigurationSpecNotification {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecResource) DeepCopyInto(out *ConfigurationSpecResource) {
	*out = *in
	if in.AlertConfigurationID != nil {
		in, out := &in.AlertConfigurationID, &out.AlertConfigurationID
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventType != nil {
		in, out := &in.EventType, &out.EventType
		*out = new(string)
		**out = **in
	}
	if in.Matcher != nil {
		in, out := &in.Matcher, &out.Matcher
		*out = make([]ConfigurationSpecMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetricThreshold != nil {
		in, out := &in.MetricThreshold, &out.MetricThreshold
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.MetricThresholdConfig != nil {
		in, out := &in.MetricThresholdConfig, &out.MetricThresholdConfig
		*out = new(ConfigurationSpecMetricThresholdConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]ConfigurationSpecNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.ThresholdConfig != nil {
		in, out := &in.ThresholdConfig, &out.ThresholdConfig
		*out = new(ConfigurationSpecThresholdConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecResource.
func (in *ConfigurationSpecResource) DeepCopy() *ConfigurationSpecResource {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecThresholdConfig) DeepCopyInto(out *ConfigurationSpecThresholdConfig) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecThresholdConfig.
func (in *ConfigurationSpecThresholdConfig) DeepCopy() *ConfigurationSpecThresholdConfig {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecThresholdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatus) DeepCopyInto(out *ConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatus.
func (in *ConfigurationStatus) DeepCopy() *ConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
