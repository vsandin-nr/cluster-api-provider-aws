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

type LaunchTemplateSpecification struct {
	LaunchTemplateID   string `json:"launchTemplateId,omitempty"`
	LaunchTemplateName string `json:"launchTemplateName,omitempty"`
	Version            string `json:"version,omitempty"`
}
type LaunchTemplate struct {
	LaunchTemplateSpecification LaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty"`
	Overrides                   []Overrides                 `json:"overrides,omitempty"`
}

type Overrides struct {
	InstanceType string `json:"InstanceType"`
}

type InstancesDistribution struct {
	OnDemandAllocationStrategy          string `json:"onDemandAllocationStrategy,omitempty"`
	OnDemandBaseCapacity                int    `json:"onDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity int    `json:"onDemandPercentageAboveBaseCapacity,omitempty"`
	SpotAllocationStrategy              string `json:"spotAllocationStrategy,omitempty"`
}
type MixedInstancesPolicy struct {
	LaunchTemplate        LaunchTemplate        `json:"launchTemplate,omitempty"`
	InstancesDistribution InstancesDistribution `json:"instancesDistribution,omitempty"`
}

type Tags struct {
	ResourceID        string `json:"resourceId,omitempty"`
	ResourceType      string `json:"resourceType,omitempty"`
	Key               string `json:"key,omitempty"`
	Value             string `json:"value,omitempty"`
	PropagateAtLaunch bool   `json:"propagateAtLaunch,omitempty"`
}

// AWSMachinePoolSpec defines the desired state of AWSMachinePool
type AWSMachinePoolSpec struct {
	AutoScalingGroupName             string               `json:"autoScalingGroupName,omitempty"`
	MixedInstancesPolicy             MixedInstancesPolicy `json:"mixedInstancesPolicy,omitempty"`
	MinSize                          int                  `json:"minSize,omitempty"`
	MaxSize                          int                  `json:"maxSize,omitempty"`
	DesiredCapacity                  int                  `json:"desiredCapacity,omitempty"`
	DefaultCooldown                  int                  `json:"defaultCooldown,omitempty"`
	AvailabilityZones                []string             `json:"availabilityZones,omitempty"`
	HealthCheckType                  string               `json:"healthCheckType,omitempty"`
	HealthCheckGracePeriod           int                  `json:"healthCheckGracePeriod,omitempty"`
	VPCZoneIdentifier                string               `json:"vpcZoneIdentifier,omitempty"`
	Tags                             []Tags               `json:"tags,omitempty"`
	TerminationPolicies              []string             `json:"terminationPolicies,omitempty"`
	NewInstancesProtectedFromScaleIn bool                 `json:"newInstancesProtectedFromScaleIn,omitempty"`
	ServiceLinkedRoleARN             string               `json:"serviceLinkedRoleARN,omitempty"`
}

// AWSMachinePoolStatus defines the observed state of AWSMachinePool
type AWSMachinePoolStatus struct {
	AutoScalingGroupARN string `json:"autoScalingGroupARN,omitempty"`
}

// +kubebuilder:object:root=true

// AWSMachinePool is the Schema for the awsmachinepools API
type AWSMachinePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSMachinePoolSpec   `json:"spec,omitempty"`
	Status AWSMachinePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSMachinePoolList contains a list of AWSMachinePool
type AWSMachinePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSMachinePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSMachinePool{}, &AWSMachinePoolList{})
}
