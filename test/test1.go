package main

import "fmt"

func main() {
	sum := func() {
		var sum = 0.0
		for i := 1.0; i <= 1e6; i++ {
			sum += i
		}
		fmt.Println(sum)
	}
	sum()
}



