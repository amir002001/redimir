package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

const (
	headerSize     = 4
	maxMessageSize = 4096
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

	if err := query(tcpConnection, "hello1"); err != nil {
		log.Fatal(err)
	}
	if err := query(tcpConnection, "hello2"); err != nil {
		log.Fatal(err)
	}
	if err := query(tcpConnection, "hello3"); err != nil {
		log.Fatal(err)
	}
}

func query(connection *net.TCPConn, message string) error {
	sendHeader := uint32(len(message))
	if err := binary.Write(connection, binary.LittleEndian, sendHeader); err != nil {
		return err
	}
	if _, err := connection.Write([]byte(message)); err != nil {
		return err
	}
	var receiveHeader uint32
	if err := binary.Read(connection, binary.LittleEndian, &receiveHeader); err != nil {
		return err
	}
	if receiveHeader > maxMessageSize {
		return fmt.Errorf("Message too big")
	}
	receiveMessage := make([]byte, receiveHeader)
	_, err := connection.Read(receiveMessage)
	fmt.Printf("server says: %s\n", receiveMessage)
	return err
}
