// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using TrafficLabel within kubernetes types, where deepcopy-gen is used.
func (in *TrafficLabel) DeepCopyInto(out *TrafficLabel) {
	p := proto.Clone(in).(*TrafficLabel)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficLabel. Required by controller-gen.
func (in *TrafficLabel) DeepCopy() *TrafficLabel {
	if in == nil {
		return nil
	}
	out := new(TrafficLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficLabel. Required by controller-gen.
func (in *TrafficLabel) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficLabelRule within kubernetes types, where deepcopy-gen is used.
func (in *TrafficLabelRule) DeepCopyInto(out *TrafficLabelRule) {
	p := proto.Clone(in).(*TrafficLabelRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficLabelRule. Required by controller-gen.
func (in *TrafficLabelRule) DeepCopy() *TrafficLabelRule {
	if in == nil {
		return nil
	}
	out := new(TrafficLabelRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficLabelRule. Required by controller-gen.
func (in *TrafficLabelRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LabelVar within kubernetes types, where deepcopy-gen is used.
func (in *LabelVar) DeepCopyInto(out *LabelVar) {
	p := proto.Clone(in).(*LabelVar)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelVar. Required by controller-gen.
func (in *LabelVar) DeepCopy() *LabelVar {
	if in == nil {
		return nil
	}
	out := new(LabelVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LabelVar. Required by controller-gen.
func (in *LabelVar) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPLabelRule within kubernetes types, where deepcopy-gen is used.
func (in *HTTPLabelRule) DeepCopyInto(out *HTTPLabelRule) {
	p := proto.Clone(in).(*HTTPLabelRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPLabelRule. Required by controller-gen.
func (in *HTTPLabelRule) DeepCopy() *HTTPLabelRule {
	if in == nil {
		return nil
	}
	out := new(HTTPLabelRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPLabelRule. Required by controller-gen.
func (in *HTTPLabelRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
