// @Title chunk.go
// @Description Chunk 编码处理
// @Author Zero - 2024/8/15 23:31:38

package main

import (
	"bufio"
	"errors"
	"io"
)

// Chunk 编码
// HTTP/1.1 引入了新的编码方式 ———— `Chunk编码` 顾名思义就是将报文主体分块传输.
// 根据 Content-Length 解析请求体存在一个问题：必须事先知道请求体的长度，而有时候我们希望数据边产生边发送，实现流式传输,所以这样没有办法知道要发送多少数据。
// 因此在 HTTP/1.1 除了长连接之外的另一个改进就是引入了 Chunk编码, 需要在请求头中设置 `Transfer-Encoding: chunked`
//
// 格式如下:
// HTTP/1.1 200 OK\r\n
// Content-Type: text/plain\r\n
// Transfer-Encoding: chunked\r\n
// \r\n
// 以下为body
// 17\r\n							#chunk size
// hello, this is chunked \r\n		#chunk data
// D\r\n							#chunk size
// data sent by \r\n				#chunk data
// 7\r\n							#chunk size
// client!\r\n						#chunk data
// 0\r\n\r\n						#end

// 分块输出流，用于读取 Chunk 编码格式的请求体
type chunkReader struct {
	n    int           // 当前处理的块剩多少字节未读取完
	bufR *bufio.Reader // 请求体的读缓冲
	done bool          // 请求体是否读取完毕
	crlf [2]byte       // 读取 '\r\n' 回车换行符
}

// 从请求体中读取数据
func (cr *chunkReader) Read(buf []byte) (n int, err error) {
	if cr.done {
		// 已读完
		return 0, io.EOF
	}
	if cr.n == 0 {
		// 当前块已处理完，处理下一块
		cr.n, err = cr.NextChunkSize()
		if err != nil {
			return
		}
	}
}

// NextChunkSize 获取请求体下一个数据块大小
func (cr *chunkReader) NextChunkSize() (size int, e error) {
	// 读取一行内容，该内容为 ChunkSize,也就是数据块的长度
	line, err := readLine(cr.bufR)
	if err != nil {
		return
	}
	// TODO fix 待优化
	// 获取行的长度，并且将16进制 --> 10进制
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch  {
		case 'a' <= c && c <= 'f':
			size = size * 16 + int(c - 'a') + 10
		case 'A' <= c && c <= 'F':
			size = size * 16 + int(c - 'A') + 10
		case '0' <= c && c <= '9':
			size = size * 16 + int(c - '0')
		default:
			return 0, errors.New("illegal hex number")
		}
	}
	return
}