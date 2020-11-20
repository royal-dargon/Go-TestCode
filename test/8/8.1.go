package main

import "fmt"

func main() {
	week := make(map[string]int )
	week["Monday"] = 1
	week["Tuesday"]= 2
	week["Wensday"]= 3
	if _, ok := week["Tuesday"]; ok {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
	if _,ok := week["Hollyday"]; ok {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}

