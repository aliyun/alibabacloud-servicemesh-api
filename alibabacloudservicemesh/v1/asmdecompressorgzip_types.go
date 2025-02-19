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

type GzipDecompressor struct {

	// Value from 9 to 15 that represents the base two logarithmic of the decompressor's window size.
	// The decompression window size needs to be equal or larger than the compression window size.
	// The default window size is 15.
	// This is so that the decompressor can decompress a response compressed by a compressor with any compression window size.
	// For more details about this parameter, please refer to `zlib manual <https://www.zlib.net/manual.html>`_ > inflateInit2.
	WindowBits *uint32 `protobuf:"bytes,1,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Value for zlib's decompressor output buffer. If not set, defaults to 4096.
	// See https://www.zlib.net/manual.html for more details.
	ChunkSize *uint32 `protobuf:"bytes,2,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	// An upper bound to the number of times the output buffer is allowed to be bigger than the size of
	// the accumulated input. This value is used to prevent decompression bombs. If not set, defaults to 100.
	// [#comment:TODO(rojkov): Re-design the Decompressor interface to handle compression bombs gracefully instead of this quick solution.
	// See https://github.com/envoyproxy/envoy/commit/d4c39e635603e2f23e1e08ddecf5a5fb5a706338 for details.]
	MaxInflateRatio *uint32 `protobuf:"bytes,3,opt,name=max_inflate_ratio,json=maxInflateRatio,proto3" json:"max_inflate_ratio,omitempty"`
}
