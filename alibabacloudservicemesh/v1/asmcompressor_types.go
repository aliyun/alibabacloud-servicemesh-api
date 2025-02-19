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

// refer to https://github.com/envoyproxy/go-control-plane/blob/84c5a14eb8c8b5b99efeecd3801fea84d628852d/envoy/extensions/filters/http/compressor/v3/compressor.pb.go

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ASMCompressorSpec defines the desired state of ASMCompressor
type ASMCompressorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	WorkloadSelector *WorkloadSelector `protobuf:"bytes,1,opt,name=workload_selector,json=workloadSelector,proto3" json:"workloadSelector,omitempty"`
	IsGateway        bool              `protobuf:"varint,2,opt,name=is_gateway,json=isGateway,proto3" json:"isGateway,omitempty"`

	// The service port/gateway port to which traffic is being
	// sent/received. If not specified, matches all listeners. Even though
	// inbound listeners are generated for the instance/pod ports, only
	// service ports should be used to match listeners.
	PortNumber uint32 `protobuf:"varint,3,opt,name=port_number,json=portNumber,proto3" json:"portNumber,omitempty"`

	// A compressor library to use for compression. Currently only
	// :ref:`envoy.compression.gzip.compressor<envoy_v3_api_msg_extensions.compression.gzip.compressor.v3.Gzip>`
	// is included in Envoy.
	// [#extension-category: envoy.compression.compressor]
	CompressorLibrary *CompressorLibrary `protobuf:"bytes,6,opt,name=compressor_library,json=compressorLibrary,proto3" json:"compressor_library,omitempty"`
	// Configuration for request compression. Compression is disabled by default if left empty.
	RequestDirectionConfig *Compressor_RequestDirectionConfig `protobuf:"bytes,7,opt,name=request_direction_config,json=requestDirectionConfig,proto3" json:"request_direction_config,omitempty"`
	// Configuration for response compression. Compression is enabled by default if left empty.
	//
	// .. attention::
	//
	//    If the field is not empty then the duplicate deprecated fields of the ``Compressor`` message,
	//    such as ``content_length``, ``content_type``, ``disable_on_etag_header``,
	//    ``remove_accept_encoding_header`` and ``runtime_enabled``, are ignored.
	//
	//    Also all the statistics related to response compression will be rooted in
	//    ``<stat_prefix>.compressor.<compressor_library.name>.<compressor_library_stat_prefix>.response.*``
	//    instead of
	//    ``<stat_prefix>.compressor.<compressor_library.name>.<compressor_library_stat_prefix>.*``.
	ResponseDirectionConfig *Compressor_ResponseDirectionConfig `protobuf:"bytes,8,opt,name=response_direction_config,json=responseDirectionConfig,proto3" json:"response_direction_config,omitempty"`
	// If true, chooses this compressor first to do compression when the q-values in `Accept-Encoding` are same.
	// The last compressor which enables choose_first will be chosen if multiple compressor filters in the chain have choose_first as true.
	ChooseFirst     *bool                         `protobuf:"varint,9,opt,name=choose_first,json=chooseFirst,proto3" json:"choose_first,omitempty"`
	PerRouteConfigs []ASMCompressorPerRouteConfig `protobuf:"bytes,11,rep,name=per_route_configs,json=perRouteConfigs,proto3" json:"per_route_configs,omitempty"`
}

type ASMCompressorPerRouteConfig struct {
	// The route match to which this config applies.
	RouteMatch *RouteConfigurationMatch `protobuf:"bytes,0,opt,name=route_match,json=routeMatch,proto3" json:"route_match,omitempty"`
	// The config to use when a route matches.
	Disabled *bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

type CompressorLibrary struct {
	//Algorithm *string `protobuf:"bytes,1,opt,name=algorithm,json=algorithm,proto3" json:"algorithm,omitempty"`

	Gzip   *GzipCompressor   `protobuf:"bytes,2,opt,name=gzip,json=gzip,proto3,oneof" json:"gzip,omitempty"`
	Brotli *BrotliCompressor `protobuf:"bytes,3,opt,name=brotli,json=brotli,proto3,oneof" json:"brotli,omitempty"`
	Zstd   *ZstdCompressor   `protobuf:"bytes,4,opt,name=zstd,json=zstd,proto3,oneof" json:"zstd,omitempty"`
}

// Configuration for filter behavior on the request direction.
type Compressor_RequestDirectionConfig struct {
	CommonConfig *Compressor_CommonDirectionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
}

// Configuration for filter behavior on the response direction.
type Compressor_ResponseDirectionConfig struct {
	CommonConfig *Compressor_CommonDirectionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// If true, disables compression when the response contains an etag header. When it is false, the
	// filter will preserve weak etags and remove the ones that require strong validation.
	DisableOnEtagHeader *bool `protobuf:"varint,2,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	// If true, removes accept-encoding from the request headers before dispatching it to the upstream
	// so that responses do not get compressed before reaching the filter.
	//
	// .. attention::
	//
	//    To avoid interfering with other compression filters in the same chain use this option in
	//    the filter closest to the upstream.
	RemoveAcceptEncodingHeader *bool `protobuf:"varint,3,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
}

type Compressor_CommonDirectionConfig struct {

	// Runtime flag that controls whether compression is enabled or not for the direction this
	// common config is put in. If set to false, the filter will operate as a pass-through filter
	// in the chosen direction, unless overridden by CompressorPerRoute.
	// If the field is omitted, the filter will be enabled.
	Enabled *RuntimeFeatureFlag `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Minimum value of Content-Length header of request or response messages (depending on the direction
	// this common config is put in), in bytes, which will trigger compression. The default value is 30.
	MinContentLength *uint32 `protobuf:"bytes,2,opt,name=min_content_length,json=minContentLength,proto3" json:"min_content_length,omitempty"`
	// Set of strings that allows specifying which mime-types yield compression; e.g.,
	// application/json, text/html, etc. When this field is not defined, compression will be applied
	// to the following mime-types: "application/javascript", "application/json",
	// "application/xhtml+xml", "image/svg+xml", "text/css", "text/html", "text/plain", "text/xml"
	// and their synonyms.
	ContentType []string `protobuf:"bytes,3,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

// Runtime derived bool with a default when not specified.
type RuntimeFeatureFlag struct {

	// Default value if runtime value is not available.
	DefaultValue *bool `protobuf:"bytes,1,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Runtime key to get value for comparison. This value is used if defined. The boolean value must
	// be represented via its
	// `canonical JSON encoding <https://developers.google.com/protocol-buffers/docs/proto3#json>`_.
	RuntimeKey *string `protobuf:"bytes,2,opt,name=runtime_key,json=runtimeKey,proto3" json:"runtime_key,omitempty"`
}

// ASMCompressorStatus defines the observed state of ASMCompressor
type ASMCompressorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ASMCompressor is the Schema for the ASMCompressors API
// +genclient
type ASMCompressor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMCompressorSpec   `json:"spec,omitempty"`
	Status ASMCompressorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ASMCompressorList contains a list of ASMCompressor
type ASMCompressorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMCompressor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMCompressor{}, &ASMCompressorList{})
}
