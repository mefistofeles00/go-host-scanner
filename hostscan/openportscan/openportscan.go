package openportscan

import (
	"fmt"
	"net"
	"os"
)

func ScanPort() {
	var port string
	var host string
	fmt.Print("port numarasi giriniz => ")
	fmt.Scan(&port)
	fmt.Print("host numarasi giriniz => ")
	fmt.Scan(&host)
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Printf("port %s kapali\n", port)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("port %s acik\n", port)
}
