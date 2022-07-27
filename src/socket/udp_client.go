package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务器失,err:", err.Error())
		return
	}
	defer socket.Close()
	inputReader := bufio.NewReader(os.Stdin) //获取终端输入
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := socket.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("wirte失败,err:", err.Error())
			continue
		}
		var buf [1024]byte
		n, addr, err := socket.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read失败,err:", err.Error())
			continue
		}
		fmt.Printf("data:%v n:%v addr:%v\n", string(buf[:n]), n, addr)
	}
}
