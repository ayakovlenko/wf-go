package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"wf"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var templateCmd = &cli.Command{
	Name:    "template",
	Aliases: []string{"tpl"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:     "json",
			Required: false,
		},
	},
	Action: runRemplateCmd,
	Subcommands: []*cli.Command{
		templatePathCmd,
	},
}

var templatePathCmd = &cli.Command{
	Name: "path",
	Action: func(cCtx *cli.Context) error {
		tplDir, err := locateTemplateDir()
		if err != nil {
			return err
		}

		fmt.Print(tplDir)

		return nil
	},
}

func runRemplateCmd(cCtx *cli.Context) error {
	env := map[string]interface{}{}
	tplName := cCtx.Args().First()
	params := cCtx.Args().Tail()
	for _, param := range params {
		k, v := parseTemplateParam(param)
		env[k] = v
	}

	wfItem, err := loadTemplate(tplName, env)
	if err != nil {
		return errors.Wrap(err, "error loading template")
	}

	if cCtx.Bool("json") {
		jsonOut, err := json.MarshalIndent(wfItem, "", "  ")
		if err != nil {
			return errors.Wrap(err, "failed to marshall to xml")
		}

		fmt.Println(string(jsonOut))
		return nil
	}

	outline := wf.ToOPML(wfItem)

	xmlOut, err := xml.MarshalIndent(outline, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshall to json")
	}

	fmt.Println(string(xmlOut))
	return nil
}

func printAvailableTemplates() error {
	tpls, err := getAvailableTemplates()
	if err != nil {
		return err
	}

	fmt.Printf("available templates:\n\n")
	for _, k := range tpls {
		fmt.Printf("- %s\n", k)
	}

	return nil
}

func parseTemplateParam(arg string) (key string, value interface{}) {
	idx := strings.Index(arg, "=")

	if idx == -1 {
		key = arg
		value = true
		return
	}

	key = arg[:idx]
	value = arg[idx+1:]
	return
}

func locateTemplateDir() (string, error) {
	wfDir, defined := os.LookupEnv("WF_DIR")
	if !defined {
		return "", errors.New("WF_DIR is not defined")
	}

	return filepath.Join(wfDir, "templates"), nil
}

func loadTemplate(name string, env map[string]interface{}) (*wf.Item, error) {
	tplDir, err := locateTemplateDir()
	if err != nil {
		return nil, err
	}

	filename := filepath.Join(tplDir, name+".js")

	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file: %s", filename)
	}

	return wf.EvalTemplate(name+".js", f, env)
}

func getAvailableTemplates() ([]string, error) {
	tplDir, err := locateTemplateDir()
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(tplDir)
	if err != nil {
		return nil, err
	}

	tpls := []string{}
	for _, f := range files {
		name := f.Name()

		if strings.HasSuffix(name, ".js") {
			tpls = append(tpls, strings.TrimSuffix(name, ".js"))
		}
	}

	return tpls, nil
}
