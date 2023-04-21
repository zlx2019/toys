/**
  @author: Zero
  @date: 2023/3/27 20:45:56
  @desc: convert 单元测试

**/

package converts

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyToBytes(t *testing.T) {
	bytes, _ := ToBytes(111111.12)
	fmt.Println(bytes)
}

func TestToJson(t *testing.T) {
	json, _ := AnyToJson([]string{"小张", "小王", "小刘"})
	fmt.Println(json) //["小张","小王","小刘"]

	json, _ = AnyToJson(map[string]any{"name": "小张", "age": 18})
	fmt.Println(json) //{"age":18,"name":"小张"}
}

func TestToReadOnlyChannel(t *testing.T) {
	channel := ToReadOnlyChannel([]int{1, 2, 3, 4, 5})
	for item := range channel {
		fmt.Println(item)
	}
}

type Users struct {
	Name string
	Age  int
}

func TestToMap(t *testing.T) {
	us := []Users{Users{Name: "小明", Age: 18}}
	uMap := SliceToMap(us, func(u Users) (string, Users) {
		return u.Name, u
	})
	fmt.Println(uMap) //map[小明:{小明 18}]
}

func TestEncoderAndDecoder(t *testing.T) {
	u := &Users{
		Name: "满城雪",
		Age:  22,
	}
	bytes, _ := EncoderBytes(u)

	u2 := &Users{}
	DecoderBytes(bytes, u2)
	fmt.Println(u)  //&{满城雪 22}
	fmt.Println(u2) //&{满城雪 22}
}

// 测试json序列化
type JsonStruct struct {
	UserName string `json:"user_name"`
	NickName string `json:"nike_name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Locked   bool   `json:"locked"`
	Student  struct {
		Name string `json:"stu_name"`
		Age  int    `json:"stu_age"`
	}
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
	result := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"Student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"
	jsonStr, err := AnyToJson(&u)
	if err != nil {
		panic(err)
	}
	is.Equal(jsonStr, result)

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
