package main

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/knz/bubbline"
)

func main() {
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

		fmt.Println(val)
		// Handle regular input.
		err = m.AddHistory(val)
		if err != nil {
			log.Fatal(err)
		}
	}
}
