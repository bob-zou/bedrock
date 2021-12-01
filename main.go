package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var _showVersion bool

var RootCmd = &cobra.Command{
	Use: "bedrock",
	RunE: func(c *cobra.Command, args []string) error {
		if _showVersion {
			fmt.Printf("bedrock version %s\n", version)
			return nil
		}
		return c.Usage()
	},
}

func init() {
	RootCmd.AddCommand(New)
	RootCmd.AddCommand(Upgrade)
	RootCmd.Flags().BoolVarP(&_showVersion, "version", "v", false, "show version")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
