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

type ZstdCompressor struct {

	// Set compression parameters according to pre-defined compression level table.
	// Note that exact compression parameters are dynamically determined,
	// depending on both compression level and source content size (when known).
	// Value 0 means default, and default level is 3.
	// Setting a level does not automatically set all other compression parameters
	// to default. Setting this will however eventually dynamically impact the compression
	// parameters which have not been manually set. The manually set
	// ones will 'stick'.
	CompressionLevel *uint32 `protobuf:"bytes,1,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"`
	// A 32-bits checksum of content is written at end of frame. If not set, defaults to false.
	EnableChecksum *bool `protobuf:"varint,2,opt,name=enable_checksum,json=enableChecksum,proto3" json:"enable_checksum,omitempty"`
	// The higher the value of selected strategy, the more complex it is,
	// resulting in stronger and slower compression.
	// Special: value 0 means "use default strategy".
	Strategy *string `protobuf:"varint,3,opt,name=strategy,proto3,enum=envoy.extensions.compression.zstd.compressor.v3.Zstd_Strategy" json:"strategy,omitempty"`
	// A dictionary for compression. Zstd offers dictionary compression, which greatly improves
	// efficiency on small files and messages. Each dictionary will be generated with a dictionary ID
	// that can be used to search the same dictionary during decompression.
	// Please refer to `zstd manual <https://github.com/facebook/zstd/blob/dev/programs/zstd.1.md#dictionary-builder>`_
	// to train a specific dictionary for compression.
	Dictionary *DataSource `protobuf:"bytes,4,opt,name=dictionary,proto3" json:"dictionary,omitempty"`
	// Value for compressor's next output buffer. If not set, defaults to 4096.
	ChunkSize *uint32 `protobuf:"bytes,5,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

type DataSource struct {
	// Local filesystem data source.
	Filename *string `protobuf:"bytes,1,opt,name=filename,proto3,oneof" json:"filename,omitempty"`
	// Bytes inlined in the configuration.
	InlineBytes *string `protobuf:"bytes,2,opt,name=inline_bytes,json=inlineBytes,proto3,oneof" json:"inlineBytes,omitempty"`
	// String inlined in the configuration.
	InlineString *string `protobuf:"bytes,3,opt,name=inline_string,json=inlineString,proto3,oneof" json:"inlineString,omitempty"`
	// Environment variable data source.
	EnvironmentVariable *string `protobuf:"bytes,4,opt,name=environment_variable,json=environmentVariable,proto3,oneof" json:"environmentVariable,omitempty"`
}

// Reference to http://facebook.github.io/zstd/zstd_manual.html
type Zstd_Strategy int32

const (
	Zstd_DEFAULT  Zstd_Strategy = 0
	Zstd_FAST     Zstd_Strategy = 1
	Zstd_DFAST    Zstd_Strategy = 2
	Zstd_GREEDY   Zstd_Strategy = 3
	Zstd_LAZY     Zstd_Strategy = 4
	Zstd_LAZY2    Zstd_Strategy = 5
	Zstd_BTLAZY2  Zstd_Strategy = 6
	Zstd_BTOPT    Zstd_Strategy = 7
	Zstd_BTULTRA  Zstd_Strategy = 8
	Zstd_BTULTRA2 Zstd_Strategy = 9
)

// Enum value maps for Zstd_Strategy.
var (
	Zstd_Strategy_name = map[int32]string{
		0: "DEFAULT",
		1: "FAST",
		2: "DFAST",
		3: "GREEDY",
		4: "LAZY",
		5: "LAZY2",
		6: "BTLAZY2",
		7: "BTOPT",
		8: "BTULTRA",
		9: "BTULTRA2",
	}
	Zstd_Strategy_value = map[string]int32{
		"DEFAULT":  0,
		"FAST":     1,
		"DFAST":    2,
		"GREEDY":   3,
		"LAZY":     4,
		"LAZY2":    5,
		"BTLAZY2":  6,
		"BTOPT":    7,
		"BTULTRA":  8,
		"BTULTRA2": 9,
	}
)
