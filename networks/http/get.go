/**
  @author: Zero
  @date: 2023/5/3 19:34:42
  @desc: HTTP GET请求函数库

**/

package http

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

// GET 发送GET请求,将Response Body的JSON数据映射到T对象中,并且返回
func GET[T any](url string) (T, error) {
	return GETWithQuery[T](url, nil)
}

// GETWithQuery 发送GET请求,将Response Body的JSON数据映射到T对象中,并且返回
func GETWithQuery[T any](url string, query map[string]string) (T, error) {
	var payload T
	_, err := GetWithOptions(url, query, nil, &payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

// Get 发送 GET 请求,返回响应对象,由上层处理。
func Get(url string) (*resty.Response, error) {
	return GetWithOptions(url, nil, nil, nil)
}

// GetWithQuery 发送 GET 请求,携带Query参数,返回响应对象,由上层处理。
func GetWithQuery(url string, query map[string]string) (*resty.Response, error) {
	return GetWithOptions(url, query, nil, nil)
}

// GetWithOptions 发送GET请求,携带Query参数、Headers，返回响应对象,由上层处理。
func GetWithOptions(url string, query map[string]string, headers map[string]string, resultPayload any) (*resty.Response, error) {
	return Request(url, http.MethodGet, resultPayload, nil, headers, query)
}
