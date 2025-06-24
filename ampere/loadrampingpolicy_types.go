// +kubebuilder:object:generate=true
// +groupName=istio.alibabacloud.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LoadRampingPolicyStatus defines the observed state of Policy.
type LoadRampingPolicyStatus struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:storageversion
// Policy is the Schema for the policies API.
// +genclient
type LoadRampingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	//+kubebuilder:pruning:PreserveUnknownFields
	Spec   LoadRampingPolicySpec   `json:"spec,omitempty"`
	Status LoadRampingPolicyStatus `json:"status,omitempty"`
}

type Criteria struct {
	Threshold *float64 `json:"threshold,omitempty"`
}

type DriverCriteria struct {
	Backward *Criteria `json:"backward,omitempty"`
	Forward  *Criteria `json:"forward,omitempty"`
	Reset    *Criteria `json:"reset,omitempty"`
}

type AverageLatencyDriver struct {
	Selectors     *[]Selector     `json:"selectors,omitempty"`
	MatchRequests []MatchRequest  `json:"match_requests,omitempty"`
	Criteria      *DriverCriteria `json:"criteria,omitempty"`
}

type PercentileLatencyDriver struct {
	Criteria   *DriverCriteria `json:"criteria,omitempty"`
	FluxMeter  *Meter          `json:"flux_meter,omitempty"`
	Percentile *float64        `json:"percentile,omitempty"`
}

type PromqlDriver struct {
	Criteria    *DriverCriteria `json:"criteria,omitempty"`
	QueryString *string         `json:"query_string,omitempty"`
}

type Drivers struct {
	AverageLatencyDrivers    *[]AverageLatencyDriver    `json:"average_latency_drivers,omitempty"`
	PercentileLatencyDrivers *[]PercentileLatencyDriver `json:"percentile_latency_drivers,omitempty"`
	PromqlDrivers            *[]PromqlDriver            `json:"promql_drivers,omitempty"`
}

type Sampler struct {
	Selectors                *[]Selector    `json:"selectors,omitempty"`
	MatchRequests            []MatchRequest `json:"match_requests,omitempty"`
	SessionLabelKey          *string        `json:"session_label_key,omitempty"`
	DeniedResponseStatusCode *string        `json:"denied_response_status_code,omitempty"`
}

type Step struct {
	Duration               *string  `json:"duration,omitempty"`
	TargetAcceptPercentage *float64 `json:"target_accept_percentage,omitempty"`
}

type LoadRamp struct {
	Sampler *Sampler `json:"sampler,omitempty"`
	Steps   *[]Step  `json:"steps,omitempty"`
}

type LoadRampingPolicySpec struct {
	Resources        *Resources        `json:"resources,omitempty"`
	Drivers          *Drivers          `json:"drivers,omitempty"`
	Start            *bool             `json:"start,omitempty"`
	Reset            *bool             `json:"reset,omitempty"`
	LoadRamp         *LoadRamp         `json:"load_ramp,omitempty"`
	WorkloadSelector *WorkloadSelector `json:"workload_selector,omitempty"`
	ApplyToTraffic   *ApplyToTraffic   `json:"apply_to_traffic,omitempty"`
}

//+kubebuilder:object:root=true

// PolicyList contains a list of Policy.
type LoadRampingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadRampingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoadRampingPolicy{}, &LoadRampingPolicyList{})
}
