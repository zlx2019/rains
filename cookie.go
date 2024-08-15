package main

import "strings"

// @Title       cookies.go
// @Description Cookie
// @Author      Zero.
// @Create      2024-08-15 18:29

// Cookie HTTP Cookie
type Cookie map[string]string

// Cookie 查询 Cookie 值
func (req *Request) Cookie(key string) string {
	if req.cookies == nil {
		// 首次查询cookie时解析
		req.cookieParse()
	}
	return req.cookies[key]
}

// 解析 Cookie 信息
func (req *Request) cookieParse() {
	if req.cookies != nil{
		return
	}
	req.cookies = make(Cookie)
	// 从请求头中获取 Cookie
	cookie,ok := req.Headers["Cookie"]
	if !ok {
		return
	}
	for _, line := range cookie {
		// cookie格式: name1=value1; name2=value2; name3=value3
		// 分割成多组 K-V
		items := strings.Split(strings.TrimSpace(line), ";")
		if len(items) == 1 && items[0] == "" {
			continue
		}
		// 处理每一组
		for i := 0; i < len(items); i++ {
			idx := strings.IndexByte(items[i], '=')
			if idx == -1 {
				continue
			}
			k := items[i][:idx]
			v := items[i][idx+1:]
			req.cookies[k] = v
		}
	}
}