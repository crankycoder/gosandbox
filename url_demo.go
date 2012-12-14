package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func main() {

	rawurl := "udp://someuser:somepassword@myhost.com:9001/2"
	parsed, _ := url.Parse(rawurl)

	host := parsed.Host
	port := 80

	if strings.Contains(host, ":") {
		tmp := strings.Split(host, ":")
		host = tmp[0]
		port, _ = strconv.Atoi(tmp[1])
	}

	fmt.Printf("Host: [%s] Port: [%d]\n", host, port)
	fmt.Printf("Hello, playground\n")
}
