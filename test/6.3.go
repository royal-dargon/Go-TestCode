package main

import "fmt"

func main() {
	F2(1,2,3,4)
	F3([]int{1,2,3,4})
}

func F2(s ... int) {
	for _, v :=range s {
		fmt.Println(v)
	}
}

func F3(s []int) {
	for _, v := range s {
		fmt.Println(v)
	}
}
