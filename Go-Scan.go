/*
 * Athor : Poulpy
 * Date : 2019/02/04
 * Description : For scan IP address
 * Version : 1.0.0.0
 */

package main

import (
	. "fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// for sync Goroutines
var mu = &sync.Mutex{}

func ScanAddress(addr string) {
	var result string

	// lock Goroutines
	mu.Lock()
	conn, err := net.DialTimeout("tcp", addr+":80", time.Second)
	if err != nil {
		result = addr + "\t \033[31m [x] \033[0m"
	}
	if conn != nil {
		result = addr + "\t \033[32m [✔] \033[0m"
	}

	// unlock Goroutines
	mu.Unlock()
	Println(result)
}

func main() {

	var addr string

	// for timespan calc
	start := time.Now()

	for i := 1; i < 255; i++ {
		addr = "192.168.1." + strconv.Itoa(i)

		go ScanAddress(addr)
	}

	// lock Goroutines
	mu.Lock()
	// print time
	Println(time.Now().Sub(start))
	Println("End scan")

}
