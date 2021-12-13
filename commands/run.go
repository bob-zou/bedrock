package commands

import "github.com/spf13/cobra"

var Run = &cobra.Command{
	Use:   "run",
	Short: "run bedrock service",
	Long:  "run bedrock service\nExample: \n  bedrock run",
	RunE:  run,
}

func run(c *cobra.Command, args []string) (err error) {
	return execCmd("go", "run", "cmd/main.go")
}
