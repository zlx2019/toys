/**
  @author: Zero
  @date: 2023/4/22 10:50:18
  @desc: json工具库 基准测试
		 表示运行json_benchmark_test.go文件的所有基准测试方法, json.go表示需要用到的源文件
		 go test -bench=. json_benchmark_test.go json.go
	     Command:
				-bench=.: 表示运行所有基准测试方法,可以使用*通配符表达式
			    -benchtime=5s: 表示运行基准测试的时间,默认是1s
			    -benchtime=10000x: 该参数也可以作为运行次数,表示执行1万次
				-benchmem: 表示显示内存分配情况
				-count: 表示运行基准测试的次数,默认是1次
		 示例:
		    运行所有基准测试方法,运行时间为5s,显示内存分配情况,运行3次
			go test -bench=. -benchtime=5s -benchmem -count=3 json_benchmark_test.go json.go

**/

package converts

import (
	"encoding/json"
	"testing"
)

type UserStruct struct {
	UserName string `json:"user_name"`
	NickName string `json:"nike_name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Locked   bool   `json:"locked"`
	Student  struct {
		Name string `json:"stu_name"`
		Age  int    `json:"stu_age"`
	} `json:"student"`
}

// 基准测试 使用jsoniter库来序列化结构体
func BenchmarkAnyToJson(b *testing.B) {
	u := &UserStruct{
		UserName: "满城雪",
		NickName: "海问香",
		Age:      19,
		Address:  "北京市海淀区",
		Locked:   true,
		Student: struct {
			Name string `json:"stu_name"`
			Age  int    `json:"stu_age"`
		}{
			Name: "小张同学",
			Age:  16,
		},
	}
	for i := 0; i < b.N; i++ {
		AnyToJson(&u)
	}
}

// 基准测试 使用标准库来序列化结构体
func BenchmarkStdAnyToJson(b *testing.B) {
	u := &JsonStruct{
		UserName: "满城雪",
		NickName: "海问香",
		Age:      19,
		Address:  "北京市海淀区",
		Locked:   true,
		Student: struct {
			Name string `json:"stu_name"`
			Age  int    `json:"stu_age"`
		}{
			Name: "小张同学",
			Age:  16,
		},
		Hobby: []string{"唱", "跳", "rap"},
	}
	for i := 0; i < b.N; i++ {
		json.Marshal(&u)
	}
}

// 基准测试 使用jsoniter库来反序列化结构体
func BenchmarkReadJson(b *testing.B) {
	jsonStr := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	for i := 0; i < b.N; i++ {
		ReadJson[UserStruct](jsonStr)
	}
}

// 基准测试 使用jsoniter库来反序列化结构体(不使用泛型)
func BenchmarkBeadJsonNotReflect(b *testing.B) {
	jsonStr := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	var u UserStruct
	for i := 0; i < b.N; i++ {
		ReadJsonToAny(jsonStr, &u)
	}
}

// 基准测试 使用 unsafe String 和 bytes转换 0.3307ns
func BenchmarkStringAndBytesConvert(b *testing.B) {
	jsonStr := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	for i := 0; i < b.N; i++ {
		bytes := StringToBytes(jsonStr)
		_ = BytesToString(bytes)
	}
}

// 基准测试 使用Json String 和 bytes转换  1000ns
func BenchmarkStringAndBytesConvert2(b *testing.B) {
	jsonStr := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(jsonStr)
		ReadJsonBytes[string](bytes)
	}
}
