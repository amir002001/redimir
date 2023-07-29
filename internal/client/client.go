package client

import (
	"fmt"
	"log"
	"net"
)

var redimirConnection *net.TCPConn

func init() {
	remoteAddress := net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 2807, // El Psy Kongroo
		Zone: "",
	}
	conn, err := net.DialTCP("tcp4", nil, &remoteAddress)
	if err != nil {
		log.Fatal(err)
	}
	redimirConnection = conn
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
	case "delete":
		if len(params) != 2 {
			return "", fmt.Errorf("wrong usage. try delete <key>")
		}
		return handleDelete(params[1])
	default:
		return "", fmt.Errorf("unkown command")
	}
}

func handleGet(key string) (result string, err error) {
	request := fmt.Sprintf("*2\r\n$3get\r\n$%d%s", len(key), key)
	_, err = redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}
	// TODO HANDLE RESPONSE
	return "OK", nil
}

func handleSet(key, newValue string) (result string, err error) {
	request := fmt.Sprintf("*3\r\n$3set\r\n$%d%s\r\n%d%s", len(key), key, len(newValue), newValue)
	_, err = redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}
	// TODO HANDLE RESPONSE
	return "OK", nil
}

func handleDelete(key string) (result string, err error) {
	request := fmt.Sprintf("*2\r\n$3del\r\n$%d%s", len(key), key)
	_, err = redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}
	// TODO HANDLE RESPONSE
	return "OK", nil
}
