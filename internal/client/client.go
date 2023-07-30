package client

// I KNOW THIS ISNT THREAD SAFE PLZ GO BACK TO REDDIT :3

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	redimirConnection *net.TCPConn
	redimirScanner    *bufio.Scanner
	ParseError        = fmt.Errorf("error parsing response")
)

func init() {
	remoteAddress := net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 2807, // El Psy Kongroo
	}
	conn, err := net.DialTCP("tcp4", nil, &remoteAddress)
	if err != nil {
		log.Fatal(err)
	}
	redimirConnection = conn
	redimirScanner = bufio.NewScanner(conn)
}

func SendParams(params []string) (result string, err error) {
	switch params[0] {
	case "get":
		if len(params) != 2 {
			return "", fmt.Errorf("wrong usage. try get <key>")
		}
		return handleGet(params[1])
	case "set":
		if len(params) != 3 {
			return "", fmt.Errorf("wrong usage. try set <key> <new value>")
		}
		return handleSet(params[1], params[2])
	case "del":
		if len(params) != 2 {
			return "", fmt.Errorf("wrong usage. try delete <key>")
		}
		return handleDelete(params[1])
	default:
		return "", fmt.Errorf("unkown command")
	}
}

func handleGet(key string) (result string, err error) {
	request := fmt.Sprintf("*2\r\n$3\r\nget\r\n$%d\r\n%s\r\n", len(key), key)
	_, err = redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}

	return parseResponse()
}

func handleSet(key, newValue string) (result string, err error) {
	request := fmt.Sprintf(
		"*3\r\n$3\r\nset\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n",
		len(key),
		key,
		len(newValue),
		newValue,
	)

	_, err = redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}

	return parseResponse()
}

func handleDelete(key string) (result string, err error) {
	request := fmt.Sprintf("*2\r\n$3\r\ndel\r\n$%d\r\n%s\r\n", len(key), key)
	_, err = redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}

	return parseResponse()
}

func parseResponse() (response string, err error) {
	redimirScanner.Scan()
	initial := redimirScanner.Text()
	if len(initial) == 0 {
		return "", ParseError
	}

	switch initial[0] {
	case '$':
		if initial[1] == '-' {
			return "(nil)", nil
		}
		redimirScanner.Scan()
		actualString := redimirScanner.Text()
		return actualString, nil
	case '+':
		return initial[1:], nil
	case ':':
		return initial[1:], nil
	default:
		return "", ParseError
	}
}
