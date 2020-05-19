package main

import (
	"os"
	"path/filepath"
	"wf"

	"github.com/pkg/errors"
)

func loadTemplate(name string, env map[string]interface{}) (*wf.Item, error) {
	wfDir, ok := os.LookupEnv("WF_DIR")
	if !ok {
		return nil, errors.New("WF_DIR is not defined")
	}
	filename := filepath.Join(wfDir, "templates", name+".js")

	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file: %s", filename)
	}

	return wf.EvalTemplate(name+".js", f, env)
}
