package main

import (
	"log"

	"github.com/amir002001/redimir/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
