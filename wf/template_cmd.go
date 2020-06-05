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
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	RunE:    runRemplateCmd,
}

func runRemplateCmd(cmd *cobra.Command, args []string) error {
	jsonOutFlag, err := cmd.Flags().GetBool("json")
	if err != nil {
		return errors.Wrapf(err, "failed to get flag")
	}

	if len(args) == 0 {
		if err := printAvailableTemplates(); err != nil {
			return err
		}

		return nil
	}

	env := map[string]interface{}{}
	tplName := args[0]
	for _, param := range args[1:] {
		k, v := parseTemplateParam(param)
		env[k] = v
	}

	wfItem, err := loadTemplate(tplName, env)
	if err != nil {
		return errors.Wrap(err, "error loading template")
	}

	if jsonOutFlag {
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

func init() {
	templateCmd.Flags().Bool("json", false, "output as JSON")

	rootCmd.AddCommand(
		templateCmd,
	)
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
	wfDir, ok := os.LookupEnv("WF_DIR")
	if !ok {
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
