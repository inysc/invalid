package rule

import (
	"bytes"
	"fmt"
	"log"
)

type notRule struct {
	idx   int
	rule  string
	sName string // 结构体名
	fName string // 字段名
	fType string
	val   string
	isPtr bool
}

var _ Ruler = &notRule{}

func (nr *notRule) Name() string {
	return fmt.Sprintf("_%s_invalid_not_%d_", nr.fName, nr.idx)
}

func (nr *notRule) Meth() string {
	sb := &bytes.Buffer{}
	if nr.isPtr && nr.val != "nil" {
		notPtrTmpl.Execute(sb, map[string]any{
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
		isPtr: fieldType[0] == '*',
	}
}
