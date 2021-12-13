package main

import (
	"fmt"
	"os"

	"github.com/bob-zou/bedrock/commands"

	"github.com/spf13/cobra"
)

var _showVersion bool

var RootCmd = &cobra.Command{
	Use:  "bedrock",
	Long: "bedrock: an simple tool for creating go microservices.",
	RunE: func(c *cobra.Command, args []string) error {
		if _showVersion {
			fmt.Printf("bedrock version %s\n", version)
			return nil
		}
		return c.Usage()
	},
}

func init() {
	RootCmd.AddCommand(commands.New)
	RootCmd.AddCommand(commands.Upgrade)
	RootCmd.AddCommand(commands.Docs)
	RootCmd.AddCommand(commands.Run)
	RootCmd.Flags().BoolVarP(&_showVersion, "version", "v", false, "show version")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
