// +kubebuilder:object:generate=true
// +groupName=istio.alibabacloud.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// controller-gen crd:crdVersions=v1,allowDangerousTypes=true paths="./pkg/apis/adaptivescheduler/policy/v1/..." output:crd:artifacts:config=istiocrds/adaptivescheduler/controlplane
// go run ./pkg/crds/adaptivescheduler/controlplane/generate.go
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RateLimitingPolicyStatus defines the observed state of Policy.
type RateLimitingPolicyStatus struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type Alerter struct {
	AlertName      string            `json:"alert_name,omitempty"`
	Labels         map[string]string `json:"labels,omitempty"`
	ResolveTimeout *metav1.Duration  `json:"resolve_timeout,omitempty"`
	Severity       *string           `json:"severity,omitempty"`
}

type RateLimiterParametersLazySync struct {
	Enabled *bool  `json:"enabled,omitempty"`
	NumSync *int64 `json:"num_sync,omitempty"`
}

type RateLimiterParameters struct {
	ContinuousFill   *bool                          `json:"continuous_fill,omitempty"`
	DelayInitialFill *bool                          `json:"delay_initial_fill,omitempty"`
	Interval         *metav1.Duration               `json:"interval,omitempty"`
	LazySync         *RateLimiterParametersLazySync `json:"lazy_sync,omitempty"`
	LimitByLabelKey  *string                        `json:"limit_by_label_key,omitempty"`
	MaxIdleTime      *metav1.Duration               `json:"max_idle_time,omitempty"`
}

type RateLimiterRequestParameters struct {
	DeniedResponseStatusCode *int    `json:"denied_response_status_code,omitempty"`
	TokensLabelKey           *string `json:"tokens_label_key,omitempty"`
}

type RateLimiter struct {
	Alerter           *Alerter                      `json:"alerter,omitempty"`
	BucketCapacity    *float64                      `json:"bucket_capacity,omitempty"`
	FillAmount        *float64                      `json:"fill_amount,omitempty"`
	Parameters        *RateLimiterParameters        `json:"parameters,omitempty"`
	RequestParameters *RateLimiterRequestParameters `json:"request_parameters,omitempty"`
	Selectors         []Selector                    `json:"selectors,omitempty"`
	MatchRequests     []MatchRequest                `json:"match_requests,omitempty"`
}

type RateLimitingPolicySpec struct {
	Resources        *Resources        `json:"resources,omitempty"`
	RateLimiter      *RateLimiter      `json:"rate_limiter,omitempty"`
	WorkloadSelector *WorkloadSelector `json:"workload_selector,omitempty"`
	ApplyToTraffic   *ApplyToTraffic   `json:"apply_to_traffic,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:storageversion
// Policy is the Schema for the policies API.
// +genclient
type RateLimitingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	//+kubebuilder:pruning:PreserveUnknownFields
	Spec   RateLimitingPolicySpec   `json:"spec,omitempty"`
	Status RateLimitingPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PolicyList contains a list of Policy.
type RateLimitingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RateLimitingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RateLimitingPolicy{}, &RateLimitingPolicyList{})
}
