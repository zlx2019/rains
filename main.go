package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
)

// @Title       main.go
// @Description Main
// @Author      Zero.
// @Create      2024-08-14 11:08


type myHandler struct {

}

func (*myHandler) Handle(rw ResponseWriter, request *Request)  {
	slog.Info("Hello")

	// 构建响应数据
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "[Query]name=%s\n", request.Query("name"))
	fmt.Fprintf(buf, "[Query]token=%s\n", request.Query("token"))
	fmt.Fprintf(buf, "[Cookie]foo1=%s\n", request.Cookie("foo1"))
	fmt.Fprintf(buf, "[Cookie]foo2=%s\n", request.Cookie("foo2"))
	fmt.Fprintf(buf, "[Header]User-Agent=%s\n", request.Headers.Get("User-Agent"))
	fmt.Fprintf(buf, "[Request]Method=%s\n", request.Method)
	fmt.Fprintf(buf, "[Request]RemoteAddr=%s\n", request.RemoteAddr)
	fmt.Fprintf(buf, "[Request]Request=%+v\n", request)


	// 手动响应 HTTP 成功报文

	// 响应行
	_, _ = io.WriteString(rw, "HTTP/1.1 200 OK\r\n")
	// 响应头
	_, _ = io.WriteString(rw, fmt.Sprintf("Content-Length: %d\r\n", buf.Len()))
	_, _ = io.WriteString(rw, "\r\n")
	// 响应体
	_, _ = io.Copy(rw, buf)

}


// Main
func main() {
	server := NewHTTPServer(":9090", new(myHandler))
	panic(server.Startup())
}
