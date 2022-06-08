package rule

import (
	"bytes"
	"fmt"
	"strings"
)

type RegexRule struct {
	idx   int
	rule  string
	sName string
	fName string
	fType string

	Val map[string]string
	Pkg map[string]struct{}
}

func (*RegexRule) Prio() int {
	return PrioOther
}

func (rr *RegexRule) Check() string {
	sb := &bytes.Buffer{}
	if strings.HasPrefix(rr.fType, "*[]") {
		regexSlicePtrTmpl.Execute(sb, map[string]any{
			"rule":        rr.rule,
			"index":       rr.idx,
			"struct_name": rr.sName,
			"field_name":  rr.fName,
			"field_type":  rr.fType,
			"regex_value": rr.rule[1:],
		})
	} else if strings.HasPrefix(rr.fType, "*") {
		regexPtrTmpl.Execute(sb, map[string]any{
			"rule":        rr.rule,
			"index":       rr.idx,
			"struct_name": rr.sName,
			"field_name":  rr.fName,
			"field_type":  rr.fType,
			"regex_value": rr.rule[1:],
		})
	} else if strings.HasPrefix(rr.fType, "[]") {
		regexSliceTmpl.Execute(sb, map[string]any{
			"rule":        rr.rule,
			"index":       rr.idx,
			"struct_name": rr.sName,
			"field_name":  rr.fName,
			"field_type":  rr.fType,
			"regex_value": rr.rule[1:],
		})
	} else {
		regexTmpl.Execute(sb, map[string]any{
			"rule":        rr.rule,
			"index":       rr.idx,
			"struct_name": rr.sName,
			"field_name":  rr.fName,
			"field_type":  rr.fType,
			"regex_value": rr.rule[1:],
		})
	}
	return sb.String()
}

func NewRegexRule(structName, fieldType, fieldName, rule string) *RegexRule {
	index++
	vname := fmt.Sprintf("regex_%s_%d *regexp.Regexp", strings.ToLower(structName), index)
	val := fmt.Sprintf(`regex_%s_%d, err = regexp.Compile(%s)
			if err != nil {
				panic(err)
			}
		`, strings.ToLower(structName), index, rule[1:])
	return &RegexRule{
		idx:   index,
		rule:  rule,
		sName: structName,
		fName: fieldName,
		fType: fieldType,
		Val:   map[string]string{vname: val},
		Pkg:   map[string]struct{}{"errors": {}, "regexp": {}},
	}
}
