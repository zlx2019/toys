/**
  @author: Zero
  @date: 2023/5/4 12:27:42
  @desc:

**/

package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试POST 请求
func TestPOST(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	// Request Body
	body := map[string]any{"teacher_id": 10001, "name": "Go入门到入土"}
	// 发起请求,指定结果映射载体为string类型
	result, err := POST[string]("http://127.0.0.1:8080/courses/", body)
	is.NoError(err)
	is.Equal(result, "Success")
}
