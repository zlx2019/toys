/**
  @author: Zero
  @date: 2023/4/24 19:32:50
  @desc: 性能比较基准测试
		 结构体序列化为[]byte性能比较
		 go test -bench=. compare_benchmark_test.go json.go convert.go
**/

package converts

import (
	"encoding/json"
	"testing"
)

type compare struct {
	UserName string
	NickName string
	Age      int
	Address  string
	Locked   bool
}

// 直接使用json库序列化为[]byte(推荐)
func BenchmarkWithJsonByte(b *testing.B) {
	c := &compare{
		UserName: "满城雪",
		NickName: "海问香",
		Age:      19,
		Address:  "北京市海淀区",
		Locked:   true,
	}
	for i := 0; i < b.N; i++ {
		AnyToJsonBytes(c)
	}
}

func BenchmarkWithStdJson(b *testing.B) {
	c := &compare{
		UserName: "满城雪",
		NickName: "海问香",
		Age:      19,
		Address:  "北京市海淀区",
		Locked:   true,
	}
	for i := 0; i < b.N; i++ {
		json.Marshal(c)
	}
}

// 使用标准库gob序列化为[]byte(性能极差)
func BenchmarkWithGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := &compare{
			UserName: "满城雪",
			NickName: "海问香",
			Age:      19,
			Address:  "北京市海淀区",
			Locked:   true,
		}
		EncoderBytes(c)
	}
}

// 使用json库序列化为string,再通过unsafe序列化[]byte
func BenchmarkWithJsonString(b *testing.B) {
	c := &compare{
		UserName: "满城雪",
		NickName: "海问香",
		Age:      19,
		Address:  "北京市海淀区",
		Locked:   true,
	}
	for i := 0; i < b.N; i++ {
		toJson, _ := AnyToJson(c)
		_ = StringToBytes(toJson)
	}
}
