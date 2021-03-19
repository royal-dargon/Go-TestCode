package main


/*
两个goroutine
一个生成零到一百的数字发送到channel1中
第二个取出并计算平方，发送到channel2中
*/

import "fmt"

func f1(ch chan int){
	for i :=0; i < 100; i ++{
		ch <- i 
	}
	close(ch)
}

func f2(ch1 chan int, ch2 chan int){
	//第一个取值的方式
	for {
		temp, ok := <- ch1
		if !ok{
			break
		}
		ch2 <- temp * temp
	}
	close(ch2)
}
func main(){
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go f1(ch1)
	go f2(ch1, ch2)
	//从通道中取值的法二
	for ret := range ch2{
		fmt.Println(ret)
	}
}

