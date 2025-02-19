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
	"math"
)

// refer to https://github.com/envoyproxy/go-control-plane/blob/84c5a14eb8c8b5b99efeecd3801fea84d628852d/envoy/extensions/filters/http/compressor/v3/compressor.pb.go

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BrotliCompressor struct {
	// Value from 0 to 11 that controls the main compression speed-density lever.
	// The higher quality, the slower compression. The default value is 3.
	Quality *uint32 `protobuf:"bytes,1,opt,name=quality,proto3" json:"quality,omitempty"`
	// A value used to tune encoder for specific input. For more information about modes,
	// please refer to brotli manual: https://brotli.org/encode.html#aa6f
	// This field will be set to "DEFAULT" if not specified.
	EncoderMode *string `protobuf:"varint,2,opt,name=encoder_mode,json=encoderMode,proto3,enum=envoy.extensions.compression.brotli.compressor.v3.Brotli_EncoderMode" json:"encoder_mode,omitempty"`
	// Value from 10 to 24 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 18.
	// For more details about this parameter, please refer to brotli manual:
	// https://brotli.org/encode.html#a9a8
	WindowBits *uint32 `protobuf:"bytes,3,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Value from 16 to 24 that represents the base two logarithmic of the compressor's input block
	// size. Larger input block results in better compression at the expense of memory usage. The
	// default is 24. For more details about this parameter, please refer to brotli manual:
	// https://brotli.org/encode.html#a9a8
	InputBlockBits *uint32 `protobuf:"bytes,4,opt,name=input_block_bits,json=inputBlockBits,proto3" json:"input_block_bits,omitempty"`
	// Value for compressor's next output buffer. If not set, defaults to 4096.
	ChunkSize *uint32 `protobuf:"bytes,5,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	// If true, disables "literal context modeling" format feature.
	// This flag is a "decoding-speed vs compression ratio" trade-off.
	DisableLiteralContextModeling *bool `protobuf:"varint,6,opt,name=disable_literal_context_modeling,json=disableLiteralContextModeling,proto3" json:"disable_literal_context_modeling,omitempty"`
}

type Brotli_EncoderMode int32

const (
	Brotli_DEFAULT Brotli_EncoderMode = 0
	Brotli_GENERIC Brotli_EncoderMode = 1
	Brotli_TEXT    Brotli_EncoderMode = 2
	Brotli_FONT    Brotli_EncoderMode = 3
)

// Enum value maps for Brotli_EncoderMode.
var (
	Brotli_EncoderMode_name = map[int32]string{
		0: "DEFAULT",
		1: "GENERIC",
		2: "TEXT",
		3: "FONT",
	}
	Brotli_EncoderMode_value = map[string]int32{
		"DEFAULT": 0,
		"GENERIC": 1,
		"TEXT":    2,
		"FONT":    3,
	}
)
