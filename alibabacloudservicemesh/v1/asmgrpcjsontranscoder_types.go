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

type GrpcJsonTranscoderConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type GrpcJsonTranscoder_PrintOptions struct {

	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace *bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"addWhitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields *bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"alwaysPrintPrimitiveFields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts *bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"alwaysPrintEnumsAsInts,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames *bool `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserveProtoFieldNames,omitempty"`
}

// ASMGrpcJsonTranscoderSpec defines the desired state of ASMGrpcJsonTranscoder
type ASMGrpcJsonTranscoderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	WorkloadSelector *WorkloadSelector `protobuf:"bytes,3,opt,name=workload_selector,json=workloadSelector,proto3" json:"workloadSelector,omitempty"`
	IsGateway        bool              `protobuf:"varint,4,opt,name=is_gateway,json=isGateway,proto3" json:"isGateway,omitempty"`

	// Priority defines the order in which patch sets are applied within a context.
	// When one patch depends on another patch, the order of patch application
	// is significant. The API provides two primary ways to order patches.
	// Patch sets in the root namespace are applied before the patch sets in the
	// workload namespace. Patches within a patch set are processed in the order
	// that they appear in the `configPatches` list.
	//
	// The default value for priority is 0 and the range is [ min-int32, max-int32 ].
	// A patch set with a negative priority is processed before the default. A patch
	// set with a positive priority is processed after the default.
	//
	// It is recommended to start with priority values that are multiples of 10
	// to leave room for further insertion.
	//
	// Patch sets are sorted in the following ascending key order:
	// priority, creation time, fully qualified resource name.
	Priority int32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`

	//Configs *GrpcJsonTranscoderConfig `protobuf:"bytes,6,opt,name=configs,json=configs,proto3" json:"configs,omitempty"`

	// The service port/gateway port to which traffic is being
	// sent/received. If not specified, matches all listeners. Even though
	// inbound listeners are generated for the instance/pod ports, only
	// service ports should be used to match listeners.
	PortNumber uint32 `protobuf:"varint,6,opt,name=port_number,json=portNumber,proto3" json:"portNumber,omitempty"`

	// Supplies the base64 encoded string content of
	// :ref:`the proto descriptor set <config_grpc_json_generate_proto_descriptor_set>` for the gRPC
	// services.
	ProtoDescriptorBin string `protobuf:"bytes,7,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3" json:"protoDescriptorBin,omitempty"`

	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,8,rep,name=services,proto3" json:"services,omitempty"`

	// Builtin Proto Descriptor, e.g. the value "KServe_Predict_V2" means the proto descriptor
	// for KServe predict v2 grpc;
	BuiltinProtoDescriptor BuiltinProtoDescriptor `protobuf:"varint,9,opt,name=builtin_proto_descriptor,enum=BuiltinProtoDescriptor,json=builtinProtoDescriptor,proto3" json:"builtinProtoDescriptor,omitempty"`

	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,10,opt,name=print_options,json=printOptions,proto3" json:"printOptions,omitempty"`

	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute *bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"matchIncomingRequestRoute,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//	service Bookstore {
	//	  rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//	    option (google.api.http) = {
	//	      get: "/shelves/{shelf}"
	//	    };
	//	  }
	//	}
	//
	//	message GetShelfRequest {
	//	  int64 shelf = 1;
	//	}
	//
	//	message Shelf {}
	//
	// The request “/shelves/100?foo=bar“ will not be mapped to “GetShelf``` because variable
	// binding for “foo“ is not defined. Adding “foo“ to “ignored_query_parameters“ will allow
	// the same request to be mapped to “GetShelf“.
	IgnoredQueryParameters *[]string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignoredQueryParameters,omitempty"`
	// Whether to route methods without the “google.api.http“ option.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//	package bookstore;
	//
	//	service Bookstore {
	//	  rpc GetShelf(GetShelfRequest) returns (Shelf) {}
	//	}
	//
	//	message GetShelfRequest {
	//	  int64 shelf = 1;
	//	}
	//
	//	message Shelf {}
	//
	// The client could “post“ a json body “{"shelf": 1234}“ with the path of
	// “/bookstore.Bookstore/GetShelfRequest“ to call “GetShelfRequest“.
	AutoMapping *bool `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"autoMapping,omitempty"`
	// Whether to ignore query parameters that cannot be mapped to a corresponding
	// protobuf field. Use this if you cannot control the query parameters and do
	// not know them beforehand. Otherwise use “ignored_query_parameters“.
	// Defaults to false.
	IgnoreUnknownQueryParameters *bool `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignoreUnknownQueryParameters,omitempty"`
	// Whether to convert gRPC status headers to JSON.
	// When trailer indicates a gRPC error and there was no HTTP body, take “google.rpc.Status“
	// from the “grpc-status-details-bin“ header and use it as JSON body.
	// If there was no such header, make “google.rpc.Status“ out of the “grpc-status“ and
	// “grpc-message“ headers.
	// The error details types must be present in the “proto_descriptor“.
	//
	// For example, if an upstream server replies with headers:
	//
	// .. code-block:: none
	//
	//	grpc-status: 5
	//	grpc-status-details-bin:
	//	    CAUaMwoqdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLlJlcXVlc3RJbmZvEgUKA3ItMQ
	//
	// The “grpc-status-details-bin“ header contains a base64-encoded protobuf message
	// “google.rpc.Status“. It will be transcoded into:
	//
	// .. code-block:: none
	//
	//	HTTP/1.1 404 Not Found
	//	content-type: application/json
	//
	//	{"code":5,"details":[{"@type":"type.googleapis.com/google.rpc.RequestInfo","requestId":"r-1"}]}
	//
	// In order to transcode the message, the “google.rpc.RequestInfo“ type from
	// the “google/rpc/error_details.proto“ should be included in the configured
	// :ref:`proto descriptor set <config_grpc_json_generate_proto_descriptor_set>`.
	ConvertGrpcStatus *bool `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convertGrpcStatus,omitempty"`

	TranscodeFirst bool `protobuf:"bytes,11,opt,name=transcode_first,json=transcodeFirst,proto3" json:"transcodeFirst,omitempty"`
}

type BuiltinProtoDescriptor string

const (
	BuiltinProtoDescriptor_KServe_Predict_V2 BuiltinProtoDescriptor = "KServe_Predict_V2"
)

var BuiltinProtoDescriptor_name = map[int32]BuiltinProtoDescriptor{
	0: BuiltinProtoDescriptor_KServe_Predict_V2,
}
var BuiltinProtoDescriptor_value = map[BuiltinProtoDescriptor]int32{
	BuiltinProtoDescriptor_KServe_Predict_V2: 0,
}

// ASMGrpcJsonTranscoderStatus defines the observed state of ASMGrpcJsonTranscoder
type ASMGrpcJsonTranscoderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ASMGrpcJsonTranscoder is the Schema for the ASMGrpcJsonTranscoders API
// +genclient
type ASMGrpcJsonTranscoder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMGrpcJsonTranscoderSpec   `json:"spec,omitempty"`
	Status ASMGrpcJsonTranscoderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ASMGrpcJsonTranscoderList contains a list of ASMGrpcJsonTranscoder
type ASMGrpcJsonTranscoderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMGrpcJsonTranscoder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMGrpcJsonTranscoder{}, &ASMGrpcJsonTranscoderList{})
}
