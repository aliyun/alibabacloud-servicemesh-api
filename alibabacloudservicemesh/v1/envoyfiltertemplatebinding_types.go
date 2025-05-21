/*
Copyright 2019 Alibaba Cloud.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EnvoyFilterTemplateBindingSpec defines the desired state of EnvoyFilterTemplateBinding
type EnvoyFilterTemplateBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	TemplateRef EnvoyFilterTemplateRef `json:"templateRef" protobuf:"bytes,1,opt,name=templateRef"`
	Workloads   []WorkloadRef          `json:"workloads" protobuf:"bytes,2,opt,name=workloads"`
}

type EnvoyFilterTemplateRef struct {
	Name         string `json:"name" protobuf:"bytes,2,opt,name=name"`
	IstioVersion string `json:"istioVersion" protobuf:"bytes,3,opt,name=istioVersion"`
}

type WorkloadRef struct {
	Namespace string            `json:"namespace" protobuf:"bytes,2,opt,name=namespace"`
	Name      string            `json:"name" protobuf:"bytes,2,opt,name=name"`
	Kind      string            `json:"kind" protobuf:"bytes,3,opt,name=kind"`
	Selector  map[string]string `json:"selector,omitempty" protobuf:"bytes,4,opt,name=selector"`
}

// EnvoyFilterTemplateBindingStatus defines the observed state of EnvoyFilterTemplateBinding
type EnvoyFilterTemplateBindingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Status       ConfigState `json:"Status"`
	ErrorMessage string      `json:"ErrorMessage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvoyFilterTemplateBinding is the Schema for the envoyfiltertemplatebindings API
// +genclient
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=envoyfiltertemplatebindings,scope=Namespaced
type EnvoyFilterTemplateBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvoyFilterTemplateBindingSpec   `json:"spec,omitempty"`
	Status EnvoyFilterTemplateBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvoyFilterTemplateBindingList contains a list of EnvoyFilterTemplateBinding
type EnvoyFilterTemplateBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvoyFilterTemplateBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnvoyFilterTemplateBinding{}, &EnvoyFilterTemplateBindingList{})
}
