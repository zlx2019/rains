package main

import (
	"bufio"
	"io"
	"log/slog"
	"net"
)

// @Title       conn.go
// @Description HTTP Connection
// @Author      Zero.
// @Create      2024-08-14 11:15

// Connection TCP连接扩展
type Connection struct {
	net.Conn
	// 读缓冲
	bufR *bufio.Reader
	limitR *io.LimitedReader
	// 写缓冲
	bufW *bufio.Writer

	serv *Server
}

// 建立新的连接
func wrapConn(conn net.Conn, server *Server) *Connection {
	lr := &io.LimitedReader{R: conn, N: 1 << 20}
	return &Connection{
		conn,
		bufio.NewReaderSize(conn, 4 << 10),
		lr,
		bufio.NewWriterSize(conn, 4 << 10), // 写缓冲大小为4kb
		server,
	}
}

// 处理每一条 Connection 连接
func (c *Connection) serve() {
	defer c.Close()
	defer func() {
		if err := recover(); err != nil {
			slog.Error("conn panic err: %v", err)
		}
		_ = c.Close()
	}()
	// HTTP/1.1 支持 Keep-alive 长连接，暂时使用 for 不停的解析 HTTP 请求
	for i := 0; i < 1; i++ {
		// 解析 Request
		request, err := requestParse(c)
		if err != nil {
			slog.Error("parsing http request failed: %v", err.Error())
			return
		}
		// 设置 Response
		response := mountResponse(c)
		// 通过外部传入的 处理器来处理HTTP请求
		c.serv.Handler.Handle(response, request)
		// 请求完成，清理
		if err = request.finish(); err != nil {
			return
		}
	}
}
