package rule

import (
	"bytes"
	_ "embed"
	"io/ioutil"
	"testing"
)

func TestDefalut(t *testing.T) {
	sb := &bytes.Buffer{}
	sb.WriteString("package demo\n\n")
	{
		defaultTmpl.Execute(sb, map[string]any{
			"rule":          "",
			"struct_name":   "Pill",
			"field_name":    "IDInt",
			"field_type":    "int",
			"default_value": "234",
		})

		defaultPtrTmpl.Execute(sb, map[string]any{
			"rule":          "",
			"struct_name":   "Pill",
			"field_name":    "ID",
			"field_type":    "*int64",
			"default_value": "234",
		})
	}

	{
		notTmpl.Execute(sb, map[string]any{
			"rule":            "",
			"index":           1,
			"struct_name":     "Pill",
			"field_name":      "IDInt",
			"field_type":      "int",
			"forbidden_value": "234",
		})

		notPtrTmpl.Execute(sb, map[string]any{
			"rule":            "",
			"index":           2,
			"struct_name":     "Pill",
			"field_name":      "ID",
			"field_type":      "*int64",
			"forbidden_value": "234",
		})
	}

	{
		enumTmpl.Execute(sb, map[string]any{
			"rule":        "",
			"index":       1,
			"struct_name": "Pill",
			"field_name":  "IDInt",
			"field_type":  "int",
			"enum_value":  "{234, 123}",
		})

		enumPtrTmpl.Execute(sb, map[string]any{
			"rule":        "",
			"index":       2,
			"struct_name": "Pill",
			"field_name":  "ID",
			"field_type":  "*int64",
			"enum_value":  "{234, 123}",
		})
	}

	err = ioutil.WriteFile("../../demo/pill_invalid.go", sb.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
