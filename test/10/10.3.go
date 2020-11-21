package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
} // 定义一个二维point结构体

// 定义一个三维polar的结构体
type Polar struct {
	x,y,z float64
}

//定义一个平方的函数
func My_double(a float64)float64 {
	b := a * a
	return b
}

func main() {
	var a [3]Point
	var b [3]Polar 
	var c [3]float64
	var d [3]float64
	for i := 0;i < 1;i ++ {
		fmt.Scanf("%f",&a[i].x)
		fmt.Scanf("%f",&a[i].y)
		fmt.Scanf("%f",&b[i].x)
		fmt.Scanf("%f",&b[i].y)
		fmt.Scanf("%f",&b[i].z)
	}
	fmt.Println(a[0].x,a[0].y,b[0].x,b[0].z,b[0].y)
	for i := 0;i < 1;i ++ {
		c[i] = math.Sqrt(My_double(a[i].x) + My_double(a[i].y))
		d[i] = math.Sqrt(My_double(b[i].x) + My_double(b[i].y) +
		My_double(b[i].z))
	}
	for i := 0;i < 1;i ++ {
		fmt.Println(c[i],d[i])
	}
}



