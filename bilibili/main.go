package main

import (
	"fmt"
	"sync"
	"time"

	//"strconv"
	"bilibili/model"
)

var wg sync.WaitGroup

/*
func main() {
	DB := model.Initdb()
	defer DB.Close()
	t1 := time.Now()
	wg.Add(100)
	for Id := 100; Id <= 200; Id++ {
		go model.WorkerOne(Id, wg)
	}
	wg.Wait()
	t2 := time.Since(t1)
	fmt.Println(t2)

}

*/

func main() {
	//var url1 chan int
	Url1 := make(chan int, 101)
	DB := model.Initdb()
	defer DB.Close()
	t1 := time.Now()

	// 开goroutine
	for i := 0; i < 50; i++ {
		go model.WorkerTwo(Url1)
	}
	// 发布任务
	for Id := 100; Id <= 200; Id++ {
		Url1 <- Id
		//wg.Add(1)
	}
	//close(url1)
	t2 := time.Since(t1)
	fmt.Println(t2)
	//wg.Wait()
	time.Sleep(time.Second * 5)
}
