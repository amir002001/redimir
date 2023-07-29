package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "client",
	Short: "Initialize the client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server init")
	},
}
