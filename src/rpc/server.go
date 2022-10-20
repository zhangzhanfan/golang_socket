package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

//该方法向外暴露，提供计算圆形面积的服务
func (mu *MathUtil) CalculateCireArea(request float32, response *float32) error {
	*response = math.Pi * request * request
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
	// 通过该函数把mathUtil中提供的服务注册到HTTP服务上，方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()
	// 在特定的接口上进行监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}
