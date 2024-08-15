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

// 连接处理逻辑
func (c *Connection) serve() {
	defer c.Close()
	// panic handler
	defer func() {
		if err := recover(); err != nil {
			slog.Error("conn panic err: %v", err)
		}
		_ = c.Close()
	}()
	// keep-alive 长连接处理

	for i := 0; i < 1; i++ {
		// 读取连接数据流，解析为 HTTP Request
		req, err := RequestParse(c)
		if err != nil {
			slog.Error("conn unpack request err: %v", err)
			return
		}
		// 分配连接响应流
		response := wrapResponse(c)
		// 处理 HTTP 请求
		c.serv.Handler.Handle(response, req)
		if err = c.bufW.Flush(); err != nil {
			return
		}
	}
}
