package main

import "fmt"

func main() {
	var a,b int 
	fmt.Printf("请输入两门成绩\n")
	fmt.Scanf("%d %d",&a,&b)
	switch {
		case a < 0 || a >100 :
			fmt.Printf("it is error\n")
		case b < 0 || b >100 :
			fmt.Printf("it is error\n")
		case a < 60 || b < 60 :
			fmt.Printf("it is not pass\n")
		case a >= 60 || b >= 60 :
			fmt.Printf("it is pass\n")
		}
	}
