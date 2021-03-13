package main

import(
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 1; i < 10; i ++{
		fmt.Println("A:", i)
	}
	wg.Done()
}

func b() {
	for i := 1; i < 10; i ++{
		fmt.Println("B:", i)
	}
	wg.Done()
}

func main(){
	runtime.GOMAXPROCS(8) //编辑占用CPU的核心数
	wg.Add(2)
	go a()
	go b()
	//time.Sleep(time.Second)
	wg.Wait()
}