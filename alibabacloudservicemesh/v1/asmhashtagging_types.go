package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ASMHashTaggingPolicy struct {
	PartitionSize uint32 `json:"partitionSize,omitempty"`
	TagValue      string `json:"tagValue,omitempty"`
}

type ASMHashTaggingMatch struct {
	Host *string `json:"host,omitempty"`
}

type ASMHashTaggingRule struct {
	Name        *string                `json:"name,omitempty"`
	Match       *ASMHashTaggingMatch   `json:"match,omitempty"`
	Description *string                `json:"description,omitempty"`
	Header      string                 `json:"header"`
	Modulo      uint32                 `json:"modulo"`
	TagHeader   string                 `json:"tagHeader"`
	Partitions  []ASMHashTaggingPolicy `json:"partitions,omitempty"`
}

type ASMHashTaggingSpec struct {
	WorkloadSelector map[string]string    `json:"workloadSelector,omitempty"`
	Rules            []ASMHashTaggingRule `json:"rules,omitempty"`
}

type ASMHashTaggingStatus struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=ht

// ASMHashTagging is the Schema for the ASMHashTagging API
// +genclient
type ASMHashTagging struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ASMHashTaggingSpec   `json:"spec,omitempty"`
	Status            ASMHashTaggingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ASMHashTaggingList contains a list of ASMHashTagging
type ASMHashTaggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMHashTagging `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMHashTagging{}, &ASMHashTaggingList{})
}
