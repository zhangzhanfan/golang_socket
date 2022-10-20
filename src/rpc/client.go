package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	var request float32
	request = 10
	var response *float32
	// 同步的请求调用
	// err = client.Call("MathUtil.CalculateCireArea", request, &response)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(*response)

	// 异步的请求调用
	syncCall := client.Go("MathUtil.CalculateCireArea", request, &response, nil) //第四个参数为通道，Done:nil
	replayDone := <-syncCall.Done                                                //通过通道来感知异步调用是否结束，当Done有值才可以读取，继续执行下面的代码；否则一直阻塞状态
	fmt.Println(replayDone)
	fmt.Println(*response)
}
