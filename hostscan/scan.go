package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var choice int
	fmt.Println("1. portlari Tara")
	fmt.Println("2. dns adresi bul")
	fmt.Println("seciminizi yapin(1 veya 2)")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		portTara()
	case 2:
		dnsScan()
	default:
		fmt.Println("secim basarisiz, program sonlandiriliyor kuzen")
		os.Exit(1)
	}
}
func portTara() {
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

func dnsScan() {
	var host string
	fmt.Print("web sitesi: example.com => ")
	fmt.Scan(&host)
	ipAddresses, err := resolveDNS(host)

	if err != nil {
		fmt.Printf("DNS çözümü başarısız: %v\n", err)
		return
	}

	fmt.Printf("%s DNS çözümü:\n", host)
	for _, ipAddress := range ipAddresses {
		fmt.Println(ipAddress)
	}
}

func resolveDNS(host string) ([]net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	return ips, nil
}
