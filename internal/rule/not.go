package rule

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

type notRule struct {
	idx   int
	rule  string
	sName string // 结构体名
	fName string // 字段名
	fType string
	val   string

	Val map[string]string
	Pkg map[string]struct{}
}

var _ Ruler = &notRule{}

func (nr *notRule) Prio() int {
	if nr.fType[0] == '*' {
		return PrioNotNil
	}
	return PrioNot
}

func (nr *notRule) Name() string {
	return fmt.Sprintf("_%s_invalid_not_%d_", nr.fName, nr.idx)
}

func (nr *notRule) Check() string {
	sb := &bytes.Buffer{}
	if strings.HasPrefix(nr.fType, "*") && nr.val == "nil" {
		notTmpl.Execute(sb, map[string]any{
			"rule":            nr.rule,
			"index":           nr.idx,
			"struct_name":     nr.sName,
			"field_name":      nr.fName,
			"field_type":      nr.fType,
			"forbidden_value": nr.val,
		})
	} else if strings.HasPrefix(nr.fType, "*[]") && nr.val != "nil" {
		notSlicePtrTmpl.Execute(sb, map[string]any{
			"rule":            nr.rule,
			"index":           nr.idx,
			"struct_name":     nr.sName,
			"field_name":      nr.fName,
			"field_type":      nr.fType,
			"forbidden_value": nr.val,
		})
	} else if strings.HasPrefix(nr.fType, "*") && nr.val != "nil" {
		notPtrTmpl.Execute(sb, map[string]any{
			"rule":            nr.rule,
			"index":           nr.idx,
			"struct_name":     nr.sName,
			"field_name":      nr.fName,
			"field_type":      nr.fType,
			"forbidden_value": nr.val,
		})
	} else if strings.HasPrefix(nr.fType, "[]") {
		notSliceTmpl.Execute(sb, map[string]any{
			"rule":            nr.rule,
			"index":           nr.idx,
			"struct_name":     nr.sName,
			"field_name":      nr.fName,
			"field_type":      nr.fType,
			"forbidden_value": nr.val,
		})
	} else {
		notTmpl.Execute(sb, map[string]any{
			"rule":            nr.rule,
			"index":           nr.idx,
			"struct_name":     nr.sName,
			"field_name":      nr.fName,
			"field_type":      nr.fType,
			"forbidden_value": nr.val,
		})
	}
	return sb.String()
}

func NewNotRule(structName, fieldType, fieldName, rule string) *notRule {
	if rule[0] != '!' {
		log.Panicf("invalid not rule<%s>", rule)
	}

	index++
	return &notRule{
		idx:   index,
		rule:  rule,
		sName: structName,
		fName: fieldName,
		fType: fieldType,
		val:   rule[1:],

		Val: map[string]string{},
		Pkg: map[string]struct{}{"errors": {}},
	}
}
