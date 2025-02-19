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

type BrotliDecompressor struct {

	// If true, disables "canny" ring buffer allocation strategy.
	// Ring buffer is allocated according to window size, despite the real size of the content.
	DisableRingBufferReallocation *bool `protobuf:"varint,1,opt,name=disable_ring_buffer_reallocation,json=disableRingBufferReallocation,proto3" json:"disable_ring_buffer_reallocation,omitempty"`
	// Value for decompressor's next output buffer. If not set, defaults to 4096.
	ChunkSize *uint32 `protobuf:"bytes,2,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}
