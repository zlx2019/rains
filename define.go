// @Title define.go
// @Description 常量
// @Author Zero - 2024/8/15 22:19:19

package main

const CRLF = "\r\n"
const CRLF2 = "\r\n\r\n"
const headSep = ':'

const (
	// ReadWriteBufferSize 连接的读写缓冲大小
	ReadWriteBufferSize = 4 << 10
	// ConnReadLimitSize 最多从连接中读取的数据量
	ConnReadLimitSize = 1 << 20
)

// HTTP Methods
const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

// HTTP Request Header Keys
const (
	// ContentLength 请求正文长度
	ContentLength = "Content-Length"
	// TransferEncoding 请求体编码格式
	TransferEncoding = "Transfer-Encoding"
	Expect           = "Expect"
)

// HTTP Request Header Values
const (
	ExpectValue = "100-continue"
)

const (
	// 允许客户端发送请求体的报文
	continueMessage = "HTTP/1.1 100 Continue\r\n\r\n"
)
