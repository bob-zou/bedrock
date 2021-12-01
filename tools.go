package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/toolkits/file"

	"github.com/fatih/color"
)

type Tool struct {
	Name         string   `json:"name"`
	Alias        string   `json:"alias"`
	Install      string   `json:"install"`
	Summary      string   `json:"summary"`
	Requirements []string `json:"requirements"`
	Platform     []string `json:"platform"`
	requires     []*Tool
}

var tools = []*Tool{
	{
		Name:         "wire",
		Alias:        "wire",
		Install:      "go install github.com/google/wire/cmd/wire@latest",
		Summary:      "",
		Requirements: []string{},
		Platform:     []string{"darwin", "linux", "windows"},
	},
}

func installTools() (err error) {
	for _, t := range tools {
		if err = t.install(); err != nil {
			return
		}
	}
	return
}

func findTool(name string) *Tool {
	for i, tool := range tools {
		if tool.Name == name {
			return tools[i]
		}
	}
	return nil
}

func (t *Tool) check() (ok bool) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(os.Getenv("HOME"), "go")
	}

	if !file.IsExist(filepath.Join(gopath, "bin", t.Name)) {
		return false
	}

	return true

}

func (t *Tool) install() (err error) {
	if t == nil {
		return
	}

	if t.check() {
		return
	}

	for _, r := range t.Requirements {
		if err = findTool(r).install(); err != nil {
			return
		}
	}

	cmds := strings.Split(t.Install, " ")
	if len(cmds) > 0 {
		err = execCmd(cmds[0], cmds[1:]...)
		if err != nil {
			color.Red("%s: install failed!", t.Name)
		} else {
			color.Green("%s: install success!", t.Name)
		}
	}

	return
}
