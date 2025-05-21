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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// `WorkloadSelector` specifies the criteria used to determine if the
// `Gateway`, `Sidecar`, `EnvoyFilter`, or `ServiceEntry`, or `ASMLocalRateLimiter`
// configuration can be applied to a proxy. The matching criteria
// includes the metadata associated with a proxy, workload instance
// info such as labels attached to the pod/VM, or any other info that
// the proxy provides to Istio during the initial handshake. If
// multiple conditions are specified, all conditions need to match in
// order for the workload instance to be selected. Currently, only
// label based selection mechanism is supported.
type WorkloadSelector struct {
	// One or more labels that indicate a specific set of pods/VMs
	// on which the configuration should be applied. The scope of
	// label search is restricted to the configuration namespace in which the
	// the resource is present.
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

type HeaderMatcher struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If specified, this regex string is a regular expression rule which implies the entire request
	// header value must match the regex. The rule will not match if only a subsequence of the
	// request header value matches the regex.
	RegexMatch string `protobuf:"bytes,2,opt,name=regex_match,json=regexMatch,proto3" json:"regex_match,omitempty"`
	// If specified, header match will be performed based on the value of the header.
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3" json:"exact_match,omitempty"`
	// * The prefix *abcd* matches the value *abcdxyz*, but not for *abcxyz*.
	PrefixMatch string `protobuf:"bytes,4,opt,name=prefix_match,json=prefixMatch,proto3" json:"prefix_match,omitempty"`
	// * The suffix *abcd* matches the value *xyzabcd*, but not for *xyzbcd*.
	SuffixMatch string `protobuf:"bytes,5,opt,name=suffix_match,json=suffixMatch,proto3" json:"suffix_match,omitempty"`
	// If specified as true, header match will be performed based on whether the header is in the
	// request. If specified as false, header match will be performed based on whether the header is absent.
	PresentMatch bool `protobuf:"varint,6,opt,name=present_match,json=presentMatch,proto3" json:"present_match,omitempty"`
	// If specified, the match result will be inverted before checking. Defaults to false.
	// * The regex ``\d{3}`` does not match the value *1234*, so it will match when inverted.
	InvertMatch bool `protobuf:"varint,7,opt,name=invert_match,json=invertMatch,proto3" json:"invert_match,omitempty"`
	// if specified, the exact match the value ""
	IsExactMatchEmpty    bool     `protobuf:"varint,8,opt,name=is_exact_match_empty,json=isExactMatchEmpty,proto3" json:"is_exact_match_empty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type RouteMatch struct {
	NameMatch   string           `protobuf:"bytes,1,opt,name=name_match,proto3" json:"name_match,omitempty"`
	HeaderMatch []*HeaderMatcher `protobuf:"bytes,2,rep,name=header_match,proto3" json:"header_match,omitempty"`
}

type VirtualHostMatch struct {
	Name  string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port  int32      `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Route RouteMatch `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
}

type RateLimitMatch struct {
	VHost *VirtualHostMatch `protobuf:"bytes,1,opt,name=vhost,proto3" json:"vhost,omitempty"`
}

type Duration struct {
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive. Note: these bounds are computed from:
	// 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Envoy Issue: https://github.com/envoyproxy/envoy/issues/21513
type LimitConfig struct {
	FillInterval            *Duration `protobuf:"bytes,1,opt,name=fill_interval,proto3" json:"fill_interval,omitempty"`
	Quota                   int32     `protobuf:"varint,2,opt,name=quota,proto3" json:"quota,omitempty"`
	PerDownStreamConnection bool      `json:"per_downstream_connection,omitempty"`
	//+kubebuilder:validation:MinLength=1
	CustomResponseBody string `json:"custom_response_body,omitempty"`
	//+kubebuilder:validation:MinProperties=1
	//Deprecated,chaned to ResponseHeadersToAdd
	ResponseHeaderToAdd map[string]string `json:"response_header_to_add,omitempty"`
	//+kubebuilder:validation:MinProperties=1
	ResponseHeadersToAdd map[string]string `json:"response_headers_to_add,omitempty"`
}

type LocalRateLimiterConfig struct {
	Name  string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Match *RateLimitMatch `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`
	Limit *LimitConfig    `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`

	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// ASMLocalRateLimiterSpec defines the desired state of ASMLocalRateLimiter
type ASMLocalRateLimiterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	WorkloadSelector *WorkloadSelector `protobuf:"bytes,3,opt,name=workload_selector,json=workloadSelector,proto3" json:"workloadSelector,omitempty"`

	Configs []*LocalRateLimiterConfig `protobuf:"bytes,4,opt,name=configs,json=configs,proto3" json:"configs,omitempty"`

	IsGateway bool `protobuf:"varint,5,opt,name=is_gateway,json=isGateway,proto3" json:"isGateway,omitempty"`
}

// ASMLocalRateLimiterStatus defines the observed state of ASMLocalRateLimiter
type ASMLocalRateLimiterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
}

//+kubebuilder:storageversion
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ASMLocalRateLimiter is the Schema for the asmlocalratelimiters API
type ASMLocalRateLimiter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMLocalRateLimiterSpec   `json:"spec,omitempty"`
	Status ASMLocalRateLimiterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ASMLocalRateLimiterList contains a list of ASMLocalRateLimiter
type ASMLocalRateLimiterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMLocalRateLimiter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMLocalRateLimiter{}, &ASMLocalRateLimiterList{})
}
