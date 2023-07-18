package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	remoteAddr := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3632,
	}

	tcpConnection, err := net.DialTCP("tcp4", nil, &remoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	bytes := make([]byte, 1024)

	mLen, err := tcpConnection.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes[:mLen]))
}
