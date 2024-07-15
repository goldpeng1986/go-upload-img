package tools

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 用于将来自其他请求的返回数据进行结构化
type T_Return_Status struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Return_Status(r *ghttp.Request, mData g.Map, code string, msg string) {
	if code == "" {
		code = "200"
	}
	if msg == "" {
		msg = "success"
	}
	r.Response.WriteJson(g.Map{
		"code": code,
		"data": mData,
		"msg":  msg,
	})
}
