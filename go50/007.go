package main

import "fmt"

func main() {
	var arr1 [1000]int
	var n int
	var flag int = 1
	fmt.Printf("请输入一个整数")
	fmt.Scanf("%d",&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d",&arr1[i])
	}
	//输入数字
	for i := 0; i < n; i++ {
		if arr1[i] % 2 == 1 {
			flag = flag * arr1[i]
		}
	}
	fmt.Printf("%d\n",flag)
}

