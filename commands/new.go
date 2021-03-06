package commands

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/fatih/color"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
	"github.com/toolkits/file"
)

type project struct {
	Year      string // current year
	Name      string // project name
	ModPrefix string // mod prefix
	path      string // project dir
}

var (
	_p        project // project object
	_httpOnly bool    // enable in next version
	_grpcOnly bool    // enable in next version
)

var New = &cobra.Command{
	Use:   "new [project name]",
	Short: "create a bedrock project",
	Long:  "create a project using the repository template\nExample: \n  bedrock new\n  bedrock new bedrock-service\n  bedrock new bedrock-service -d /tmp",
	RunE:  newProject,
}

func init() {
	New.Flags().StringVarP(&_p.path, "dir", "d", "", "specific dir of project")
	New.Flags().BoolVar(&_httpOnly, "http", false, "use http only, not support yet")
	New.Flags().BoolVar(&_grpcOnly, "grpc", false, "use grpc only, not support yet")
}

func newProject(c *cobra.Command, args []string) (err error) {
	// 先按照依赖的工具
	if insErr := installTools(); insErr != nil {
		return
	}

	if len(args) != 1 {
		_p.Name = "bedrock-demo"
	} else {
		_p.Name = args[0]
	}

	if _p.path != "" {
		if _p.path, err = filepath.Abs(_p.path); err != nil {
			return
		}
		_p.path = filepath.Join(_p.path, _p.Name)
	} else {
		pwd, _ := os.Getwd()
		_p.path = filepath.Join(pwd, _p.Name)
	}

	if !isEmpty(_p.path) {
		color.Red("%s is not empty.\n", _p.path)
		return
	}

	_p.ModPrefix = modPath(_p.path)
	_p.Year = strconv.Itoa(time.Now().Year())

	// create a project
	if err = create(); err != nil {
		return err
	}

	return nil
}

func isEmpty(path string) bool {
	if !file.IsExist(path) {
		return true
	}

	f, _ := os.Open(path)
	defer func() { _ = f.Close() }()

	_, err := f.Readdirnames(1)
	if err == io.EOF {
		return true
	}
	return false
}

func modPath(p string) string {
	dir := filepath.Dir(p)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			content, _ := ioutil.ReadFile(filepath.Join(dir, "go.mod"))
			mod := regexpReplace(`module\s+(?P<name>[\S]+)`, string(content), "$name")
			name := strings.TrimPrefix(filepath.Dir(p), dir)
			name = strings.TrimPrefix(name, string(os.PathSeparator))
			if name == "" {
				return fmt.Sprintf("%s/", mod)
			}
			return fmt.Sprintf("%s/%s/", mod, name)
		}
		parent := filepath.Dir(dir)
		if dir == parent {
			return ""
		}
		dir = parent
	}
}

func regexpReplace(reg, src, temp string) string {
	result := []byte{}
	pattern := regexp.MustCompile(reg)
	for _, submatches := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, temp, src, submatches)
	}
	return string(result)
}

//go:generate packr2
func create() (err error) {
	box := packr.New("all", "../templates/http")
	if err = os.MkdirAll(_p.path, 0755); err != nil {
		return
	}
	for _, name := range box.List() {
		if _p.ModPrefix != "" && name == "go.mod.tmpl" {
			continue
		}
		tmpl, _ := box.FindString(name)
		i := strings.LastIndex(name, string(os.PathSeparator))
		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(_p.path, dir), 0755); err != nil {
				return
			}
		}
		if strings.HasSuffix(name, ".tmpl") {
			name = strings.TrimSuffix(name, ".tmpl")
		}
		if err = write(filepath.Join(_p.path, name), tmpl); err != nil {
			return
		}
	}

	if err = execCmd("go", "mod", "tidy"); err != nil {
		return
	}
	if err = execCmdNoneOutput("swag", "init", "--dir", "cmd"); err != nil {
		return
	}
	if err = execCmd("go", "generate", "-x", "./..."); err != nil {
		return
	}
	if err = execCmd("swag", "init", "-d", "cmd", "--parseInternal", "--parseDependency", "--parseDepth", "2"); err != nil {
		return
	}
	return
}

func execCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = _p.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func execCmdNoneOutput(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = _p.path
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run()
}

func write(path, tpl string) (err error) {
	data, err := parse(tpl)
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, data, 0644)
}

func parse(s string) ([]byte, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, _p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
