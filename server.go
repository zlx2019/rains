package main

import (
	"log/slog"
	"net"
)

// @Title       server.go
// @Description HTTP Server
// @Author      Zero.
// @Create      2024-08-14 11:20

// Server HTTP 服务端
type Server struct {
	Addr string
	Handler Handler
}

func NewHTTPServer(addr string, handler Handler) *Server {
	return &Server{
		addr,
		handler,
	}
}

// Startup 启动HTTP服务
func (s *Server) Startup() error {
	listen, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	slog.Info("ListenAndServe on " + s.Addr)
	// 接收每个TCP连接，包装为Connection后启用协程处理
	for  {
		conn, err := listen.Accept()
		if err != nil {
			slog.Error("accept connection failed: %v", err)
			continue
		}
		c := wrapConn(conn, s)
		go c.serve()
	}
}