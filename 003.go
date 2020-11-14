package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { println(a) }

func m() {
	a := "0"
	println(a)
}
