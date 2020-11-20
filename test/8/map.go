package main

import "fmt"

func main() {
	var a string
	var m1 map[string]int
	m1 = make(map[string]int,10) //初始化
	m1["张旷"] = 18
	m1["汤先宁"] = 18
	m1["钟雨农"] = 999
	m1["周科宇"] = 666
	fmt.Scanf("%s",&a)
	fmt.Printf("%d\n",m1[a])
}
