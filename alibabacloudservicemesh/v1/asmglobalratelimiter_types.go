/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GlobalRateLimitTimeUnit string

const (
	TIMEUNIT_SECOND GlobalRateLimitTimeUnit = "SECOND"
	TIMEUNIT_MINUTE GlobalRateLimitTimeUnit = "MINUTE"
)

type GlobalServiceLimitConfig struct {
	Quota int32                   `json:"quota,omitempty"`
	Unit  GlobalRateLimitTimeUnit `json:"unit,omitempty"`
}

type GlobalRateLimitOverrideConfig struct {
	RequestMatch RequestMatcher           `json:"request_match,omitempty"`
	Limit        GlobalServiceLimitConfig `json:"limit,omitempty"`
}

type GlobalRateLimiterConfig struct {
	Name           string                          `json:"name,omitempty"`
	Match          *RateLimitMatch                 `json:"match,omitempty"`
	Limit          *GlobalServiceLimitConfig       `json:"limit,omitempty"`
	LimitOverrides []GlobalRateLimitOverrideConfig `json:"limit_overrides,omitempty"`
}

type RateLimitServiceConfig struct {
	Host    string    `json:"host,omitempty"`
	Port    int32     `json:"port,omitempty"`
	Timeout *Duration `json:"timeout,omitempty"`
}

// ASMGlobalRateLimiterSpec defines the desired state of ASMGlobalRateLimiter
type ASMGlobalRateLimiterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	WorkloadSelector *WorkloadSelector `json:"workloadSelector,omitempty"`

	RateLimitService *RateLimitServiceConfig    `json:"rateLimitService,omitempty"`
	Configs          []*GlobalRateLimiterConfig `json:"configs,omitempty"`
	IsGateway        bool                       `json:"isGateway,omitempty"`
}

// ASMGlobalRateLimiterStatus defines the observed state of ASMGlobalRateLimiter
type ASMGlobalRateLimiterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
	// Global Rate Limit Service Config YAML
	GlobalRateLimiterServiceConfig string `json:"config.yaml,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ASMGlobalRateLimiter is the Schema for the asmglobalratelimiters API
// +genclient
type ASMGlobalRateLimiter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMGlobalRateLimiterSpec   `json:"spec,omitempty"`
	Status ASMGlobalRateLimiterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ASMGlobalRateLimiterList contains a list of ASMGlobalRateLimiter
type ASMGlobalRateLimiterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMGlobalRateLimiter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMGlobalRateLimiter{}, &ASMGlobalRateLimiterList{})
}
