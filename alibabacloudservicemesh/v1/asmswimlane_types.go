package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized

type ValidationMessageLevel string

const (
	ValidationMessageLevel_Info    ValidationMessageLevel = "info"
	ValidationMessageLevel_Warning ValidationMessageLevel = "warning"
	ValidationMessageLevel_Error   ValidationMessageLevel = "error"
)

type StringMatch struct {
	// tension: only one of the StringMatch type should be used
	Exact  string `json:"exact,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	Regex  string `json:"regex,omitempty"`
}

type RouteDestination struct {
	Host    string `json:"host,omitempty"`
	PortNum uint32 `json:"portNum,omitempty"`
}

type SwimLaneIngressRoute struct {
	Destination RouteDestination `json:"destination,omitempty"`
}

type SwimLaneIngressRequestMatch struct {
	Uri     StringMatch            `json:"uri,omitempty"`
	Headers map[string]StringMatch `json:"headers,omitempty"`
}

type SwimLaneIngressRule struct {
	Online bool                        `json:"online,omitempty"`
	Hosts  []string                    `json:"hosts,omitempty"`
	Name   string                      `json:"name,omitempty"`
	Match  SwimLaneIngressRequestMatch `json:"match,omitempty"`
	Route  SwimLaneIngressRoute        `json:"route,omitempty"`
}

type SwimLaneIngressWeight struct {
	Destination RouteDestination `json:"destination,omitempty"`
	Weight      int              `json:"weight,omitempty"`
}

type ValidationMessage struct {
	Code    string                 `json:"code,omitempty"`
	Level   ValidationMessageLevel `json:"level,omitempty"`
	Message string                 `json:"message,omitempty"`
}

// ASMSwimLaneSpec defines the desired state of ASMSwimLane
// +k8s:openapi-gen=true
type ASMSwimLaneSpec struct {
	LabelSelector map[string]string     `json:"labelSelector,omitempty"`
	Services      []ClusterServiceRef   `json:"services,omitempty"`
	IngressRules  []SwimLaneIngressRule `json:"ingressRules,omitempty"`
	IngressWeight SwimLaneIngressWeight `json:"ingressWeight,omitempty"`
}

// ASMSwimLaneStatus defines the observed state of ASMSwimLane
// +k8s:openapi-gen=true
type ASMSwimLaneStatus struct {
	Error              string              `json:"error,omitempty"`
	ValidationMessages []ValidationMessage `json:"validationMessage,omitempty"`
}

// ASMSwimLane is the Schema for the asmswimlanes API
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type ASMSwimLane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMSwimLaneSpec   `json:"spec,omitempty"`
	Status ASMSwimLaneStatus `json:"status,omitempty"`
}

// ASMSwimLaneList contains a list of ASMSwimLane
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ASMSwimLaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMSwimLane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMSwimLane{}, &ASMSwimLaneList{})
}
