// +kubebuilder:object:generate=true
// +groupName=istio.alibabacloud.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// controller-gen crd:crdVersions=v1,allowDangerousTypes=true paths="./pkg/apis/adaptivescheduler/policy/v1/..." output:crd:artifacts:config=istiocrds/adaptivescheduler/controlplane
// go run ./pkg/crds/adaptivescheduler/controlplane/generate.go
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ConcurrencySchedulingPolicyStatus struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ConcurrencyScheduler struct {
	Alerter            *Alerter                      `json:"alerter,omitempty"`
	ConcurrencyLimiter *ConcurrencyLimiterParameters `json:"concurrency_limiter,omitempty"`
	MaxConcurrency     *int64                        `json:"max_concurrency,omitempty"`
	Scheduler          *Scheduler                    `json:"scheduler,omitempty"`
	Selectors          []Selector                    `json:"selectors,omitempty"`
	MatchRequests      []MatchRequest                `json:"match_requests,omitempty"`
}

type ConcurrencySchedulingPolicySpec struct {
	Resources            *Resources            `json:"resources,omitempty"`
	ConcurrencyScheduler *ConcurrencyScheduler `json:"concurrency_scheduler,omitempty"`
	WorkloadSelector     *WorkloadSelector     `json:"workload_selector,omitempty"`
	ApplyToTraffic       *ApplyToTraffic       `json:"apply_to_traffic,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:storageversion
// Policy is the Schema for the policies API.
// +genclient
type ConcurrencySchedulingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	//+kubebuilder:pruning:PreserveUnknownFields
	Spec   ConcurrencySchedulingPolicySpec   `json:"spec,omitempty"`
	Status ConcurrencySchedulingPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PolicyList contains a list of Policy.
type ConcurrencySchedulingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConcurrencySchedulingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConcurrencySchedulingPolicy{}, &ConcurrencySchedulingPolicyList{})
}
