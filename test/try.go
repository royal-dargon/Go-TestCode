package main

import "fmt"

func main () {
	var i int = 1
	i = put_count(i)
	i = put_count(i)
	i = put_count(i)
	i = put_count(i)
	i = put_count(i)
}


func put_count(i int) ( int) {
	fmt.Printf("第%d次使用\n",i)
	i++
	return i
}
