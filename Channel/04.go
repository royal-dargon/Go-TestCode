package main

import "fmt"
// select

func main(){
	ch := make(chan int, 1)
	for i := 0; i < 10; i ++{
		select {
		case x := <- ch: //能发送
			fmt.Println(x)
		case ch <- i: // 能接收
		default :
			fmt.Println("什么都没干")
		}
	}
}