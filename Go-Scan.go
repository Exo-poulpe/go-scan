package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Scan IP address")
	for i := 1; i < 255; i++ {

		//fmt.Println("192.168.1." + strconv.Itoa(i))
		go fmt.Println(Scan("192.168.1."+strconv.Itoa(i), "80"))
	}
	fmt.Println("end")
}

func Scan(addr string, port string) string {

	var result string
	conn, err := net.DialTimeout("tcp", addr+":"+port, time.Second*2)
	if err != nil {
		result = addr + " is down"
	}
	if conn != nil {
		result = addr + " is up"
	}
	return result
}
