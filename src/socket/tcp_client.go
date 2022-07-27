package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000") //连接
	if err != nil {
		fmt.Println("conn failed err:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin) //获取终端输入
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo)) //写入数据
		if err != nil {
			return
		}
		var buf [1024]byte
		n, err := conn.Read(buf[:]) //读出数据
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("服务端收到了你的数据并发送了", string(buf[:n]))
	}
}
