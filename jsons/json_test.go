/**
  @author: Zero
  @date: 2023/5/9 20:48:41
  @desc:

**/

package jsons

import (
	"github.com/bytedance/sonic/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	// Json
	json := "{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}"

	// 获取string节点值
	username, err := GetString(json, "user_name")
	is.True(is.NoError(err))
	is.Equal(username, "满城雪")

	// 获取int节点值
	age, err := GetNode[int64](json, "age", func(node ast.Node) (int64, error) {
		return node.Int64()
	})
	is.True(is.NoError(err))
	is.Equal(age, int64(19))

}

func TestValida(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	is.True(Valida([]byte("[1,2,3,4]")))
	is.True(Valida([]byte("{\"user_name\":\"满城雪\",\"nike_name\":\"海问香\",\"age\":19,\"address\":\"北京市海淀区\",\"locked\":true,\"student\":{\"stu_name\":\"小张同学\",\"stu_age\":16}}")))
	is.False(Valida([]byte("\"name\": ")))
}
