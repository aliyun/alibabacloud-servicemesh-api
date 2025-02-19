package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LLMProviderSpec defines the desired state of LLMProvider
type LLMProviderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	WorkloadSelector *WorkloadSelector   `json:"workloadSelector,omitempty"`
	Host             *string             `json:"host,omitempty"`
	Path             *string             `json:"path,omitempty"`
	Configs          *LLMProviderConfigs `json:"configs,omitempty"`
}

type LLMProviderConfigs struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	DefaultConfig        *ProviderConfig            `json:"defaultConfig,omitempty"`
	RouteSpecificConfigs map[string]*ProviderConfig `json:"routeSpecificConfigs,omitempty"`
}

type ProviderConfig struct {
	OpenAIConfig *OpenAIConfig `json:"openAIConfig,omitempty"`
}

type OpenAIConfig struct {
	Model  *string `json:"model,omitempty"`
	Stream *bool   `json:"stream,omitempty"`
	APIKey *string `json:"apiKey,omitempty"`
}

// LLMProviderStatus defines the observed state of LLMProvider
type LLMProviderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status *string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LLMProvider is the Schema for the llmproviders API
// +genclient
type LLMProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LLMProviderSpec   `json:"spec,omitempty"`
	Status LLMProviderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LLMProviderList contains a list of LLMProvider
type LLMProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LLMProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LLMProvider{}, &LLMProviderList{})
}
