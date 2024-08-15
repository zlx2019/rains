// @Title define.go
// @Description 常量
// @Author Zero - 2024/8/15 22:19:19

package main

const CRLF = "\r\n"
const CRLF2 = "\r\n\r\n"
const headSep = ':'

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
)
