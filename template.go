package wf

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/dop251/goja"
	"github.com/pkg/errors"
)

var vm = goja.New()

func init() {
	vm.Set("MONDAY", 1)
	vm.Set("TUESDAY", 2)
	vm.Set("WEDNESDAY", 3)
	vm.Set("THURSDAY", 4)
	vm.Set("FRIDAY", 5)
	vm.Set("SATURDAY", 6)
	vm.Set("SUNDAY", 7)

	vm.Set("days", []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	})

	vm.Set("env", vm.NewObject())
}

var tplLib = `
	function Item(title, arg1, arg2) {
		if (arg1 instanceof Array) {
			note = null;
			items = arg1;
		} else {
			note = arg1;
			items = arg2;
		}

		// filter items
		if (items) {
			items = items.filter(function (item) {
				return item !== null;
			});
		}

		return {
			title: title,
			note: note || null,
			items: items,
		}
	}
`

// EvalTemplate evaluates JS templates with a provided environment.
func EvalTemplate(
	filename string,
	r io.Reader,
	env map[string]interface{},
) (*Item, error) {

	vmEnv := vm.Get("env").ToObject(vm)
	for k, v := range env {
		vmEnv.Set(k, v)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	confStr := buf.String()

	v, err := vm.RunString(tplLib + confStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to evaluate %q", filename)
	}

	out, err := v.ToObject(vm).MarshalJSON()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal json produced by goja vm")
	}

	var item Item
	if err := json.Unmarshal(out, &item); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall json")
	}

	return &item, nil
}
