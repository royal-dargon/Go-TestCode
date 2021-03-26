package main

import (
	"fmt"
	"time"

	//"strconv"
	"bilibili/model"
)

func main() {
	DB := model.Initdb()
	defer DB.Close()
	t1 := time.Now()
	for Id := 100; Id <= 200; Id++ {
		model.Worker(Id)
	}
	t2 := time.Since(t1)
	fmt.Println(t2)
}
