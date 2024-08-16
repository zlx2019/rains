package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
)

// @Title       handler.go
// @Description HTTP Request Handler
// @Author      Zero.
// @Create      2024-08-14 11:21

// Handler HTTP 处理器
type Handler interface {
	// Handle 处理函数
	Handle(ResponseWriter, *Request)
}

// defaultHandler 通过原生数据流，返回HTTP响应报文.
type defaultHandler struct {

}
func (*defaultHandler) Handle(rw ResponseWriter, request *Request)  {
	slog.Info("Hello")
	// 通过响应流，手动写回 HTTP响应报文.
	// 构建响应数据
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "[Query]name=%s\n", request.Query("name"))
	fmt.Fprintf(buf, "[Query]token=%s\n", request.Query("token"))
	fmt.Fprintf(buf, "[Cookie]foo1=%s\n", request.Cookie("foo1"))
	fmt.Fprintf(buf, "[Cookie]foo2=%s\n", request.Cookie("foo2"))
	fmt.Fprintf(buf, "[Headers]User-Agent=%s\n", request.Headers.Get("User-Agent"))
	fmt.Fprintf(buf, "[Request]Method=%s\n", request.Method)
	fmt.Fprintf(buf, "[Request]RemoteAddr=%s\n", request.RemoteAddr)
	fmt.Fprintf(buf, "[Request]Request=%+v\n", request)

	// 响应行
	_, _ = io.WriteString(rw, "HTTP/1.1 200 OK\r\n")
	// 响应头
	_, _ = io.WriteString(rw, fmt.Sprintf("Content-Length: %d\r\n", buf.Len()))
	_, _ = io.WriteString(rw, "\r\n")
	// 响应体
	_, _ = buf.WriteTo(rw)
}

// 数据回显处理器，将客户端发送的请求体，原封不动返回
// 测试 Chunk 请求体
type echoHandler struct {}
func (eh *echoHandler) Handle(rw ResponseWriter, req *Request) {
	// 读取请求体
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	const prefix = "client body:"
	length := len(body) + len(prefix)
	_, _ = io.WriteString(rw, "HTTP/1.1 200 OK\r\n")
	_, _ = io.WriteString(rw, fmt.Sprintf("Content-Length: %d\r\n", length))
	_, _ = io.WriteString(rw, "\r\n")
	_, _ = io.WriteString(rw, prefix)
	_, _ = rw.Write(body)
}