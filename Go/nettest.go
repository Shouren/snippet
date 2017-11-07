package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, addr := range addrs {
			fmt.Println(addr)
		}
	}
}
