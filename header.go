package main

// @Title       header.go
// @Description HTTP 请求头
// @Author      Zero.
// @Create      2024-08-14 12:59

// Headers HTTP 请求头
type Headers map[string][]string

// Set 添加请求头
func (h Headers) Set(key, value string)  {
	h[key] = []string{value}
}

// Put 追加多个请求头的值
func (h Headers) Put(key string, value... string) {
	h[key] = append(h[key], value...)
}

// Get 获取请求头值
func (h Headers) Get(key string) string {
	if values, ok := h[key]; ok && len(values) > 0 {
		return values[0]
	}else {
		return ""
	}
}

// GetOrDef 获取请求头,不存在则返回默认值
func (h Headers) GetOrDef(key, defVal string) string {
	if val := h.Get(key); val != "" {
		return val
	}
	return defVal
}

// Exists 查询指定的值是否存在
func (h Headers) Exists(key string)bool{
	_, ok := h[key]
	return ok
}

// Del 删除请求头
func (h Headers) Del(key string) {
	delete(h, key)
}
