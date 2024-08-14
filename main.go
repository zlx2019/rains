package main

import "log/slog"

// @Title       main.go
// @Description Main
// @Author      Zero.
// @Create      2024-08-14 11:08


type myHandler struct {

}

func (*myHandler) Handle(writer ResponseWriter, request *Request)  {
	slog.Info("Hello")
}


// Main
func main() {
	server := NewHTTPServer(":9090", new(myHandler))
	panic(server.Startup())
}
