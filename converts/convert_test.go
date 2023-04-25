/**
  @author: Zero
  @date: 2023/3/27 20:45:56
  @desc: convert 单元测试

**/

package converts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
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

// int 与 []byte 互转
func TestIntAndBytesConvert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// int
	var num = 150
	bytes, err := IntegerToBytes(num)
	fmt.Println(unsafe.Sizeof(num))
	is.NoError(err)
	result, err := BytesToInteger[int64](bytes)
	is.NoError(err)
	is.Equal(int64(num), result)

	//int8
	var num2 int8 = 100
	bytes, err = IntegerToBytes(num2)
	is.NoError(err)
	result2, err := BytesToInteger[int8](bytes)
	is.NoError(err)
	is.Equal(num2, result2)

	//int16
	var num3 int16 = 10012
	bytes, err = IntegerToBytes(num3)
	is.NoError(err)
	result3, err := BytesToInteger[int16](bytes)
	is.NoError(err)
	is.Equal(num3, result3)

	//int32
	var num4 int32 = 10012
	bytes, err = IntegerToBytes(num4)
	is.NoError(err)
	result4, err := BytesToInteger[int32](bytes)
	is.NoError(err)
	is.Equal(num4, result4)

	//int64
	var num5 int64 = 10012
	bytes, err = IntegerToBytes(num5)
	is.NoError(err)
	result5, err := BytesToInteger[int64](bytes)
	is.NoError(err)
	is.Equal(num5, result5)
}
