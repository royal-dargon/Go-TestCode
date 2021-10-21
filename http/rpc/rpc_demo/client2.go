package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":8801")
	if err != nil {
		panic(err.Error())
	}

	var result *float32
	addParma := &param.AddParma{Args1: 1.2, Args2: 2.3}
	err = client.Call("MathUtil.Add", addParma, &result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("计算结果:", *result)
}
