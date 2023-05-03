/**
  @author: Zero
  @date: 2023/5/3 19:35:48
  @desc:

**/

package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zlx2019/toys/converts"
	"regexp"
	"testing"
)

// 测试 GET 请求 Text响应数据
func TestGetResultText(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	response, err := Get("http://127.0.0.1:8080/health")
	is.NoError(err)
	// 获取Get请求 Text格式响应
	result := response.String()
	ok, err := regexp.MatchString("I'm OK. \\d times", result)
	is.NoError(err)
	is.True(ok)
}

// 测试GET 请求 Json响应数据
func TestGetResultWithJson(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	response, err := Get("http://127.0.0.1:8080/courses/")
	is.NoError(err)
	// 结果集
	var list []Course
	// 将Json响应体 手动写入到list中
	err = converts.ReadJsonBytesToAny(response.Body(), &list)
	is.NoError(err)
	fmt.Println(list)
}

// 测试GET请求 直接通过泛型映射结果
func TestGET(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	courses, err := GET[[]Course]("http://127.0.0.1:8080/courses/")
	is.NoError(err)
	fmt.Println(courses)
}
