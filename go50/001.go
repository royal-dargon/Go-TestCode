package main

import "fmt"

func main() {
	var a  int
	var b  int
	var c  int
	fmt.Scanf("%d %d %d",&a,&b,&c)
	if a+b > c && a+c > b && b+c > a {
		if a == b && a == c {
			fmt.Printf("等边三角形\n")
		} else if  a == b || a == c || b == c {
			fmt.Printf("等腰三角形\n")
		} else {
			fmt.Printf("三角形\n")
		}
	} else {
		fmt.Printf("false\n")
	}
}
