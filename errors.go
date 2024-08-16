package main

import "errors"

// @Title       errors.go
// @Description 错误字典
// @Author      Zero.
// @Create      2024-08-16 18:12

var (
	ParseHTTPRequestErr = errors.New("parsing http request failed")
)
