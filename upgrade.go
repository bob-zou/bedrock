package main

import "github.com/spf13/cobra"

var Upgrade = &cobra.Command{
	Use:   "upgrade [Module ...]",
	Short: "bedrock self upgrade",
	Long:  "",
	RunE:  upgrade,
}

func upgrade(c *cobra.Command, args []string) (err error) {
	if err = execCmd("go", "install", "github.com/bob-zou/bedrock"); err != nil {
		return
	}
	return
}
