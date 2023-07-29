package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/knz/bubbline"
)

var invalidCommandError = fmt.Errorf("invalid command")

func try() {
	// Instantiate the widget.
	m := bubbline.New()

	m.Prompt = "redimir> "
	for {
		// Read a line of input using the widget.
		val, err := m.GetLine()
		// Handle the end of input.
		if err != nil {
			if err == io.EOF {
				// No more input.
				break
			}
			if errors.Is(err, bubbline.ErrInterrupted) {
				// Entered Ctrl+C to cancel input.
				fmt.Println("^C")
			} else if errors.Is(err, bubbline.ErrTerminated) {
				fmt.Println("terminated")
				break
			} else {
				fmt.Println("error:", err)
			}
			continue
		}

		trimmedVal := strings.TrimSpace(val)
		params := strings.Split(trimmedVal, " ")
		if len(params[0]) == 0 {
			fmt.Println(invalidCommandError.Error())
			continue
		}

		if err := handleParams(params); err != nil {
			fmt.Println(err.Error())
		}

		if err := m.AddHistory(val); err != nil {
			log.Fatal(err)
		}
	}
}

func handleParams(params []string) error {
	switch params[0] {
	case "set":
		if len(params) != 3 {
			return invalidCommandError
		}
	case "get":
		if len(params) != 2 {
			return invalidCommandError
		}
	case "delete":
		if len(params) != 2 {
			return invalidCommandError
		}
	default:
		return invalidCommandError
	}

	return nil
}
