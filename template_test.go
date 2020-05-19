package wf

import (
	"reflect"
	"strings"
	"testing"
)

func TestEvalTemplate(t *testing.T) {

	filename := "dummy.js"

	t.Run("eval simple template", func(t *testing.T) {
		r := strings.NewReader(`
			Item("2020-05-18");
		`)

		want := &Item{"2020-05-18", nil, nil}
		have, _ := EvalTemplate(filename, r, nil)

		if !reflect.DeepEqual(want, have) {
			t.Errorf("want: %+v; have: %+v", want, have)
		}
	})

	t.Run("eval parameterized template", func(t *testing.T) {
		r := strings.NewReader(`
			var date = new Date("2020-05-18");

			if (env.date === "tomorrow") {
				date.setDate(date.getDate() + 1);
			}

			var isoDate = date.toISOString().split("T")[0];

			Item(isoDate);
		`)

		want := &Item{"2020-05-19", nil, nil}
		have, err := EvalTemplate(
			filename,
			r,
			map[string]interface{}{
				"date": "tomorrow",
			},
		)

		if err != nil {
			t.Errorf("got error: %+v", err)
		}

		if !reflect.DeepEqual(want, have) {
			t.Errorf("want: %+v; have: %+v", want, have)
		}
	})
}
