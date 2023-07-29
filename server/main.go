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

func try() {
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
	// read
	var header uint32

	if err := binary.Read(connection, binary.LittleEndian, &header); err != nil {
		return err
	} else if header > maxMessageSize {
		return fmt.Errorf("message too long")
	}

	message := make([]byte, header)
	if _, err := connection.Read(message); err != nil {
		return err
	}
	fmt.Printf("Client says: %s\n", string(message))

	// reply
	reply := "world"
	if err := binary.Write(connection, binary.LittleEndian, uint32(len(reply))); err != nil {
		return err
	}

	_, err := connection.Write([]byte(reply))
	return err
}
