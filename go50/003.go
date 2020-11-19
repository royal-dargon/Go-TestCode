package main

import "fmt"

func main() {
	var a,b int
	var n int = 1
	fmt.Printf("请输入一个奇数:")
	fmt.Scanf("%d",&a)
	b = a
	for i := 1; i <= a; i += 2 {
		for m := 1; m <= n; m++ {
			fmt.Printf(" ")
		}
		for j := b; j > 0; j-- {
			fmt.Printf("*")
		}
		n += 1
		b -= 2
		fmt.Printf("\n")
	}
	n = (a+1)/2 - 1 
	b = 3
	for i := 3; i <= a; i += 2 {
		for m := 1; m <= n; m++ {
			fmt.Printf(" ")
		}
		for j := b; j > 0; j-- {
			fmt.Printf("*")
		}
		n -= 1
		b += 2
		fmt.Printf("\n")
	}

}
