package cmd

import "github.com/spf13/cobra"

func init() {
	rootCommand.AddCommand(clientCommand)
}

var rootCommand = &cobra.Command{
	Use:   "redimir",
	Short: "a redis implementation made by yours truly, Amir :3",
}

func Execute() error {
	return rootCommand.Execute()
}
