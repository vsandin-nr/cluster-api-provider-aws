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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	iamv1alpha1 "sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/api/iam/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSIAMConfiguration) DeepCopyInto(out *AWSIAMConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSIAMConfiguration.
func (in *AWSIAMConfiguration) DeepCopy() *AWSIAMConfiguration {
	if in == nil {
		return nil
	}
	out := new(AWSIAMConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSIAMConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSIAMConfigurationSpec) DeepCopyInto(out *AWSIAMConfigurationSpec) {
	*out = *in
	if in.NameSuffix != nil {
		in, out := &in.NameSuffix, &out.NameSuffix
		*out = new(string)
		**out = **in
	}
	in.ControlPlane.DeepCopyInto(&out.ControlPlane)
	in.ClusterAPIControllers.DeepCopyInto(&out.ClusterAPIControllers)
	in.Nodes.DeepCopyInto(&out.Nodes)
	in.BootstrapUser.DeepCopyInto(&out.BootstrapUser)
	if in.EKS != nil {
		in, out := &in.EKS, &out.EKS
		*out = new(EKSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SecureSecretsBackends != nil {
		in, out := &in.SecureSecretsBackends, &out.SecureSecretsBackends
		*out = make([]v1alpha3.SecretBackend, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSIAMConfigurationSpec.
func (in *AWSIAMConfigurationSpec) DeepCopy() *AWSIAMConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(AWSIAMConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSIAMRoleSpec) DeepCopyInto(out *AWSIAMRoleSpec) {
	*out = *in
	if in.ExtraPolicyAttachments != nil {
		in, out := &in.ExtraPolicyAttachments, &out.ExtraPolicyAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraStatements != nil {
		in, out := &in.ExtraStatements, &out.ExtraStatements
		*out = make([]iamv1alpha1.StatementEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrustStatements != nil {
		in, out := &in.TrustStatements, &out.TrustStatements
		*out = make([]iamv1alpha1.StatementEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(v1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSIAMRoleSpec.
func (in *AWSIAMRoleSpec) DeepCopy() *AWSIAMRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AWSIAMRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapUser) DeepCopyInto(out *BootstrapUser) {
	*out = *in
	if in.ExtraPolicyAttachments != nil {
		in, out := &in.ExtraPolicyAttachments, &out.ExtraPolicyAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraGroups != nil {
		in, out := &in.ExtraGroups, &out.ExtraGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraStatements != nil {
		in, out := &in.ExtraStatements, &out.ExtraStatements
		*out = make([]iamv1alpha1.StatementEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(v1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapUser.
func (in *BootstrapUser) DeepCopy() *BootstrapUser {
	if in == nil {
		return nil
	}
	out := new(BootstrapUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAPIControllers) DeepCopyInto(out *ClusterAPIControllers) {
	*out = *in
	in.AWSIAMRoleSpec.DeepCopyInto(&out.AWSIAMRoleSpec)
	if in.AllowedEC2InstanceProfiles != nil {
		in, out := &in.AllowedEC2InstanceProfiles, &out.AllowedEC2InstanceProfiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAPIControllers.
func (in *ClusterAPIControllers) DeepCopy() *ClusterAPIControllers {
	if in == nil {
		return nil
	}
	out := new(ClusterAPIControllers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlane) DeepCopyInto(out *ControlPlane) {
	*out = *in
	in.AWSIAMRoleSpec.DeepCopyInto(&out.AWSIAMRoleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlane.
func (in *ControlPlane) DeepCopy() *ControlPlane {
	if in == nil {
		return nil
	}
	out := new(ControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSConfig) DeepCopyInto(out *EKSConfig) {
	*out = *in
	in.DefaultControlPlaneRole.DeepCopyInto(&out.DefaultControlPlaneRole)
	if in.ManagedMachinePool != nil {
		in, out := &in.ManagedMachinePool, &out.ManagedMachinePool
		*out = new(AWSIAMRoleSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSConfig.
func (in *EKSConfig) DeepCopy() *EKSConfig {
	if in == nil {
		return nil
	}
	out := new(EKSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodes) DeepCopyInto(out *Nodes) {
	*out = *in
	in.AWSIAMRoleSpec.DeepCopyInto(&out.AWSIAMRoleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodes.
func (in *Nodes) DeepCopy() *Nodes {
	if in == nil {
		return nil
	}
	out := new(Nodes)
	in.DeepCopyInto(out)
	return out
}
