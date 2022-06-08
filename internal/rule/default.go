package rule

import (
	"bytes"
	"fmt"
)

type defaultRule struct {
	rule  string // 对应的规则
	sName string // 结构体名
	fName string // 字段名
	fType string // 类型
	val   string // 默认值

	Val map[string]string
	Pkg map[string]struct{}
}

func (dr *defaultRule) Prio() int {
	return PrioDefault
}

var _ Ruler = &defaultRule{}

func NewDefaultRule(structName, fieldType, filedName, rule string) *defaultRule {
	return &defaultRule{
		rule:  rule,
		sName: structName,
		fName: filedName,
		fType: fieldType,
		val:   mvDefault(rule),

		Val: map[string]string{},
		Pkg: map[string]struct{}{"errors": {}},
	}
}

func (dr *defaultRule) Check() string {
	sb := &bytes.Buffer{}
	if dr.fType[0] == '*' && dr.val != "nil" {
		defaultPtrTmpl.Execute(sb, map[string]any{
			"rule":          dr.rule,
			"struct_name":   dr.sName,
			"field_name":    dr.fName,
			"field_type":    dr.fType,
			"default_value": dr.val,
		})
	} else {
		defaultTmpl.Execute(sb, map[string]any{
			"rule":          dr.rule,
			"struct_name":   dr.sName,
			"field_name":    dr.fName,
			"field_type":    dr.fType,
			"default_value": dr.val,
		})
	}

	return sb.String()
}

func (dr *defaultRule) Name() string {
	return fmt.Sprintf("_%s_invalid_default_", dr.fName)
}
