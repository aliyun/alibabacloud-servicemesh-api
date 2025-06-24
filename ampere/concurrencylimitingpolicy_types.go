// +kubebuilder:object:generate=true
// +groupName=istio.alibabacloud.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// controller-gen crd:crdVersions=v1,allowDangerousTypes=true paths="./pkg/apis/adaptivescheduler/policy/v1/..." output:crd:artifacts:config=istiocrds/adaptivescheduler/controlplane
// go run ./pkg/crds/adaptivescheduler/controlplane/generate.go
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConcurrencyPolicyStatus defines the observed state of Policy
type ConcurrencyLimitingPolicyStatus struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ConcurrencyLimiterParameters struct {
	LimitByLabelKey     *string          `json:"limit_by_label_key,omitempty"`
	MaxIdleTime         *metav1.Duration `json:"max_idle_time,omitempty"`
	MaxInflightDuration *metav1.Duration `json:"max_inflight_duration,omitempty"`
}

type ConcurrencyLimiter struct {
	MaxConcurrency    *int64                        `json:"max_concurrency,omitempty"`
	Parameters        *ConcurrencyLimiterParameters `json:"parameters,omitempty"`
	RequestParameters *RateLimiterRequestParameters `json:"request_parameters,omitempty"`
	Alerter           *Alerter                      `json:"alerter,omitempty"`
	Selectors         []Selector                    `json:"selectors,omitempty"`
	MatchRequests     []MatchRequest                `json:"match_requests,omitempty"`
}

type ConcurrencyLimitingPolicySpec struct {
	Resources          *Resources          `json:"resources,omitempty"`
	ConcurrencyLimiter *ConcurrencyLimiter `json:"concurrency_limiter,omitempty"`
	WorkloadSelector   *WorkloadSelector   `json:"workload_selector,omitempty"`
	ApplyToTraffic     *ApplyToTraffic     `json:"apply_to_traffic,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:storageversion
// Policy is the Schema for the policies API.
// +genclient
type ConcurrencyLimitingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	//+kubebuilder:pruning:PreserveUnknownFields
	Spec   ConcurrencyLimitingPolicySpec   `json:"spec,omitempty"`
	Status ConcurrencyLimitingPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PolicyList contains a list of Policy.
type ConcurrencyLimitingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConcurrencyLimitingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConcurrencyLimitingPolicy{}, &ConcurrencyLimitingPolicyList{})
}
