package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Upgrade = &cobra.Command{
	Use:   "upgrade [version]",
	Short: "bedrock self upgrade",
	Long:  "",
	RunE:  upgrade,
}

func upgrade(c *cobra.Command, args []string) (err error) {
	var (
		path = "github.com/bob-zou/bedrock@latest"
	)
	if len(args) == 1 {
		path = fmt.Sprintf("github.com/bob-zou/bedrock@%s", args[0])
	}

	if err = execCmd("go", "install", path); err != nil {
		return
	}
	return
}
