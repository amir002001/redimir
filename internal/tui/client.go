package tui

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/amir002001/redimir/internal/client"
	"github.com/knz/bubbline"
)

var invalidCommandError = fmt.Errorf("invalid command")

func InitializeClient() error {
	// Instantiate the widget.
	client, err := client.Initialize()
	if err != nil {
		return err
	}

	m := bubbline.New()

	m.Prompt = "redimir> "
	for {
		if val, err := m.GetLine(); err != nil {
			if err == io.EOF {
				// No more input.
				break
			}
			if errors.Is(err, bubbline.ErrInterrupted) {
				// Entered Ctrl+C to cancel input.
				fmt.Println("^C")
			} else {
				return err
			}
		} else {
			trimmedVal := strings.TrimSpace(val)
			params := strings.Split(trimmedVal, " ")
			if len(params[0]) == 0 {
				fmt.Println(invalidCommandError.Error())
				continue
			}

			if result, err := handleParams(client, params); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(result)
			}

			if err := m.AddHistory(val); err != nil {
				return err
			}
		}
	}
	return nil
}

func handleParams(client *client.Client, params []string) (result string, err error) {
	return client.SendParams(params)
}
