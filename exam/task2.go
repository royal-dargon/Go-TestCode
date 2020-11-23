package main

import "fmt"

func Romantolnt(romanNum string)int {
	roma := make(map[string]int)
	var sum int = 0
	roma["I"] = 1
	roma["V"] = 5
	roma["X"] = 10
	roma["L"] = 50
	roma["C"] = 100
	roma["D"] = 500
	roma["M"] = 1000
	str := []byte(romanNum)
	for i := 0; i < len(romanNum); i ++ {
		if i < len(romanNum) - 1 {
			if string(str[i]) == "I" && string(str[i+1]) == "V" {
				sum = sum - 1
			} else if string(str[i]) == "I" && string(str[i+1]) == "X"{
				sum = sum - 1
			} else if string(str[i]) == "X" && string(str[i+1]) == "L"{
				sum = sum - 10
			} else if string(str[i]) == "x" && string(str[i+1]) == "C"{
				sum = sum - 10
			} else if string(str[i]) == "C" && string(str[i+1]) == "D"{
				sum = sum - 100
			} else if string(str[i]) == "C" && string(str[i+1]) == "M"{
				sum = sum - 100
			} else {
				sum = sum + roma[string(str[i])]
			}
		} else {
		sum = sum + roma[string(str[i])] 
		}
	}
	return sum
}

func main() {
	var s string
	fmt.Printf("输入一个罗马字符")
	fmt.Scanf("%s",&s)
	sum := Romantolnt(s)
	fmt.Println(sum)
}







