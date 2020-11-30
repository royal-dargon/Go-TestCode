package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
func World(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "World")
}
func main() {
	server := http.Server{
		// 端口号要开在 1024-65535 这个范围内,如果页面进不了，可能是端口号被占用
		// 我电脑8080就被占用了
		Addr: "127.0.0.1:1024",
	}
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/world", World)
	server.ListenAndServe()
}
