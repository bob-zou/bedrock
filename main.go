package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	_version bool
	version  = "<UNDEFINED>"
	commit   = "<UNDEFINED>"
)

var RootCmd = &cobra.Command{
	Use: "bedrock",
	RunE: func(c *cobra.Command, args []string) error {
		if _version {
			fmt.Printf("bedrock version %s, build %s\n", version, commit)
			return nil
		}
		return c.Usage()
	},
}

func init() {
	RootCmd.AddCommand(New)
	RootCmd.AddCommand(Upgrade)
	RootCmd.Flags().BoolVarP(&_version, "version", "v", false, "show version")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
