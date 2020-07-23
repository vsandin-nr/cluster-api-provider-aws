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
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
)

// EBS from describe-launch-templates
type EBS struct {
	Encrypted  bool   `json:"encrypted,omitempty"`
	VolumeSize int64  `json:"volumeSize,omitempty"`
	VolumeType string `json:"volumeType,omitempty"`
}

// BlockDeviceMappings from describe-launch-templates
type BlockDeviceMapping struct {
	DeviceName string `json:"deviceName,omitempty"`
	Ebs        EBS    `json:"ebs,omitempty"`
}

// NetworkInterface from describe-launch-templates
type NetworkInterface struct {
	DeviceIndex int64    `json:"deviceIndex,omitempty"`
	Groups      []string `json:"groups,omitempty"`
}

// AwsLaunchTemplate defines the desired state of AWSLaunchTemplate
type AWSLaunchTemplate struct {
	// all the things needed for a launch template

	IamInstanceProfile  string               `json:"iamInstanceProfile,omitempty"`
	BlockDeviceMappings []BlockDeviceMapping `json:"blockDeviceMappings,omitempty"`
	NetworkInterfaces   []NetworkInterface   `json:"networkInterfaces,omitempty"`

	// todo: use a helper
	AMI infrav1.AWSResourceReference `json:"ami,omitempty"`

	// InstanceType is the type of instance to create. Example: m4.xlarge
	InstanceType string `json:"instanceType,omitempty"`

	// SSHKeyName is the name of the ssh key to attach to the instance. Valid values are empty string (do not use SSH keys), a valid SSH key name, or omitted (use the default SSH key name)
	// +optional
	SSHKeyName *string `json:"sshKeyName,omitempty"`

	VersionNumber *int64 `json:"versionNumber,omitempty"`
}

// LaunchTemplateSpecification from describe-auto-scaling-groups
type LaunchTemplateSpecification struct {
	LaunchTemplateID   string `json:"launchTemplateId,omitempty"`
	LaunchTemplateName string `json:"launchTemplateName,omitempty"`
	Version            string `json:"version,omitempty"`
}

// LaunchTemplate from describe-auto-scaling-groups
type LaunchTemplate struct {
	LaunchTemplateSpecification LaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty"`
	Overrides                   []Overrides                 `json:"overrides,omitempty"`
}

// Overrides from describe-auto-scaling-groups
type Overrides struct {
	InstanceType string `json:"InstanceType"`
}

// InstancesDistribution from describe-auto-scaling-groups
type InstancesDistribution struct {
	OnDemandAllocationStrategy          string `json:"onDemandAllocationStrategy,omitempty"`
	OnDemandBaseCapacity                int    `json:"onDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity int    `json:"onDemandPercentageAboveBaseCapacity,omitempty"`
	SpotAllocationStrategy              string `json:"spotAllocationStrategy,omitempty"`
}

// MixedInstancesPolicy from describe-auto-scaling-groups
type MixedInstancesPolicy struct {
	LaunchTemplate        LaunchTemplate        `json:"launchTemplate,omitempty"`
	InstancesDistribution InstancesDistribution `json:"instancesDistribution,omitempty"`
}

// Tags
type Tags map[string]string

// AutoScalingGroup describes an AWS autoscaling group.
type AutoScalingGroup struct {
	ID string `json:"id"`
	// The tags associated with the instance.
	Tags                 map[string]string `json:"tags,omitempty"`
	AutoScalingGroupName string
	DesiredCapacity      int64
	// LaunchTemplateSpecification *autoscaling.LaunchTemplateSpecification
	MaxSize int64
	MinSize int64
	// MixedInstancesPolicy        *autoscaling.MixedInstancesPolicy
	PlacementGroup    string
	VPCZoneIdentifier []string
}
