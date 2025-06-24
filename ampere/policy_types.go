// +kubebuilder:object:generate=true
// +groupName=istio.alibabacloud.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	Status string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:storageversion
// Policy is the Schema for the policies API.
// +genclient
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	//+kubebuilder:pruning:PreserveUnknownFields
	Spec runtime.RawExtension `json:"spec,omitempty"`
	//+kubebuilder:pruning:PreserveUnknownFields
	DynamicConfig runtime.RawExtension `json:"dynamicConfig,omitempty"`
	Status        PolicyStatus         `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PolicyList contains a list of Policy.
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
