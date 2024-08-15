package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// @Title       request.go
// @Description HTTP Request
// @Author      Zero.
// @Create      2024-08-14 11:14

const lineEnd = "\r\n"
const headEnd = "\r\n\r\n"
const headSep = ':'

// Request 表示 HTTP 请求信息
type Request struct {
	//======= 请求行 ========
	Method     string // 请求方法
	RequestURI string // 请求的资源Uri
	Protocol   string // 协议以及版本

	Headers Header    // 请求头
	Body    io.Reader // 请求体

	Url         *url.URL          // 请求URL
	RemoteAddr  string            // 客户端地址信息
	conn        *Connection       // 客户端连接
	cookies     map[string]string //  Cookie 信息
	queryParams Query             //  Query 查询参数
}

// RequestParse
// 从流中读取字节流，解析为 Request 请求体.
// HTTP 协议请求报文，分为三个部分:
//   - 请求行\r\n
//   - 请求头\r\n
//   - \r\n
//   - 请求体
func RequestParse(c *Connection) (req *Request, err error) {
	req = new(Request)
	req.conn = c
	req.RemoteAddr = c.RemoteAddr().String()

	// 解析请求行
	req.Method, req.RequestURI, req.Protocol, err = requestLineParse(c.bufR)
	if err != nil {
		return
	}

	// 解析Query请求参数
	if err = req.requestQueryParse(); err != nil {
		return
	}

	// 解析请求头
	req.Headers, err = requestHeadersParse(c.bufR)
	if err != nil {
		return
	}

	req.Body = newResponse(c)

	return
}

// 解析 HTTP 请求行
// +--------+---------+----------+
// | Method |   URL   | Version  |
// +--------+---------+----------+
// | GET    | /index  | HTTP/1.1 |
// +--------+---------+----------+
func requestLineParse(reader *bufio.Reader) (m string, u string, v string, e error) {
	line, e := readLine(reader)
	if e != nil {
		return
	}
	// 按空格分割，得到 Method、URL、Version
	_, e = fmt.Sscanf(string(line), "%s%s%s", &m, &u, &v)
	return
}

// 解析 HTTP 请求头
// k1:v1\r\n
// k2:v2\r\n
// k3:v3\r\n
// .....\r\n\r\n
func requestHeadersParse(reader *bufio.Reader) (Header, error) {
	header := make(Header)
	for {
		// 以 \r\n 结尾，读取每一行数据，直到读取到 \r\n\r\n
		line, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		if len(line) == 0 {
			break
		}
		// 每一组请求头以 ':' 分割
		index := bytes.IndexByte(line, headSep)
		if index == -1 {
			return nil, errors.New("invalid request header")
		}
		if index == len(line)-1 {
			continue
		}
		key := string(line[:index])
		value := strings.TrimSpace(string(line[index+1:]))
		header[key] = append(header[key], value)
	}
	return header, nil
}

// 从流中读取一行内容
func readLine(reader *bufio.Reader) ([]byte, error) {
	line, prefix, err := reader.ReadLine()
	if err != nil {
		return line, err
	}
	var content []byte
	// 行内容太长，还未读取完毕
	for prefix {
		content, prefix, err = reader.ReadLine()
		if err != nil {
			break
		}
		line = append(line, content...)
	}
	return line, err
}
