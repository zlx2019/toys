/**
  @author: Zero
  @date: 2023/5/6 17:00:41
  @desc: HTTP Put请求函数库

**/

package http

import "net/http"

// Put HTTP-PUT请求
// url 请求地址
// body 请求体参数
func Put[T any](url string, body any) (T, error) {
	var payload T
	_, err := Request(url, http.MethodPut, &payload, body, nil, nil)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
