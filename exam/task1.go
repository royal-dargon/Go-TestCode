package main

import "fmt"

func Replace (s string, target, replace byte) string {
	var s3 []string
	var s2 string
	for n,m := range s {
		s1 := []byte(s)
		if byte(m) == target {
			s1[n] = replace
		}
		s3 = append (s3,string(s1[n]))
	}
	s2 = string(s3)
	return s2
}

func main() {
	var s string
	var a byte
	var b byte
	fmt.Printf("请输入一个字符串\n")
	fmt.Scanf("%s",&s)
	fmt.Printf("请输入要替换的字符")
	fmt.Scanf("%c",&a)
	fmt.Printf("请输入替换后的字符")
	fmt.Scanf("%c",&b)
	s3 := Replace(s,a,b)
	fmt.Printf("%v\n",s3)
}
