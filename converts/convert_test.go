/**
  @author: Zero
  @date: 2023/3/27 20:45:56
  @desc:

**/

package converts

import (
	"fmt"
	"testing"
)

func TestAnyToBytes(t *testing.T) {
	bytes, _ := ToBytes(111111.12)
	fmt.Println(bytes)
}

func TestToJson(t *testing.T) {
	json, _ := ToJson([]string{"小张", "小王", "小刘"})
	fmt.Println(json) //["小张","小王","小刘"]

	json, _ = ToJson(map[string]any{"name": "小张", "age": 18})
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
