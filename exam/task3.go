package main

import "fmt"

//定义结构体
type cube struct {
	lenth int
	wide  int
	heigh int
}

//定义一个求表面积的方法
func (tn *cube) addarea() int {
	s1 := tn.lenth * tn.wide
	s2 := tn.lenth * tn.heigh
	s3 := tn.wide   * tn.heigh
	s4 := 2 * s1 + 2 * s2 + 2 * s3
	return s4
}

//定义一个求体积的方法
func (ts *cube) addvolume() int {
	v1 := ts.lenth * ts.wide * ts.heigh
	return v1
}

func main() {
	test1 := new(cube)
	fmt.Printf("请分别输入长宽高\n")
	fmt.Scanf("%d %d %d",&test1.lenth,&test1.wide,&test1.heigh)
	fmt.Println(test1.addarea())
	fmt.Println(test1.addvolume())
}

