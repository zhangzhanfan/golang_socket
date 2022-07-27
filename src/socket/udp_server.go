package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen is failed err: ", err.Error())
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read is failed err: ", err.Error())
			continue
		}
		fmt.Printf("从客户端读到的数据:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP([]byte("ok"), addr)
		if err != nil {
			fmt.Println("wirte is failed err: ", err.Error())
			continue
		}
	}
}
