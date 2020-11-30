package main

import(
	"fmt"
	"http/net"
)

//定义用户的结构体
type user struct {
	name  string
	fname string
	code  int
}

func main() {
	servr := http.Server{
		//开一个属于自己的端口
		Addr := 127.0.0.1:1024
	}
