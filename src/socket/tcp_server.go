package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端

func Process(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read failed, err:", err)
			break
		}
		str := string(buf[:n])
		fmt.Println("我收到了客户端发来的数据：", str)
		conn.Write([]byte("OK")) //发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() //保持时刻建立连接
		if err != nil {
			fmt.Println("conn failed, err:", err)
			continue
		}
		go Process(conn)
	}
}
