package main

import "fmt"

/*func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is: %c\n", i, str[i])
	}
	str2 := "日本话"
	fmt.Printf("The length of str2 is %d\n", len(str2))
	for i := 0; i < len(str2); i++ {
		fmt.Printf("Character on position %d is %c\n", i, str2[i])
	}
}

func main() {
	for i := 0; i < 25; i++ {
		for v := 0; v <= i; v++ {
			fmt.Printf("*")
			if v == i {
				fmt.Printf("\n")
			}
		}
	}
}
*/
func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
