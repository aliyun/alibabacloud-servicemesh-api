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

// EnvoyFilterTemplateSpec defines the desired state of EnvoyFilterTemplate
type EnvoyFilterTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Templates []TemplateValue `json:"templates" protobuf:"bytes,1,opt,name=templates"`
}

type TemplateValue struct {
	// IstioVersion representes which istio version this template is corresponed to
	IstioVersion string `json:"istioVersion" protobuf:"bytes,1,opt,name=istioVersion"`
	// value representes the EnvoyFilter spec value of this template
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

// EnvoyFilterTemplateStatus defines the observed state of EnvoyFilterTemplate
type EnvoyFilterTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Status       ConfigState `json:"Status"`
	ErrorMessage string      `json:"ErrorMessage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvoyFilterTemplate is the Schema for the envoyfiltertemplates API
// +genclient
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=envoyfiltertemplates,scope=Namespaced
type EnvoyFilterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvoyFilterTemplateSpec   `json:"spec,omitempty"`
	Status EnvoyFilterTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvoyFilterTemplateList contains a list of EnvoyFilterTemplate
type EnvoyFilterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvoyFilterTemplate `json:"items"`
}