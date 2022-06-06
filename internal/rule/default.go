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
	isPtr bool
}

var _ Ruler = &defaultRule{}

func NewDefaultRule(structName, fieldType, filedName, rule string) *defaultRule {
	return &defaultRule{
		rule:  rule,
		sName: structName,
		fName: filedName,
		fType: fieldType,
		val:   mvRefTag(rule),
		isPtr: fieldType[0] == '*',
	}
}

func (dr *defaultRule) Meth() string {
	sb := &bytes.Buffer{}
	if dr.isPtr && dr.val != "nil" {
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
