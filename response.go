package main

// @Title       response.go
// @Description HTTP Response
// @Author      Zero.
// @Create      2024-08-14 11:15

// HTTP 响应流
type response struct {
	conn *Connection
}

// 为客户端连接分配响应流
func mountResponse(c *Connection) *response{
	return &response{c}
}

// ResponseWriter HTTP 响应流写入器
type ResponseWriter interface {
	Write([]byte) (int, error)
}

// 实现写入函数，向连接中的 写缓冲写入数据
func (r *response) Write(buf []byte) (int,error) {
	return r.conn.bufW.Write(buf)
}