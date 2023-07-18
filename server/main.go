package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

const (
	maxMessageSize uint32 = 4096
	headerSize     uint32 = 4
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

		if err = processClient(connection); err != nil {
			log.Fatal(err)
		}
	}
}

func processClient(connection net.Conn) error {
	for {
		if err := oneRequest(connection); err != nil {
			log.Println(err)
			break
		}
	}
	err := connection.Close()
	return err
}

func oneRequest(connection net.Conn) error {
	readBuffer := make([]byte, maxMessageSize)

	err := readFull(connection, readBuffer, headerSize)
	if err != nil {
		return err
	}

	len := binary.LittleEndian.Uint32(readBuffer[:4])
	if len > maxMessageSize {
		return fmt.Errorf("message too long")
	}

	err = readFull(connection, readBuffer[headerSize:], len)
	if err != nil {
		return err
	}
	fmt.Printf("Client says: %s\n", string(readBuffer[4:4+len]))
	return nil
}

func readFull(connection net.Conn, readBuffer []byte, readN uint32) error {
	return nil
}

func writeAll() {
	return nil
}
