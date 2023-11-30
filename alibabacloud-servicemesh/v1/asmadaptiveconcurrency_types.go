// Copyright 2023 Alibaba Cloud Service Mesh
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Percent defines a type that represents a percent value
type Percent struct {
	Value int32 `json:"value,omitempty"`
}

type ConcurrencyLimitParamsConfiguration struct {
	// The allowed upper-bound on the calculated concurrency limit. Defaults to 1000.
	MaxConcurrencyLimit *int32 `json:"max_concurrency_limit,omitempty"`
	// (Duration, eg: 5s)The period of time samples are taken to recalculate the concurrency limit.
	ConcurrencyUpdateInterval string `json:"concurrency_update_interval,omitempty"`
}

type MinRTTCalcParamsConfiguration struct {
	Interval       string   `json:"interval,omitempty"`
	RequestCount   *int32   `json:"request_count,omitempty"`
	Jitter         *Percent `json:"jitter,omitempty"`
	MinConcurrency *int32   `json:"min_concurrency,omitempty"`
	Buffer         *Percent `json:"buffer,omitempty"`
}

// ASMAdaptiveConcurrencySpec defines the desired state of ASMAdaptiveConcurrency
type ASMAdaptiveConcurrencySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	WorkloadSelector          *WorkloadSelector                   `json:"workload_selector,omitempty"`
	SampleAggregatePercentile *Percent                            `json:"sample_aggregate_percentile,omitempty"`
	ConcurrencyLimitParams    ConcurrencyLimitParamsConfiguration `json:"concurrency_limit_params,omitempty"`
	MinRTTCalcParams          MinRTTCalcParamsConfiguration       `json:"min_rtt_calc_params,omitempty"`
}

// ASMAdaptiveConcurrencyStatus defines the observed state of ASMAdaptiveConcurrency
type ASMAdaptiveConcurrencyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMAdaptiveConcurrency is the Schema for the asmadaptiveconcurrencies API
type ASMAdaptiveConcurrency struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMAdaptiveConcurrencySpec   `json:"spec,omitempty"`
	Status ASMAdaptiveConcurrencyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMAdaptiveConcurrencyList contains a list of ASMAdaptiveConcurrency
type ASMAdaptiveConcurrencyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMAdaptiveConcurrency `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMAdaptiveConcurrency{}, &ASMAdaptiveConcurrencyList{})
}
