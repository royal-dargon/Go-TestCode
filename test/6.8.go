package main

import "fmt"

func main() {
	f()
}

func f() {
	var a string = "hello,world"
	fv := func(a string) { fmt.Printf("%s",a) }
	fv(a)
	fmt.Printf("- g is of type %T and has value %v\n",fv,fv)
}
