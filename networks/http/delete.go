/**
  @author: Zero
  @date: 2023/5/6 17:04:08
  @desc: HTTP Delete请求函数库

**/

package http

import "net/http"

// DELETE HTTP-DELETE请求
// url 请求地址
// body 请求体参数
func DELETE[T any](url string, body any) (T, error) {
	var payload T
	_, err := Request(url, http.MethodDelete, &payload, body, nil, nil)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
