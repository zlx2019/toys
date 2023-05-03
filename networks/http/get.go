/**
  @author: Zero
  @date: 2023/5/3 19:34:42
  @desc: HTTP GET请求函数库

**/

package http

import "github.com/go-resty/resty/v2"

// Get 发送 GET 请求,返回响应对象,由上层处理。
func Get(url string) (*resty.Response, error) {
	return GetWithQueryAndHeaders(url, nil, nil, nil)
}

// GetWithQuery 发送 GET 请求,携带Query参数,返回响应对象,由上层处理。
func GetWithQuery(url string, query map[string]string) (*resty.Response, error) {
	return GetWithQueryAndHeaders(url, query, nil, nil)
}

// GetWithQueryAndHeaders 发送GET请求,携带Query参数、Headers，返回响应对象,由上层处理。
func GetWithQueryAndHeaders(url string, query map[string]string, headers map[string]string, payload any) (*resty.Response, error) {
	request := httpClient.R()
	// 设置结果映射对象
	if payload != nil {
		request.SetResult(payload)
	}
	// 设置Query参数
	if query != nil {
		request.SetQueryParams(query)
	}
	//设置请求头
	if headers != nil {
		request.SetHeaders(headers)
	}
	// 发起请求
	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GET 发送GET请求,将Response Body的JSON数据映射到T对象中,并且返回
func GET[T any](url string) (T, error) {
	return GETWithQuery[T](url, nil)
}

// GETWithQuery 发送GET请求,将Response Body的JSON数据映射到T对象中,并且返回
func GETWithQuery[T any](url string, query map[string]string) (T, error) {
	var payload T
	_, err := GetWithQueryAndHeaders(url, query, nil, &payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
