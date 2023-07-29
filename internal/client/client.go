package client

import "fmt"

func SendParams(params []string) (result string, err error) {
	switch params[0] {
	case "get":
		if len(params) != 2 {
			return "", fmt.Errorf("wrong usage. Try get <key>")
		}
		return handleGet(params[1])
	case "set":
		if len(params) != 3 {
			return "", fmt.Errorf("wrong usage. Try set <key> <new value>")
		}
		return handleSet(params[1], params[2])
	case "delete":
		if len(params) != 2 {
			return "", fmt.Errorf("wrong usage. Try delete <key>")
		}
		return handleDelete(params[1])
	default:
		return "", fmt.Errorf("unkown command")
	}
}

func handleGet(key string) (result string, err error) {
	return "ok", nil
}

func handleSet(key, newValue string) (result string, err error) {
	return "ok", nil
}

func handleDelete(key string) (result string, err error) {
	return "ok", nil
}
