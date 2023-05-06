/**
  @author: Zero
  @date: 2023/5/4 12:09:03
  @desc: HTTP Post请求函数库

**/

package http

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

// POST 发起Post类型请求 将Response Body的JSON数据映射到T对象中,并且返回
func POST[T any](url string, body any) (T, error) {
	var payload T
	_, err := PostWithOptions(url, body, &payload, nil, nil)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

// PostWithOptions 发起Post类型请求,携带请求体、结果载体、Query参数、请求头
// 返回响应对象,由上层处理.
func PostWithOptions(url string, body, resultPayload any, query, headers map[string]string) (*resty.Response, error) {
	return Request(url, http.MethodPost, resultPayload, body, headers, query)
}
