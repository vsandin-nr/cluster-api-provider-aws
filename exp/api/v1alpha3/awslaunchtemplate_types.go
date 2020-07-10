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

// AWSLaunchTemplateSpec defines the desired state of AWSLaunchTemplate
type AWSLaunchTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of AWSLaunchTemplate. Edit AWSLaunchTemplate_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// AWSLaunchTemplateStatus defines the observed state of AWSLaunchTemplate
type AWSLaunchTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
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
