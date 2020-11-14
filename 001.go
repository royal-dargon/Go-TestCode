package main

var a string

func main() {
	a = "G"
	println(a)
	f1()
}

func f1() {
	a := "o"
	println(a)
	f2()
}

func f2() {
	println(a)
}
