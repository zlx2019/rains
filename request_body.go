// @Title request_body.go
// @Description HTTP 请求体
// @Author Zero - 2024/8/15 22:34:31

package main

import (
	"io"
	"strconv"
)

var empty = &emptyBody{}

// 表示请求没有请求体，比如 GET 请求
type emptyBody struct {}
// 直接返回 EOF
func (*emptyBody) Read([]byte) (int,error) {
	return 0, io.EOF
}

// RequestBody HTTP 请求体
type RequestBody io.Reader


// RequestBodyParse  解析 HTTP 请求体
// 从 Headers 中读出 `Content-Length`，该属性值为请求正文的数据大小.
func RequestBodyParse(req *Request) {
	if GET == req.Method {
		// GET 请求默认为无主体处理
		goto blank
	}else if val := req.Headers.Get(ContentLength); val != ""{
		// 设置了正文长度则处理
		length, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			goto blank
		}
		// 通过 LimitReader 控制每次读取请求体的数据量
		// 一旦请求体读完就返回 EOF, 而不是一直阻塞读取直到连接Close.
		req.Body = io.LimitReader(req.conn.bufR, length)
	}
	blank:
		req.Body = empty
}


