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

type GzipCompressor struct {
	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. The default value is 5.
	MemoryLevel *uint32 `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	// A value used for selecting the zlib compression level. This setting will affect speed and
	// amount of compression applied to the content. "BEST" provides higher compression at the cost of
	// higher latency, "SPEED" provides lower compression with minimum impact on response time.
	// "DEFAULT" provides an optimal result between speed and compression. This field will be set to
	// "DEFAULT" if not specified.
	CompressionLevel *string `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	// A value used for selecting the zlib compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT" will be the best choice, though
	// there are situations which changing this parameter might produce better results. For example,
	// run-length encoding (RLE) is typically used when the content is known for having sequences
	// which same data occurs many consecutive times. For more information about each strategy, please
	// refer to zlib manual.
	CompressionStrategy *string `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits *uint32 `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Set of configuration parameters common for all compression filters. You can define
	// ``content_length``, ``content_type`` and other parameters in this field.
	//Compressor *v3.Compressor `protobuf:"bytes,10,opt,name=compressor,proto3" json:"compressor,omitempty"`
	// Value for Zlib's next output buffer. If not set, defaults to 4096.
	// See https://www.zlib.net/manual.html for more details. Also see
	// https://github.com/envoyproxy/envoy/issues/8448 for context on this filter's performance.
	ChunkSize *uint32 `protobuf:"bytes,11,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

// Enum value maps for Gzip_CompressionLevel_Enum.
var (
	Gzip_CompressionLevel_Enum_name = map[int32]string{
		0: "DEFAULT",
		1: "BEST",
		2: "SPEED",
	}
	Gzip_CompressionLevel_Enum_value = map[string]int32{
		"DEFAULT": 0,
		"BEST":    1,
		"SPEED":   2,
	}
)

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

// Enum value maps for Gzip_CompressionStrategy.
var (
	Gzip_CompressionStrategy_name = map[int32]string{
		0: "DEFAULT",
		1: "FILTERED",
		2: "HUFFMAN",
		3: "RLE",
	}
	Gzip_CompressionStrategy_value = map[string]int32{
		"DEFAULT":  0,
		"FILTERED": 1,
		"HUFFMAN":  2,
		"RLE":      3,
	}
)
