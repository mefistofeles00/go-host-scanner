package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var target string

	fmt.Print("hedef host (ornegin, 12.0.0.1): ")
	fmt.Scan(&target)
	fmt.Printf("'%s' uzerindeki portlar araniyor...\n", target)

	for port := 1; port <= 1024; port++ {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)
		if err != nil {
			continue
		}

		conn.Close()
		fmt.Printf("port %d acik\n", port)
	}
}
