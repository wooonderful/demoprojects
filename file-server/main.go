package main

import (
    "net/http"
	"net"
	"errors"
	"fmt"
)

func main() {
	ip, err := getClientIp()
	if err != nil {
		fmt.Println("can not get the ip address")
	}
	fmt.Printf("you can access the website by: http://%s \n", ip)

    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    http.ListenAndServe(":80", nil)
}

func getClientIp() (string ,error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!")
}