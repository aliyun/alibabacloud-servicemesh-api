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
	"fmt"
	"math"

	"github.com/gogo/protobuf/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CircuitBreakerMatch struct {
	VHost *VirtualHostMatch `protobuf:"bytes,1,opt,name=vhost,proto3" json:"vhost,omitempty"`
}

// +kubebuilder:validation:MinProperties=1
type CustomResponse struct {
	//+kubebuilder:validation:Minimum=200
	//+kubebuilder:validation:Maximum=599
	StatusCode *int32 `json:"status_code,omitempty"`
	//+kubebuilder:validation:MinProperties=1
	ResponseHeaderToAdd map[string]string `json:"header_to_add,omitempty"`
	//+kubebuilder:validation:MinLength=1
	Body *string `json:"body,omitempty"`
}

// +kubebuilder:validation:MinProperties=1
type BreakerConfig struct {
	//+kubebuilder:validation:Type=string
	AverageRequestRT *string `json:"average_request_rt,omitempty"`
	//+kubebuilder:validation:Type=string
	SlowRequestRT    *string         `json:"slow_request_rt,omitempty"`
	MaxSlowRequests  *uint32         `json:"max_slow_requests,omitempty"`
	ErrorPercent     *Percent        `json:"error_percent,omitempty"`
	MinRequestAmount *uint32         `json:"min_request_amount,omitempty"`
	CustomResponse   *CustomResponse `json:"custom_response,omitempty"`
	//+kubebuilder:validation:Type=string
	BreakDuration *string `json:"break_duration,omitempty"`
	//+kubebuilder:validation:Type=string
	WindowSize *string `json:"window_size,omitempty"`
}

type CircuitBreakerConfig struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//+kubebuilder:validation:Required
	Match *CircuitBreakerMatch `protobuf:"bytes,2,opt,name=match,proto3" json:"match"`
	//+kubebuilder:validation:Required
	Breaker *BreakerConfig `protobuf:"bytes,3,opt,name=breaker_config,proto3" json:"breaker_config"`

	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// ASMCircuitBreakerSpec defines the desired state of ASMCircuitBreaker
type ASMCircuitBreakerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	//+kubebuilder:validation:Optional
	WorkloadSelector *WorkloadSelector `protobuf:"bytes,3,opt,name=workload_selector,json=workloadSelector,proto3" json:"workloadSelector"`
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:MinItems=1
	Configs []*CircuitBreakerConfig `protobuf:"bytes,4,opt,name=configs,json=configs,proto3" json:"configs"`

	IsGateway bool `protobuf:"varint,5,opt,name=is_gateway,json=isGateway,proto3" json:"isGateway,omitempty"`
}

// ASMCircuitBreakerStatus defines the observed state of ASMCircuitBreaker
type ASMCircuitBreakerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message []string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCircuitBreaker is the Schema for the asmcircuitbreaker API
type ASMCircuitBreaker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	//+kubebuilder:validation:Required
	Spec   ASMCircuitBreakerSpec   `json:"spec"`
	Status ASMCircuitBreakerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ASMCircuitBreakerList contains a list of ASMCircuitBreaker
type ASMCircuitBreakerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMCircuitBreaker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMCircuitBreaker{}, &ASMCircuitBreakerList{})
}
