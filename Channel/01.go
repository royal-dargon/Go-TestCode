package main

import "fmt"

func main() {
	var ch1 chan int // 引用类型需要初始化
	// ch1 = make(chan int) 是同步通道，无缓冲区
	ch1 = make(chan int, 1) // 1是缓冲区，无缓冲区与有缓冲区是有区别的，无缓冲是不能暂存值
	ch1 <- 20
	x := <- ch1
	fmt.Println(x)
	//len(ch1)  获取通道中元素的数量
	//cap(ch1)  获得通道的容量
	close(ch1)
}