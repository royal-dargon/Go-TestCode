package main

import (
	"fmt"
	"io"
)

func main() {
	var a [1000]int
	//定义k 为指数 e 为系数
	var k, e, count int
	for {
		_, err := fmt.Scanf("%d %d", &k, &e)
		if err != io.EOF {
			a[k] = e
		}
	}
	a[0] = 0
	for i := 1; i <= 1000; i++ {
		a[i-1] = a[i] * i
		//清空上一位的数
		a[i] = 0
		if a[i-1] != 0 {
			count = count + 1
		}
	}
	if count == 0 {
		fmt.Printf("0 0")
	}
	for i := 1000; i >= 0; i-- {
		if a[i] != 0 {
			fmt.Printf("%v %v", a[i], i)
			count--
			if count == 0 {
				fmt.Printf("  ")
			}
		}
	}
}
