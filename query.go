package main

import (
	"net/url"
	"strings"
)

// @Title       query.go
// @Description Request Query
// @Author      Zero.
// @Create      2024-08-15 18:21

// Query 表示 HTTP Request Query
type Query map[string]string

// Query 查询 Query 值
func (req *Request) Query(key string) string {
	return req.queryParams[key]
}


// 解析 HTTP Url 中的Query参数，方便后续获取
func (req *Request) requestQueryParse() (e error) {
	// 将字符串形式的URI 转换为 url.URL形式
	req.Url, e = url.ParseRequestURI(req.RequestURI)
	if e != nil {
		return
	}
	// 解析 Query 参数
	req.queryParams, e = queryParamsParse(req.Url.RawQuery)
	return
}

// HTTP Query 参数解析
// name=admin&age=18&address=xxx&...
func queryParamsParse(queryStr string) (Query,error) {
	items := strings.Split(queryStr, "&")
	query := make(Query, len(items))
	for _, item := range items {
		index := strings.IndexByte(item, '=')
		if index == -1 || index == len(item)-1 {
			continue
		}
		key := strings.TrimSpace(item[:index])
		value := strings.TrimSpace(item[index+1:])
		// 对Query的值进行解码，否则中文会乱码.
		v, err := url.QueryUnescape(value)
		if err != nil {
			return nil, err
		}
		query[key] = v
	}
	return query, nil
}
