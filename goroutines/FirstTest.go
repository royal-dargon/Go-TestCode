package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	wg.Add(10000) // 发送令牌
	for i := 0; i < 10000; i ++{
		go func(i int) {
			fmt.Println("hello",i)
			wg.Done() // 回收令牌
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait()   // 等待令牌全部回收
}