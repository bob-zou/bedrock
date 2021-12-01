package main

import (
	"github.com/spf13/cobra"
)

var Upgrade = &cobra.Command{
	Use:   "upgrade [version]",
	Short: "bedrock self upgrade",
	Long:  "tool for self-upgrade the bedrock.\nExample: \n  bedrock upgrade\n  bedrock upgrade v1.0.4",
	RunE:  upgrade,
}

func upgrade(c *cobra.Command, args []string) (err error) {
	if err = installTools(); err != nil {
		return
	}
	return execCmd("go", "get", "-u", "github.com/bob-zou/bedrock")
}
