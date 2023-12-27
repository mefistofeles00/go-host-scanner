package main

import (
	"fmt"
	"hostscan/openportscan"
	"hostscan/serverscan"
	"net"
	"os"
	"time"
)

func main() {
	var choice int
	fmt.Println("1. portlari Tara")
	fmt.Println("2. dns adresi bul")
	fmt.Println("3. Server admin panel scanner")
	fmt.Println("4. Acik port Tara")
	fmt.Println("seciminizi yapin(1  2 veya 3)")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		portTara()
	case 2:
		dnsScan()
	case 3:
		getServerScan()
	case 4:
		openportscan.ScanPort()
	default:
		fmt.Println("secim basarisiz, program sonlandiriliyor kuzen")
		os.Exit(1)
	}

}

func getServerScan() {
	var urller, adminlist, kaydet string
	var bftimeout float64
	fmt.Println(`
	##########################################
	#                                        #
	#        ServerScan Priv8 Exploit        #
	#                                        #
	##########################################
	`)

	/*	urller = os.Args[1]
		adminlist = os.Args[2]

		bftimeout, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Hata: Timeout degeri float sekilde olmadilir kuzen ")
			os.Exit(1)
		}
		kaydet = os.Args[4]*/
	fmt.Print("sitelerin listelerinin adini giriniz: ")
	fmt.Scan(&urller)

	fmt.Print("admin listesinin adini girin: ")
	fmt.Scan(&adminlist)

	fmt.Print("timweout degeri girin (float olmali orn: 100): ")
	fmt.Scan(&bftimeout)

	fmt.Print("cikan sonuclari kaydetmek istiyormusunuz? (E/H): ")
	fmt.Scan(&kaydet)
	serverscan.RunScanner(urller, adminlist, bftimeout, kaydet)
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
