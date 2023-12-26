package serverscan

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func RunScanner(urller, adminlist string, bftimeout float64, kaydet string) {
	ac, err := readLines(urller)
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	ac1, err := readLines(adminlist)
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	fmt.Println(" [+] Tarama başladı!")
	fmt.Println("")

	for _, url := range ac {
		url = strings.TrimSpace(url)
		for _, admin := range ac1 {
			admin = strings.TrimSpace(admin)
			fmt.Printf("Taraniyor %s%s\n", url, admin)
			if scanURL(url+admin, bftimeout) {
				fmt.Printf("%s%s [ WordPress ]\n", url, admin)
				if kaydet == "E" {
					writeToFile("wordpress.txt", url+admin)
				}
			}
		}
	}

	fmt.Println("")
	fmt.Println(" [+] Tarama bitti!")
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func scanURL(url string, timeout float64) bool {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	type1 := "type=\"password\""
	if strings.Contains(bodyString, type1) {
		if strings.Contains(bodyString, "wp-submit") {
			return true
		}
		if strings.Contains(bodyString, "index.php?route=") {
			return true
		}
		if strings.Contains(bodyString, "joomla") {
			return true
		}
		return true
	}

	return false
}

func writeToFile(filename, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Dosya yazma hatası:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(content + "\n"); err != nil {
		fmt.Println("Dosya yazma hatası:", err)
	}
}
