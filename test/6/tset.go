package main

import "fmt"

/*func main() {
	arr := []int{2, 7, 4, 5}
	x := Done(arr...)
	fmt.Println(x)
}

func Done(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	res := 1
	for _, n := range a {
		res = res * n
	}
	return res
}

func main() {
	fmt.Println("Start")
	defer Write()
	defer Done(5)
	fmt.Println("This is a test")
}

func Write() {
	fmt.Println("End")
}

func Done(a int) {
	for i := 0; i < a; i++ {
		for v := 1; v <= a-i; v++ {
			fmt.Printf("*")
			if v == a-i {
				fmt.Printf("\n")
			}
		}
	}
}

//完成一个斐波那契数列的练习
func main() {
	var a int
	fmt.Println("请输入一个正整数。")
	fmt.Scanf("%d", &a)
	for i := 1; i <= a; i++ {
		var res int64
		res = arr(i)
		fmt.Printf("%d ", res)
		if i%10 == 0 {
			fmt.Printf("\n")
		}
	}
}
func arr(a int) int64 {
	var res int64
	if a <= 2 {
		res = 1
	} else {
		res = arr(a-2) + arr(a-1)
	}
	return res
}
*/
//练习题用递归的方式从n打印到1
func main() {
	var a int
	fmt.Println("请输入一个正整数")
	fmt.Scanf("%d", &a)
	Pri(a)
}
func Pri(a int) {
	if a == 1 {
		fmt.Printf("%d ", a)
	} else {
		fmt.Printf("%d ", a)
		Pri(a - 1)
	}
}
