package main

import "fmt"

func main() {
	var n string
	var sum int = 0
	mnew := make(map[int]string)
	mnew[1] = "yi"
	mnew[2] = "er"
	mnew[3] = "san"
	mnew[4] = "si"
	mnew[5] = "wu"
	mnew[6] = "liu"
	mnew[7] = "qi"
	mnew[8] = "ba"
	mnew[9] = "jiu"
	mnew[0] = "ling"
	fmt.Scanf("%s",&n)
	num  :=[]byte(n)
	for i := 0; i < len(n); i ++ {
		sum = sum + int(num[i]) - 48
	}
	if sum >= 100 && sum <= 999 {
		fmt.Printf("%v %v %v",mnew[sum/100],mnew[sum/10%10],mnew[sum%10])
	} else if sum >= 10 {
		fmt.Printf("%v %v",mnew[sum/10],mnew[sum%10])
	} else {
		fmt.Printf("%v",mnew[sum])
	}
}



