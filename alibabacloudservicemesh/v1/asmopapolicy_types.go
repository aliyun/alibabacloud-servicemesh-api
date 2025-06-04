package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ASMOPAPolicySpec defines the desired state of ASMOPAPolicy
// +k8s:openapi-gen=true
type ASMOPAPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Criteria used to select the specific set of pods/VMs on which this
	// opa policy should be applied. If omitted, the opa policy will be applied to all workload instances in the same namespace.
	WorkloadSelector WorkloadSelector `json:"workloadSelector,omitempty"`
	Policy           string           `json:"policy,omitempty"`
	PolicyRegoID     string           `json:"policyRegoID,omitempty"`
}

// ASMOPAPolicyStatus defines the observed state of ASMOPAPolicy
// +k8s:openapi-gen=true
type ASMOPAPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status       ConfigState `json:"Status"`
	ErrorMessage string      `json:"ErrorMessage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMOPAPolicy is the Schema for the asmopapolicy API
// +genclient
// +k8s:openapi-gen=true
type ASMOPAPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMOPAPolicySpec   `json:"spec,omitempty"`
	Status ASMOPAPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMOPAPolicyList contains a list of ASMOPAPolicy
type ASMOPAPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMOPAPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMOPAPolicy{}, &ASMOPAPolicyList{})
}
