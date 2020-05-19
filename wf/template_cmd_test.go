package main

import (
	"testing"
)

func TestParseTemplateParam(t *testing.T) {

	t.Run("key-value param", func(t *testing.T) {
		wantKey, wantValue := "date", "tomorrow"

		haveKey, haveValue := parseTemplateParam("date=tomorrow")

		if haveKey != wantKey {
			t.Errorf("key: want: %s; have: %s", wantKey, haveKey)
		}

		if haveValue != wantValue {
			t.Errorf("value: want: %s; have: %s", wantValue, haveValue)
		}
	})

	t.Run("key-value param multi-world", func(t *testing.T) {
		wantKey, wantValue := "date", "after tomorrow"

		haveKey, haveValue := parseTemplateParam("date=after tomorrow")

		if haveKey != wantKey {
			t.Errorf("key: want: %s; have: %s", wantKey, haveKey)
		}

		if haveValue != wantValue {
			t.Errorf("value: want: %s; have: %s", wantValue, haveValue)
		}
	})

	t.Run("bool param", func(t *testing.T) {
		wantKey, wantValue := "tomorrow", true

		haveKey, haveValue := parseTemplateParam("tomorrow")

		if haveKey != wantKey {
			t.Errorf("key: want: %s; have: %s", wantKey, haveKey)
		}

		if haveValue != wantValue {
			t.Errorf("value: want: %+v; have: %+v", wantValue, haveValue)
		}
	})
}
