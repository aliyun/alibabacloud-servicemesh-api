// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using StringMatch within kubernetes types, where deepcopy-gen is used.
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	p := proto.Clone(in).(*StringMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MutualTls within kubernetes types, where deepcopy-gen is used.
func (in *MutualTls) DeepCopyInto(out *MutualTls) {
	p := proto.Clone(in).(*MutualTls)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutualTls. Required by controller-gen.
func (in *MutualTls) DeepCopy() *MutualTls {
	if in == nil {
		return nil
	}
	out := new(MutualTls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MutualTls. Required by controller-gen.
func (in *MutualTls) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Jwt within kubernetes types, where deepcopy-gen is used.
func (in *Jwt) DeepCopyInto(out *Jwt) {
	p := proto.Clone(in).(*Jwt)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jwt. Required by controller-gen.
func (in *Jwt) DeepCopy() *Jwt {
	if in == nil {
		return nil
	}
	out := new(Jwt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Jwt. Required by controller-gen.
func (in *Jwt) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Jwt_TriggerRule within kubernetes types, where deepcopy-gen is used.
func (in *Jwt_TriggerRule) DeepCopyInto(out *Jwt_TriggerRule) {
	p := proto.Clone(in).(*Jwt_TriggerRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jwt_TriggerRule. Required by controller-gen.
func (in *Jwt_TriggerRule) DeepCopy() *Jwt_TriggerRule {
	if in == nil {
		return nil
	}
	out := new(Jwt_TriggerRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Jwt_TriggerRule. Required by controller-gen.
func (in *Jwt_TriggerRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PeerAuthenticationMethod within kubernetes types, where deepcopy-gen is used.
func (in *PeerAuthenticationMethod) DeepCopyInto(out *PeerAuthenticationMethod) {
	p := proto.Clone(in).(*PeerAuthenticationMethod)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthenticationMethod. Required by controller-gen.
func (in *PeerAuthenticationMethod) DeepCopy() *PeerAuthenticationMethod {
	if in == nil {
		return nil
	}
	out := new(PeerAuthenticationMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthenticationMethod. Required by controller-gen.
func (in *PeerAuthenticationMethod) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using OriginAuthenticationMethod within kubernetes types, where deepcopy-gen is used.
func (in *OriginAuthenticationMethod) DeepCopyInto(out *OriginAuthenticationMethod) {
	p := proto.Clone(in).(*OriginAuthenticationMethod)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginAuthenticationMethod. Required by controller-gen.
func (in *OriginAuthenticationMethod) DeepCopy() *OriginAuthenticationMethod {
	if in == nil {
		return nil
	}
	out := new(OriginAuthenticationMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new OriginAuthenticationMethod. Required by controller-gen.
func (in *OriginAuthenticationMethod) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Policy within kubernetes types, where deepcopy-gen is used.
func (in *Policy) DeepCopyInto(out *Policy) {
	p := proto.Clone(in).(*Policy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy. Required by controller-gen.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Policy. Required by controller-gen.
func (in *Policy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TargetSelector within kubernetes types, where deepcopy-gen is used.
func (in *TargetSelector) DeepCopyInto(out *TargetSelector) {
	p := proto.Clone(in).(*TargetSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSelector. Required by controller-gen.
func (in *TargetSelector) DeepCopy() *TargetSelector {
	if in == nil {
		return nil
	}
	out := new(TargetSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TargetSelector. Required by controller-gen.
func (in *TargetSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PortSelector within kubernetes types, where deepcopy-gen is used.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	p := proto.Clone(in).(*PortSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
