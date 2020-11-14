package main

import "fmt"

func main() {
	const (
	a = iota
	b
	c
	d
	e = 5
	f = 7
)
	fmt.Println(a,b,c,d,e,f)
}
	
