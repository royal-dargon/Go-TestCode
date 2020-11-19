package main

import "fmt"

func main() {
	var a int
	var n int = 0
	fmt.Printf("请输入一个整数:")
	fmt.Scanf("%d",&a)
	for a != 1  {
		if a == 1 {
			break;
		} else if a % 2 == 0 {
			a = a/2
			n ++
		} else {
			a = (3*a+1)/2
			n++
		}
	}
	fmt.Printf("%d\n",n)
}


