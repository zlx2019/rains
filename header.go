package main

// @Title       header.go
// @Description HTTP 请求头
// @Author      Zero.
// @Create      2024-08-14 12:59

// Header HTTP 请求头
type Header map[string][]string

// Set 添加请求头
func (h Header) Set(key, value string)  {
	h[key] = []string{value}
}

// Put 追加多个请求头的值
func (h Header) Put(key string, value... string) {
	h[key] = append(h[key], value...)
}

// Get 获取请求头值
func (h Header) Get(key string) string {
	if values, ok := h[key]; ok && len(values) > 0 {
		return values[0]
	}else {
		return ""
	}
}

// Del 删除请求头
func (h Header) Del(key string) {
	delete(h, key)
}
