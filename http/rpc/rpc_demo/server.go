package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type AddParma struct {
	Args1 float32 //第一个参数
	Args2 float32 //第二个参数
}
type MathUtil struct {
}

//该方法向外暴露：提供计算圆形面积的服务
func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req //圆形的面积 s = π * r * r
	return nil                  //返回类型
}

func (mu *MathUtil) Add(param AddParma, resp *float32) error {
	*resp = param.Args1 + param.Args2 //实现两数相加的功能
	return nil
}

func main() {
	// 初始化指针数据类型
	mathUtil := new(MathUtil)

	// 调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err.Error())
	}

	// 将服务注册到HTTP协议上
	rpc.HandleHTTP()

	// 在特定的端口进行监听
	listen, err := net.Listen("tcp", ":8881")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}
