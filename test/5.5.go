package main

import "fmt"

func main() {
	var m int = 1
	for i:=1;i <= 25;i++ {
		for j := 1;j <= m;j++ {
			fmt.Printf("G")
		} 
		m++
		fmt.Printf("\n")
	}
}
