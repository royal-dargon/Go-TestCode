package main

import "fmt"

type cube struct {
	lenth int
	wide  int
	heigh int
}

func volume (x,y,z int) int {
	return x * y *z
}

func area (x,y,z int) int {
	s1 := x * y
	s2 := x * z
	s3 := z * y
	s4 := 2 * s1 + 2 * s2 + 2 * s3
	return s4
}

func main() {
	var cube1 cube
	fmt.Scanf("%d",&cube1.lenth)
	fmt.Scanf("%d",&cube1.wide)
	fmt.Scanf("%d",&cube1.heigh)
	volume1 := volume(cube1.lenth,cube1.wide,cube1.heigh)
	area1   := area(cube1.lenth,cube1.wide,cube1.heigh)
	fmt.Printf("volume : %d ; area : %d\n",volume1,area1)
}
