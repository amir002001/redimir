package cmd

import (
	"log"

	"github.com/amir002001/redimir/internal/tui"
	"github.com/spf13/cobra"
)

var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "Initialize the client",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.InitializeClient(); err != nil {
			log.Fatal(err)
		}
	},
}
