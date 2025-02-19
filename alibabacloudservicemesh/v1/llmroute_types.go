package v1

import (
	networkingv1beta1 "istio.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LLMRouteSpec defines the desired state of LLMRoute
type LLMRouteSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Host     *string    `json:"host,omitempty"`
	Gateways []string   `json:"gateways,omitempty"` // ingress or egress
	Rules    []*LLMRule `json:"rules,omitempty"`
}

type LLMRule struct {
	Name        *string            `json:"name,omitempty"`
	Matches     []*LLMRequestMatch `json:"matches,omitempty"`
	BackendRefs []*LLMBackendRef   `json:"backendRefs,omitempty"`
}

type LLMRequestMatch struct {
	Headers      map[string]*networkingv1beta1.StringMatch `json:"headers,omitempty"`
	SourceLabels map[string]string                         `json:"sourceLabels,omitempty"`
	Gateways     []string                                  `json:"gateways,omitempty"`
}

type LLMBackendRef struct {
	ProviderHost *string `json:"providerHost,omitempty"`
	Weight       int32   `json:"weight,omitempty"`
	HostRewrite  *string `json:"hostRewrite,omitempty"`
}

// LLMRouteStatus defines the observed state of LLMRoute
type LLMRouteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LLMRoute is the Schema for the llmroutes API
// +genclient
type LLMRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LLMRouteSpec   `json:"spec,omitempty"`
	Status LLMRouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LLMRouteList contains a list of LLMRoute
type LLMRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LLMRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LLMRoute{}, &LLMRouteList{})
}
