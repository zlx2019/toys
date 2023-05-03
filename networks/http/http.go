/**
  @author: Zero
  @date: 2023/5/3 08:23:00
  @desc: http请求相关函数库,基于go-resty库封装

**/

package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/zlx2019/toys/converts"
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
