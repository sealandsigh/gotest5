package main

import (
	"fmt"
	"net"
	"strings"
)

func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}

func main() {
	ip, err := GetOutboundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
