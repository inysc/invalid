package rule

import (
	"bytes"
	"fmt"
	"strings"
)

type duckRule struct {
	idx   int
	rule  string
	sName string
	fName string
	fType string
}

var _ Ruler = &duckRule{}

func (dr *duckRule) Prio() int {
	return PrioOther
}

func (dr *duckRule) Name() string {
	return fmt.Sprintf("_%s_invalid_duck_%d_", dr.fName, dr.idx)
}

func (dr *duckRule) Check() string {
	sb := &bytes.Buffer{}
	if strings.HasPrefix(dr.fType, "*[]") {
		duckSlicePtrTmpl.Execute(sb, map[string]any{
			"index":       dr.idx,
			"rule":        dr.rule,
			"struct_name": dr.sName,
			"field_name":  dr.fName,
			"field_type":  dr.fType,
		})
	} else if strings.HasPrefix(dr.fType, "[]") {
		duckSliceTmpl.Execute(sb, map[string]any{
			"index":       dr.idx,
			"rule":        dr.rule,
			"struct_name": dr.sName,
			"field_name":  dr.fName,
			"field_type":  dr.fType,
		})
	} else {
		duckTmpl.Execute(sb, map[string]any{
			"index":       dr.idx,
			"rule":        dr.rule,
			"struct_name": dr.sName,
			"field_name":  dr.fName,
			"field_type":  dr.fType,
		})
	}
	return sb.String()
}

func NewDuckRule(structName, fieldType, fieldName, rule string) *duckRule {
	index++
	return &duckRule{
		idx:   index,
		rule:  rule,
		sName: structName,
		fName: fieldName,
		fType: fieldType,
	}
}
