/**
  @author: Zero
  @date: 2023/4/26 16:27:38
  @desc: 基准测试
		 go test -bench=. packer_benchmark_test.go normal_packer.go packer.go
**/

package packer

import (
	"bytes"
	"github.com/zlx2019/toys/networks/protocol"
	"testing"
)

// 52435900                22.58 ns/op
func BenchmarkNormalPacker_UnPack(b *testing.B) {
	m := &protocol.Message{
		ID:      100001,
		Payload: []byte("Hello,~你好"),
	}
	packer := NewNormalPacker()
	pack, _ := packer.Pack(m)
	for i := 0; i < b.N; i++ {
		_, _ = packer.UnPack(pack)
	}
}

// 72374902                16.43 ns/op
func BenchmarkNormalPacker_UnPackReader(b *testing.B) {
	m := &protocol.Message{
		ID:      100001,
		Payload: []byte("Hello,~你好"),
	}
	packer := NewNormalPacker()
	pack, _ := packer.Pack(m)
	reader := bytes.NewReader(pack)
	for i := 0; i < b.N; i++ {
		_, _ = packer.UnPackReader(reader)
	}
}
