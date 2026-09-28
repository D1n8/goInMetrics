package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
)

func main() {
	netConn, _ := net.Connections("all")
	fmt.Println(netConn)
}
