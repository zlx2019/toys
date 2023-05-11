/**
  @author: Zero
  @date: 2023/4/22 10:47:31
  @desc: json相关函数库单元测试

**/

package converts

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试json序列化
type JsonStruct struct {
	UserName string  `json:"user_name"`
	NickName string  `json:"nike_name"`
	Age      float64 `json:"age"`
	Address  string  `json:"address"`
	Locked   bool    `json:"locked"`
	Student  struct {
		Name string `json:"stu_name"`
		Age  int    `json:"stu_age"`
	} `json:"student"`
	Hobby []string `json:"hobby"`
}

// 测试 结构体序列化为Json
func TestAnyToJson(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
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
	}
	result := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	jsonStr, err := AnyToJson(&u)
	if err != nil {
		panic(err)
	}
	is.Equal(jsonStr, result)
	fmt.Println(jsonStr)
}

// 测试 Json序列化为结构体
func TestReadJson(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	var json = "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"Student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	jsonStruct, err := ReadJson[JsonStruct](json)
	is.NoError(err)
	is.Equal(jsonStruct.UserName, "满城雪")
	fmt.Println(jsonStruct)
}

// 测试通过json路径,获取json节点的值
func TestGetJsonNode(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	var json = "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	stuName, err := GetJsonNode[string](json, "student.stu_name", func(any jsoniter.Any) string {
		return any.ToString()
	})
	is.NoError(err)
	is.Equal(stuName, "小张同学")

}

// 测试通过泛型,获取json节点的值
func TestGetJsonNodeToString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var json = "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	val, err := GetJsonNodeToString(json, "student.stu_age")
	is.NoError(err)
	is.Equal(val, "16")
}

// 测试通过泛型,获取json节点的值
func TestGetJsonNodeFromPaths(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var json = "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	val, err := GetJsonNodeFromPaths(json, "student", "stu_age")
	is.NoError(err)
	is.Equal(val, "16")
}

// 测试使用 unsafe String 和 bytes转换
func TestStringAndBytesConvert(t *testing.T) {
	jsonStr := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	bytes := StringToBytes(jsonStr)
	bytesToString := BytesToString(bytes)
	fmt.Println(jsonStr)
	fmt.Println(bytesToString)
}

// 测试 使用Json String 和 bytes转换
func TestStringToBytes2(t *testing.T) {
	jsonStr := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	bytes, _ := AnyToJsonBytes(jsonStr)
	bytesToString, _ := ReadJsonBytes[string](bytes)
	fmt.Println(jsonStr)
	fmt.Println(bytesToString)
}
