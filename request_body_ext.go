package main

import (
	"bufio"
	"io"
	"sync"
)

// @Title       request_body_ext.go
// @Description 特殊处理的请求体
// @Author      Zero.
// @Create      2024-08-16 17:09


// 有些HTTP客户端在发送请求时，会在请求头中添加 `Expect: 100-continue`, 然后等待服务端响应 `HTTP/1.1 100 Continue\r\n\r\n` 报文
// 接收到该报文后，客户端才会继续发送 Request Body, 这样能够防止网络资源的浪费

// 扩展请求体，每个连接首次读取该请求体时，会向客户端发送 `HTTP/1.1 100 Continue\r\n\r\n` 报文.
type extendRequestBody struct {
	// 保证只响应一次即可
	once sync.Once

	// 连接的请求体
	body io.Reader
	// 客户端连接的写缓冲
	writer *bufio.Writer
}

func wrapExtendBody(req *Request) *extendRequestBody {
	return &extendRequestBody{
		body: req.Body,
		writer: req.conn.bufW,
	}
}

func (rb *extendRequestBody) Read(buf []byte) (int,error) {
	// 首次读取请求体时，响应客户端 Continue 报文
	rb.once.Do(func() {
		_, _ = rb.writer.WriteString(continueMessage)
		_ = rb.writer.Flush()
	})
	return rb.body.Read(buf)
}

// 检查请求头中是否包含 `Expect: 100-continue`
// 如果包含则将请求体升级为 extendRequestBody
func (req *Request) extraChecker() {
	if req.Headers.Get(Expect) != ExpectValue {
		return
	}
	req.Body = wrapExtendBody(req)
}