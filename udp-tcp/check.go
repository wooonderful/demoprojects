package main

import (
	"flag"
	"fmt"
	"net"
)

const version = "0.1.0"

var Input_protocol = flag.String("p", "tcp", "The protocol you want to check")

func tcp(url string) int {
	_, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		return 1
	}
}

func udp(url string) int {
	coon, err := net.Dial("udp", url)
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		coon.Close()
		return 1
	}
}

func main1() {
	flag.Parse()
	if flag.NArg() < 1 {
		useage := "使用示例: check -p tcp 192.168.7.26:22 或者 check -p udp 192.168.7.23:123 "
		fmt.Println(useage)
		return
	}
	p := *Input_protocol
	switch {
	case p == "tcp":
		r := tcp(flag.Args()[0])
		fmt.Println(r)
	case p == "udp":
		r := udp(flag.Args()[0])
		fmt.Println(r)
	}
}

func main() {
	r := udp("11.22.33.44:31900")
	fmt.Println(r)
}

// UDP 客户端
//func main() {
//	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
//		IP:   net.IPv4(172, 16, 0, 46),
//		Port: 29999,
//	})
//	if err != nil {
//		fmt.Println("连接UDP服务器失败，err: ", err)
//		return
//	}
//	defer socket.Close()
//}
