package main

import "fmt"

func Season(int a) {
	switch i := a ;{
	case i >= 1 && i <= 3 :
		{
		fmt.Printf("这是春天！\n")
	}
	case i > 3 && i <= 6 : {
		fmt.Printf("这是夏天！\n")
	}
	case i > 6 && i <= 9 : {
		fmt.Printf("这是秋天!\n")
	}
	case i > 9 && i <= 12 : {
		fmt.Printf("这是冬天!\n")
	}
	default : fmt.Printf("不是正确格式!\n") 
}

func main() {
	var mouth int
	fmt.Scanf("%d",&mouth)
	Season(mouth)
}