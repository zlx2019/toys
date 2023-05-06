/**
  @author: Zero
  @date: 2023/5/6 17:11:27
  @desc: HTTP Patch请求函数库

**/

package http

import "net/http"

// PATCH HTTP-Patch请求
// url 请求地址
// body 请求体参数
func PATCH[T any](url string, body any) (T, error) {
	var payload T
	_, err := Request(url, http.MethodPatch, &payload, body, nil, nil)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
