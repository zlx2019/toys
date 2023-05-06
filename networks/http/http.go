/**
  @author: Zero
  @date: 2023/5/3 08:23:00
  @desc: http请求相关函数库,基于go-resty库封装

**/

package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/zlx2019/toys/converts"
	"github.com/zlx2019/toys/texts"
)

// Http客户端
var httpClient *resty.Client

func init() {
	// 初始化客户端
	httpClient = resty.New()
	// 使用自定义的Json序列化函数
	httpClient.JSONMarshal = converts.AnyToJsonBytes
	httpClient.JSONUnmarshal = converts.ReadJsonBytesToAny
}

// Request 发起HTTP请求
// url: 请求地址
// method: 请求类型
// requestBody: 请求体
// headers: 请求头
// query: Query 参数
// resultPayload: 响应结果映射载体
func Request(url, method string, resultPayload, requestBody any, headers, query map[string]string) (*resty.Response, error) {
	// 构建request
	request := httpClient.R()
	// 设置请求地址
	request.URL = url
	// 设置请求类型
	if texts.NotEmpty(method) {
		request.Method = method
	}
	// 设置请求头
	if headers != nil && len(headers) > 0 {
		request.SetHeaders(headers)
	}
	// 设置Query参数
	if query != nil && len(query) > 0 {
		request.SetQueryParams(query)
	}
	// 设置请求体参数
	if requestBody != nil {
		request.SetBody(requestBody)
	}
	// 设置响应数据载体,将ResponseBody以Json格式写入该对象
	if resultPayload != nil {
		request.SetResult(resultPayload)
	}
	// 发起请求
	return request.Send()
}
