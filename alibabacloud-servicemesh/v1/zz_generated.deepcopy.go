//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMAdaptiveConcurrency) DeepCopyInto(out *ASMAdaptiveConcurrency) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMAdaptiveConcurrency.
func (in *ASMAdaptiveConcurrency) DeepCopy() *ASMAdaptiveConcurrency {
	if in == nil {
		return nil
	}
	out := new(ASMAdaptiveConcurrency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMAdaptiveConcurrency) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMAdaptiveConcurrencyList) DeepCopyInto(out *ASMAdaptiveConcurrencyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ASMAdaptiveConcurrency, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMAdaptiveConcurrencyList.
func (in *ASMAdaptiveConcurrencyList) DeepCopy() *ASMAdaptiveConcurrencyList {
	if in == nil {
		return nil
	}
	out := new(ASMAdaptiveConcurrencyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMAdaptiveConcurrencyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMAdaptiveConcurrencySpec) DeepCopyInto(out *ASMAdaptiveConcurrencySpec) {
	*out = *in
	if in.WorkloadSelector != nil {
		in, out := &in.WorkloadSelector, &out.WorkloadSelector
		*out = new(WorkloadSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SampleAggregatePercentile != nil {
		in, out := &in.SampleAggregatePercentile, &out.SampleAggregatePercentile
		*out = new(Percent)
		**out = **in
	}
	in.ConcurrencyLimitParams.DeepCopyInto(&out.ConcurrencyLimitParams)
	in.MinRTTCalcParams.DeepCopyInto(&out.MinRTTCalcParams)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMAdaptiveConcurrencySpec.
func (in *ASMAdaptiveConcurrencySpec) DeepCopy() *ASMAdaptiveConcurrencySpec {
	if in == nil {
		return nil
	}
	out := new(ASMAdaptiveConcurrencySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMAdaptiveConcurrencyStatus) DeepCopyInto(out *ASMAdaptiveConcurrencyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMAdaptiveConcurrencyStatus.
func (in *ASMAdaptiveConcurrencyStatus) DeepCopy() *ASMAdaptiveConcurrencyStatus {
	if in == nil {
		return nil
	}
	out := new(ASMAdaptiveConcurrencyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMCircuitBreaker) DeepCopyInto(out *ASMCircuitBreaker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMCircuitBreaker.
func (in *ASMCircuitBreaker) DeepCopy() *ASMCircuitBreaker {
	if in == nil {
		return nil
	}
	out := new(ASMCircuitBreaker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMCircuitBreaker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMCircuitBreakerList) DeepCopyInto(out *ASMCircuitBreakerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ASMCircuitBreaker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMCircuitBreakerList.
func (in *ASMCircuitBreakerList) DeepCopy() *ASMCircuitBreakerList {
	if in == nil {
		return nil
	}
	out := new(ASMCircuitBreakerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMCircuitBreakerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMCircuitBreakerSpec) DeepCopyInto(out *ASMCircuitBreakerSpec) {
	*out = *in
	if in.WorkloadSelector != nil {
		in, out := &in.WorkloadSelector, &out.WorkloadSelector
		*out = new(WorkloadSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make([]*CircuitBreakerConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CircuitBreakerConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMCircuitBreakerSpec.
func (in *ASMCircuitBreakerSpec) DeepCopy() *ASMCircuitBreakerSpec {
	if in == nil {
		return nil
	}
	out := new(ASMCircuitBreakerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMCircuitBreakerStatus) DeepCopyInto(out *ASMCircuitBreakerStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMCircuitBreakerStatus.
func (in *ASMCircuitBreakerStatus) DeepCopy() *ASMCircuitBreakerStatus {
	if in == nil {
		return nil
	}
	out := new(ASMCircuitBreakerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMEgressTrafficPolicy) DeepCopyInto(out *ASMEgressTrafficPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMEgressTrafficPolicy.
func (in *ASMEgressTrafficPolicy) DeepCopy() *ASMEgressTrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(ASMEgressTrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMEgressTrafficPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMEgressTrafficPolicyList) DeepCopyInto(out *ASMEgressTrafficPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ASMEgressTrafficPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMEgressTrafficPolicyList.
func (in *ASMEgressTrafficPolicyList) DeepCopy() *ASMEgressTrafficPolicyList {
	if in == nil {
		return nil
	}
	out := new(ASMEgressTrafficPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMEgressTrafficPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMLocalRateLimiter) DeepCopyInto(out *ASMLocalRateLimiter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMLocalRateLimiter.
func (in *ASMLocalRateLimiter) DeepCopy() *ASMLocalRateLimiter {
	if in == nil {
		return nil
	}
	out := new(ASMLocalRateLimiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMLocalRateLimiter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMLocalRateLimiterList) DeepCopyInto(out *ASMLocalRateLimiterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ASMLocalRateLimiter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMLocalRateLimiterList.
func (in *ASMLocalRateLimiterList) DeepCopy() *ASMLocalRateLimiterList {
	if in == nil {
		return nil
	}
	out := new(ASMLocalRateLimiterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASMLocalRateLimiterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMLocalRateLimiterSpec) DeepCopyInto(out *ASMLocalRateLimiterSpec) {
	*out = *in
	if in.WorkloadSelector != nil {
		in, out := &in.WorkloadSelector, &out.WorkloadSelector
		*out = new(WorkloadSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make([]*LocalRateLimiterConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LocalRateLimiterConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMLocalRateLimiterSpec.
func (in *ASMLocalRateLimiterSpec) DeepCopy() *ASMLocalRateLimiterSpec {
	if in == nil {
		return nil
	}
	out := new(ASMLocalRateLimiterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASMLocalRateLimiterStatus) DeepCopyInto(out *ASMLocalRateLimiterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASMLocalRateLimiterStatus.
func (in *ASMLocalRateLimiterStatus) DeepCopy() *ASMLocalRateLimiterStatus {
	if in == nil {
		return nil
	}
	out := new(ASMLocalRateLimiterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakerConfig) DeepCopyInto(out *BreakerConfig) {
	*out = *in
	if in.AverageRequestRT != nil {
		in, out := &in.AverageRequestRT, &out.AverageRequestRT
		*out = new(string)
		**out = **in
	}
	if in.SlowRequestRT != nil {
		in, out := &in.SlowRequestRT, &out.SlowRequestRT
		*out = new(string)
		**out = **in
	}
	if in.MaxSlowRequests != nil {
		in, out := &in.MaxSlowRequests, &out.MaxSlowRequests
		*out = new(uint32)
		**out = **in
	}
	if in.ErrorPercent != nil {
		in, out := &in.ErrorPercent, &out.ErrorPercent
		*out = new(Percent)
		**out = **in
	}
	if in.MinRequestAmount != nil {
		in, out := &in.MinRequestAmount, &out.MinRequestAmount
		*out = new(uint32)
		**out = **in
	}
	if in.CustomResponse != nil {
		in, out := &in.CustomResponse, &out.CustomResponse
		*out = new(CustomResponse)
		(*in).DeepCopyInto(*out)
	}
	if in.BreakDuration != nil {
		in, out := &in.BreakDuration, &out.BreakDuration
		*out = new(string)
		**out = **in
	}
	if in.WindowSize != nil {
		in, out := &in.WindowSize, &out.WindowSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakerConfig.
func (in *BreakerConfig) DeepCopy() *BreakerConfig {
	if in == nil {
		return nil
	}
	out := new(BreakerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ByEgressGateway) DeepCopyInto(out *ByEgressGateway) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByEgressGateway.
func (in *ByEgressGateway) DeepCopy() *ByEgressGateway {
	if in == nil {
		return nil
	}
	out := new(ByEgressGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitBreakerConfig) DeepCopyInto(out *CircuitBreakerConfig) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(CircuitBreakerMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Breaker != nil {
		in, out := &in.Breaker, &out.Breaker
		*out = new(BreakerConfig)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitBreakerConfig.
func (in *CircuitBreakerConfig) DeepCopy() *CircuitBreakerConfig {
	if in == nil {
		return nil
	}
	out := new(CircuitBreakerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitBreakerMatch) DeepCopyInto(out *CircuitBreakerMatch) {
	*out = *in
	if in.VHost != nil {
		in, out := &in.VHost, &out.VHost
		*out = new(VirtualHostMatch)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitBreakerMatch.
func (in *CircuitBreakerMatch) DeepCopy() *CircuitBreakerMatch {
	if in == nil {
		return nil
	}
	out := new(CircuitBreakerMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConcurrencyLimitParamsConfiguration) DeepCopyInto(out *ConcurrencyLimitParamsConfiguration) {
	*out = *in
	if in.MaxConcurrencyLimit != nil {
		in, out := &in.MaxConcurrencyLimit, &out.MaxConcurrencyLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConcurrencyLimitParamsConfiguration.
func (in *ConcurrencyLimitParamsConfiguration) DeepCopy() *ConcurrencyLimitParamsConfiguration {
	if in == nil {
		return nil
	}
	out := new(ConcurrencyLimitParamsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResponse) DeepCopyInto(out *CustomResponse) {
	*out = *in
	if in.StatusCode != nil {
		in, out := &in.StatusCode, &out.StatusCode
		*out = new(int32)
		**out = **in
	}
	if in.ResponseHeaderToAdd != nil {
		in, out := &in.ResponseHeaderToAdd, &out.ResponseHeaderToAdd
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResponse.
func (in *CustomResponse) DeepCopy() *CustomResponse {
	if in == nil {
		return nil
	}
	out := new(CustomResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Duration) DeepCopyInto(out *Duration) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressRule) DeepCopyInto(out *EgressRule) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]*From, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(From)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]*To, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(To)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressRule.
func (in *EgressRule) DeepCopy() *EgressRule {
	if in == nil {
		return nil
	}
	out := new(EgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressTrafficPolicySpec) DeepCopyInto(out *EgressTrafficPolicySpec) {
	*out = *in
	if in.ByEgressGateway != nil {
		in, out := &in.ByEgressGateway, &out.ByEgressGateway
		*out = new(ByEgressGateway)
		**out = **in
	}
	if in.EgressRules != nil {
		in, out := &in.EgressRules, &out.EgressRules
		*out = make([]*EgressRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EgressRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressTrafficPolicySpec.
func (in *EgressTrafficPolicySpec) DeepCopy() *EgressTrafficPolicySpec {
	if in == nil {
		return nil
	}
	out := new(EgressTrafficPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressTrafficPolicyStatus) DeepCopyInto(out *EgressTrafficPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressTrafficPolicyStatus.
func (in *EgressTrafficPolicyStatus) DeepCopy() *EgressTrafficPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(EgressTrafficPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *From) DeepCopyInto(out *From) {
	*out = *in
	if in.WorkLoadSelector != nil {
		in, out := &in.WorkLoadSelector, &out.WorkLoadSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new From.
func (in *From) DeepCopy() *From {
	if in == nil {
		return nil
	}
	out := new(From)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatcher) DeepCopyInto(out *HeaderMatcher) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatcher.
func (in *HeaderMatcher) DeepCopy() *HeaderMatcher {
	if in == nil {
		return nil
	}
	out := new(HeaderMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpsUpgrade) DeepCopyInto(out *HttpsUpgrade) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpsUpgrade.
func (in *HttpsUpgrade) DeepCopy() *HttpsUpgrade {
	if in == nil {
		return nil
	}
	out := new(HttpsUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitConfig) DeepCopyInto(out *LimitConfig) {
	*out = *in
	if in.FillInterval != nil {
		in, out := &in.FillInterval, &out.FillInterval
		*out = new(Duration)
		(*in).DeepCopyInto(*out)
	}
	if in.ResponseHeaderToAdd != nil {
		in, out := &in.ResponseHeaderToAdd, &out.ResponseHeaderToAdd
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResponseHeadersToAdd != nil {
		in, out := &in.ResponseHeadersToAdd, &out.ResponseHeadersToAdd
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitConfig.
func (in *LimitConfig) DeepCopy() *LimitConfig {
	if in == nil {
		return nil
	}
	out := new(LimitConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalRateLimiterConfig) DeepCopyInto(out *LocalRateLimiterConfig) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(RateLimitMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = new(LimitConfig)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimiterConfig.
func (in *LocalRateLimiterConfig) DeepCopy() *LocalRateLimiterConfig {
	if in == nil {
		return nil
	}
	out := new(LocalRateLimiterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinRTTCalcParamsConfiguration) DeepCopyInto(out *MinRTTCalcParamsConfiguration) {
	*out = *in
	if in.RequestCount != nil {
		in, out := &in.RequestCount, &out.RequestCount
		*out = new(int32)
		**out = **in
	}
	if in.Jitter != nil {
		in, out := &in.Jitter, &out.Jitter
		*out = new(Percent)
		**out = **in
	}
	if in.MinConcurrency != nil {
		in, out := &in.MinConcurrency, &out.MinConcurrency
		*out = new(int32)
		**out = **in
	}
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(Percent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinRTTCalcParamsConfiguration.
func (in *MinRTTCalcParamsConfiguration) DeepCopy() *MinRTTCalcParamsConfiguration {
	if in == nil {
		return nil
	}
	out := new(MinRTTCalcParamsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Percent) DeepCopyInto(out *Percent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Percent.
func (in *Percent) DeepCopy() *Percent {
	if in == nil {
		return nil
	}
	out := new(Percent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitMatch) DeepCopyInto(out *RateLimitMatch) {
	*out = *in
	if in.VHost != nil {
		in, out := &in.VHost, &out.VHost
		*out = new(VirtualHostMatch)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitMatch.
func (in *RateLimitMatch) DeepCopy() *RateLimitMatch {
	if in == nil {
		return nil
	}
	out := new(RateLimitMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMatch) DeepCopyInto(out *RouteMatch) {
	*out = *in
	if in.HeaderMatch != nil {
		in, out := &in.HeaderMatch, &out.HeaderMatch
		*out = make([]*HeaderMatcher, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HeaderMatcher)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMatch.
func (in *RouteMatch) DeepCopy() *RouteMatch {
	if in == nil {
		return nil
	}
	out := new(RouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *To) DeepCopyInto(out *To) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = (*in).DeepCopy()
	}
	if in.ByEgressGateway != nil {
		in, out := &in.ByEgressGateway, &out.ByEgressGateway
		*out = new(ByEgressGateway)
		**out = **in
	}
	if in.HttpsUpgrade != nil {
		in, out := &in.HttpsUpgrade, &out.HttpsUpgrade
		*out = new(HttpsUpgrade)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new To.
func (in *To) DeepCopy() *To {
	if in == nil {
		return nil
	}
	out := new(To)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHostMatch) DeepCopyInto(out *VirtualHostMatch) {
	*out = *in
	in.Route.DeepCopyInto(&out.Route)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHostMatch.
func (in *VirtualHostMatch) DeepCopy() *VirtualHostMatch {
	if in == nil {
		return nil
	}
	out := new(VirtualHostMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSelector) DeepCopyInto(out *WorkloadSelector) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector.
func (in *WorkloadSelector) DeepCopy() *WorkloadSelector {
	if in == nil {
		return nil
	}
	out := new(WorkloadSelector)
	in.DeepCopyInto(out)
	return out
}
