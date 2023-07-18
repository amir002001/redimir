package main

import (
	"log"
	"net"
)

func main() {
	localAddress := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3632,
		Zone: "",
	}

	server, err := net.ListenTCP("tcp4", &localAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	log.Printf("listening on %d\n", localAddress.Port)

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go processClient(connection)
	}
}

func processClient(connection net.Conn) error {
	_, err := connection.Write([]byte("hello world"))

	return err
}
