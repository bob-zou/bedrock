package main

import (
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = ""
	app.Usage = "bedrock 新项目创建工具"
	app.UsageText = "项目名 [options]"
	app.HideVersion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "d",
			Value:       "",
			Usage:       "指定项目所在目录",
			Destination: &p.path,
		},
	}
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		app.Run([]string{"-h"})
		return
	}
	p.Name = os.Args[1]
	app.Action = runNew
	args := append([]string{os.Args[0]}, os.Args[2:]...)
	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
