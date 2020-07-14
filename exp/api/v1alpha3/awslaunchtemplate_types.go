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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EBS
type EBS struct {
	Encrypted  string
	VolumeSize string
	VolumeType string
}

// BlockDeviceMappings
type BlockDeviceMapping struct {
	DeviceName string
	Ebs        EBS
}

// NetworkInterface
type NetworkInterface struct {
	DeviceIndex string
	Groups      []string
}

// AWSLaunchTemplateSpec defines the desired state of AWSLaunchTemplate
type AWSLaunchTemplateSpec struct {
	// all the things needed for a launch template

	IamInstanceProfile  string               `json:"iaminstanceprofile,omitempty"`
	BlockDeviceMappings []BlockDeviceMapping `json:"blockdevicemappings,omitempty"`
	NetworkInterfaces   []NetworkInterface   `json:"networkinterfaces,omitempty"`

	// todo: use a helper
	ImageId string `json:"imageid,omitempty"`

	// InstanceType is the type of instance to create. Example: m4.xlarge
	InstanceType string `json:"instanceType,omitempty"`

	// UncompressedUserData specify whether the user data is gzip-compressed before it is sent to ec2 instance.
	// cloud-init has built-in support for gzip-compressed user data
	// user data stored in aws secret manager is always gzip-compressed.
	//
	// +optional
	UncompressedUserData *bool `json:"uncompressedUserData,omitempty"`

	// SSHKeyName is the name of the ssh key to attach to the instance. Valid values are empty string (do not use SSH keys), a valid SSH key name, or omitted (use the default SSH key name)
	// +optional
	SSHKeyName *string `json:"sshKeyName,omitempty"`
}

// AWSLaunchTemplateStatus defines the observed state of AWSLaunchTemplate
type AWSLaunchTemplateStatus struct {
	LaunchTemplateID bool `json:"launchtemplateid"`
}

// +kubebuilder:object:root=true

// AWSLaunchTemplate is the Schema for the awslaunchtemplates API
type AWSLaunchTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSLaunchTemplateSpec   `json:"spec,omitempty"`
	Status AWSLaunchTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSLaunchTemplateList contains a list of AWSLaunchTemplate
type AWSLaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSLaunchTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSLaunchTemplate{}, &AWSLaunchTemplateList{})
}
