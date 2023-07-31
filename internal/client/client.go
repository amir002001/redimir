package client

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	redimirConnection *net.TCPConn
	redimirScanner    *bufio.Scanner
}

var ParseError = fmt.Errorf("error parsing response")

func Initialize() (*Client, error) {
	remoteAddress := net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 2807, // El Psy Kongroo
	}
	conn, err := net.DialTCP("tcp4", nil, &remoteAddress)
	if err != nil {
		return nil, err
	}
	return &Client{
		redimirConnection: conn,
		redimirScanner:    bufio.NewScanner(conn),
	}, nil
}

func (client *Client) SendParams(params []string) (result string, err error) {
	switch params[0] {
	case "get":
		if len(params) != 2 {
			return "", fmt.Errorf("wrong usage. try get <key>")
		}
		return client.handleGet(params[1])
	case "set":
		if len(params) != 3 {
			return "", fmt.Errorf("wrong usage. try set <key> <new value>")
		}
		return client.handleSet(params[1], params[2])
	case "del":
		if len(params) != 2 {
			return "", fmt.Errorf("wrong usage. try delete <key>")
		}
		return client.handleDelete(params[1])
	default:
		return "", fmt.Errorf("unkown command")
	}
}

func (client *Client) handleGet(key string) (result string, err error) {
	request := fmt.Sprintf("*2\r\n$3\r\nget\r\n$%d\r\n%s\r\n", len(key), key)
	_, err = client.redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}

	return client.parseResponse()
}

func (client *Client) handleSet(key, newValue string) (result string, err error) {
	request := fmt.Sprintf(
		"*3\r\n$3\r\nset\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n",
		len(key),
		key,
		len(newValue),
		newValue,
	)

	_, err = client.redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}

	return client.parseResponse()
}

func (client *Client) handleDelete(key string) (result string, err error) {
	request := fmt.Sprintf("*2\r\n$3\r\ndel\r\n$%d\r\n%s\r\n", len(key), key)
	_, err = client.redimirConnection.Write([]byte(request))
	if err != nil {
		return "", err
	}

	return client.parseResponse()
}

func (client *Client) parseResponse() (response string, err error) {
	client.redimirScanner.Scan()
	initial := client.redimirScanner.Text()
	if len(initial) == 0 {
		return "", ParseError
	}

	switch initial[0] {
	case '$':
		if initial[1] == '-' {
			return "(nil)", nil
		}
		client.redimirScanner.Scan()
		actualString := client.redimirScanner.Text()
		return actualString, nil
	case '+':
		return initial[1:], nil
	case ':':
		return initial[1:], nil
	default:
		return "", ParseError
	}
}
