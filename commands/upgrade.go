package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Upgrade = &cobra.Command{
	Use:   "upgrade [version]",
	Short: "self upgrade",
	Long:  "tool for self-upgrade the bedrock.\nExample: \n  bedrock upgrade\n  bedrock upgrade v1.0.7",
	RunE:  upgrade,
}

func upgrade(c *cobra.Command, args []string) (err error) {
	if err = installTools(); err != nil {
		return
	}

	var (
		path = "github.com/bob-zou/bedrock@latest"
	)

	if len(args) == 1 {
		path = fmt.Sprintf("github.com/bob-zou/bedrock@%s", args[0])
	}

	return execCmd("go", "install", path)
}
