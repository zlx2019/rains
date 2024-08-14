package main

// @Title       handler.go
// @Description HTTP Request Handler
// @Author      Zero.
// @Create      2024-08-14 11:21

// Handler HTTP 处理器
type Handler interface {
	// Handle 处理函数
	Handle(ResponseWriter, *Request)
}