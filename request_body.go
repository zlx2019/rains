// @Title request_body.go
// @Description HTTP 请求体
// @Author Zero - 2024/8/15 22:34:31

package main

import (
	"io"
	"strconv"
)


// 表示空的请求体，比如 GET 请求
type emptyRequestBody struct {}
// 直接返回 EOF
func (*emptyRequestBody) Read([]byte) (int,error) {
	return 0, io.EOF
}
// 空请求体
var emptyBody = &emptyRequestBody{}


// RequestBody 表示一个最基本的 HTTP 请求体
type RequestBody io.Reader

// HTTP 请求是否使用 Chunk 编码
func (req *Request) usingChunked() bool {
	return req.Headers.Get(TransferEncoding) == "chunked"
}

// HTTP 请求头中是否设置了 Content-Length, 取出并返回.
func (req *Request) usingLength() bool {
	if val := req.Headers.Get(ContentLength); val != "" {
		return true
	}
	return false
}

// RequestBodyParse  解析 HTTP 请求体
// 从 Headers 中读出 `Content-Length`，该属性值为请求正文的数据大小.
func (req *Request) RequestBodyParse(){
	switch  {
	case req.usingChunked():
		// 通过 Chunk 编码 解析请求体
		req.Body = wrapChunkReader(req.conn.bufR)
		req.extraChecker()
		return
	case req.usingLength():
		// 通过 Content-Length 解析请求体
		val := req.Headers.Get(ContentLength)
		length, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			// 使用 LimitReader 控制每次读取请求体的数据量
			// 一旦请求体读完就返回 EOF, 而不是一直阻塞读取直到连接Close.
			req.Body = io.LimitReader(req.conn.bufR, length)
			req.extraChecker()
			return
		}
	}
	req.Body = emptyBody
	return
}


