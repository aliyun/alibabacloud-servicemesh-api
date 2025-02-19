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
	"github.com/gogo/protobuf/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math"
)

// refer to https://github.com/envoyproxy/go-control-plane/blob/84c5a14eb8c8b5b99efeecd3801fea84d628852d/envoy/extensions/filters/http/compressor/v3/compressor.pb.go

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ASMDecompressorSpec defines the desired state of ASMDecompressor
type ASMDecompressorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	WorkloadSelector *WorkloadSelector `protobuf:"bytes,1,opt,name=workload_selector,json=workloadSelector,proto3" json:"workloadSelector,omitempty"`
	IsGateway        bool              `protobuf:"varint,2,opt,name=is_gateway,json=isGateway,proto3" json:"isGateway,omitempty"`

	// The service port/gateway port to which traffic is being
	// sent/received. If not specified, matches all listeners. Even though
	// inbound listeners are generated for the instance/pod ports, only
	// service ports should be used to match listeners.
	PortNumber uint32 `protobuf:"varint,3,opt,name=port_number,json=portNumber,proto3" json:"portNumber,omitempty"`

	// A decompressor library to use for both request and response decompression. Currently only
	// :ref:`envoy.compression.gzip.compressor<envoy_v3_api_msg_extensions.compression.gzip.decompressor.v3.Gzip>`
	// is included in Envoy.
	// [#extension-category: envoy.compression.decompressor]
	DecompressorLibrary *DecompressorLibrary `protobuf:"bytes,10,opt,name=decompressor_library,json=decompressorLibrary,proto3" json:"decompressor_library,omitempty"`
	// Configuration for request decompression. Decompression is enabled by default if left empty.
	RequestDirectionConfig *Decompressor_RequestDirectionConfig `protobuf:"bytes,20,opt,name=request_direction_config,json=requestDirectionConfig,proto3" json:"request_direction_config,omitempty"`
	// Configuration for response decompression. Decompression is enabled by default if left empty.
	ResponseDirectionConfig *Decompressor_ResponseDirectionConfig `protobuf:"bytes,30,opt,name=response_direction_config,json=responseDirectionConfig,proto3" json:"response_direction_config,omitempty"`
}

type DecompressorLibrary struct {
	//Algorithm *string `protobuf:"bytes,1,opt,name=algorithm,json=algorithm,proto3" json:"algorithm,omitempty"`

	Gzip   *GzipDecompressor   `protobuf:"bytes,2,opt,name=gzip,json=gzip,proto3,oneof" json:"gzip,omitempty"`
	Brotli *BrotliDecompressor `protobuf:"bytes,3,opt,name=brotli,json=brotli,proto3,oneof" json:"brotli,omitempty"`
	Zstd   *ZstdDecompressor   `protobuf:"bytes,4,opt,name=zstd,json=zstd,proto3,oneof" json:"zstd,omitempty"`
}

// Configuration for filter behavior on the request direction.
type Decompressor_RequestDirectionConfig struct {
	CommonConfig *Decompressor_CommonDirectionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// If set to true, and response decompression is enabled, the filter modifies the Accept-Encoding
	// request header by appending the decompressor_library's encoding. Defaults to true.
	AdvertiseAcceptEncoding *bool `protobuf:"bytes,2,opt,name=advertise_accept_encoding,json=advertiseAcceptEncoding,proto3" json:"advertise_accept_encoding,omitempty"`
}

// Configuration for filter behavior on the response direction.
type Decompressor_ResponseDirectionConfig struct {
	CommonConfig *Decompressor_CommonDirectionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
}

type Decompressor_CommonDirectionConfig struct {
	// Runtime flag that controls whether the filter is enabled for decompression or not. If set to false, the
	// filter will operate as a pass-through filter. If the message is unspecified, the filter will be enabled.
	Enabled *RuntimeFeatureFlag `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// If set to true, will decompress response even if a ``no-transform`` cache control header is set.
	IgnoreNoTransformHeader *bool `protobuf:"varint,2,opt,name=ignore_no_transform_header,json=ignoreNoTransformHeader,proto3" json:"ignore_no_transform_header,omitempty"`
}

// ASMDecompressorStatus defines the observed state of ASMDecompressor
type ASMDecompressorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status defines the state of this instance
	Status string `json:"status,omitempty"`
	// Message defines the possible error message
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ASMDecompressor is the Schema for the ASMDecompressors API
// +genclient
type ASMDecompressor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ASMDecompressorSpec   `json:"spec,omitempty"`
	Status ASMDecompressorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ASMDecompressorList contains a list of ASMDecompressor
type ASMDecompressorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ASMDecompressor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ASMDecompressor{}, &ASMDecompressorList{})
}
