package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"wf"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var tplParams []string

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
	for _, arg := range tplParams {
		k, v := parseTemplateParam(arg)
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
	templateCmd.Flags().StringSliceVarP(&tplParams, "param", "p", []string{}, "")

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
