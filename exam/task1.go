package main

import "fmt"

func Replace (s string, target, replace byte) string {
	str := []byte(s)
	for i := 0; i < len(s) ; i ++ {
		if str[i] == target {
			str[i] = replace
		}
	}
	sreturn := string(str)
	return sreturn
}

func main() {
	var s string
	var a byte = 's'
	var b byte = 'k'
	fmt.Printf("请输入一个字符串\n")
	fmt.Scanf("%s",&s)
	s2 := Replace(s,a,b)
	fmt.Println(s2)
}
