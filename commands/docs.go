package commands

import (
	"github.com/spf13/cobra"
)

var Docs = &cobra.Command{
	Use:   "docs",
	Short: "init swagger docs",
	Long:  "tool for init swagger docs\nExample: \n  bedrock docs\n  bedrock docs --withoutInternal --withoutDependency\n  bedrock docs -d internal/server/http -g server.go --withoutInternal --withoutDependency",
	RunE:  docs,
}

var (
	_generalInfo       = "main.go"
	_dir               = "./cmd"
	_exclude           = ""
	_output            = "./docs"
	_withoutDependency = false
	_withoutInternal   = false
	_parseDepth        = 3
)

func init() {
	Docs.Flags().StringVarP(&_generalInfo, "generalInfo", "g", "main.go", "Go file path in which 'swagger general API Info' is written")
	Docs.Flags().StringVarP(&_dir, "dir", "d", "./cmd", "Directories you want to parse,comma separated and general-info file must be in the first one")
	Docs.Flags().StringVar(&_exclude, "exclude", "", "Exclude directories and files when searching, comma separated")
	Docs.Flags().StringVarP(&_output, "output", "o", "./docs", "Output directory for all the generated files(swagger.json, swagger.yaml and docs.go)")
	Docs.Flags().BoolVar(&_withoutDependency, "withoutDependency", false, "Do not parse go files inside dependency folder, disabled by default")
	Docs.Flags().BoolVar(&_withoutInternal, "withoutInternal", false, "Do not parse go files in internal packages, disabled by default")
	Docs.Flags().IntVar(&_parseDepth, "parseDepth", 3, "Dependency parse depth")
	Docs.Flags().BoolP("help", "h", false, "show help")
}

func docs(c *cobra.Command, args []string) (err error) {
	if err = installTools(); err != nil {
		return
	}
	cmdName := "swag"
	cmdArgs := []string{"init", "-d", _dir, "-o", _output, "-g", _generalInfo, "--parseDepth", "5", "--generatedTime"}
	if _exclude != "" {
		cmdArgs = append(cmdArgs, "--exclude", _exclude)
	}

	if !_withoutDependency {
		cmdArgs = append(cmdArgs, "--parseDependency")
	}

	if !_withoutDependency {
		cmdArgs = append(cmdArgs, "--parseInternal")
	}

	return execCmd(cmdName, cmdArgs...)
}
