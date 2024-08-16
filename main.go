package main

// @Title       main.go
// @Description Main
// @Author      Zero.
// @Create      2024-08-14 11:08



// Main
func main() {
	server := NewHTTPServer(":9898", new(echoHandler))
	panic(server.Startup())
}
