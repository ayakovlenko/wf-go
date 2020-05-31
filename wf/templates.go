package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"wf"

	"github.com/pkg/errors"
)

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
