package rule

import (
	"bytes"
	"strings"
)

type enumRule struct {
	idx   int
	rule  string
	sName string
	fName string
	fType string
	vals  string

	Val map[string]string
	Pkg map[string]struct{}
}

func (*enumRule) Prio() int {
	return PrioOther
}

func (er *enumRule) Check() string {
	sb := &bytes.Buffer{}

	if strings.HasPrefix(er.fType, "*[]") {
		enumSlicePtrTmpl.Execute(sb, map[string]any{
			"rule":        er.rule,
			"index":       er.idx,
			"struct_name": er.sName,
			"field_name":  er.fName,
			"field_type":  er.fType,
			"enum_value":  er.vals,
		})
	} else if strings.HasPrefix(er.fType, "[]") {
		enumSliceTmpl.Execute(sb, map[string]any{
			"rule":        er.rule,
			"index":       er.idx,
			"struct_name": er.sName,
			"field_name":  er.fName,
			"field_type":  er.fType,
			"enum_value":  er.vals,
		})
	} else if strings.HasPrefix(er.fType, "*") {
		enumPtrTmpl.Execute(sb, map[string]any{
			"rule":        er.rule,
			"index":       er.idx,
			"struct_name": er.sName,
			"field_name":  er.fName,
			"field_type":  er.fType,
			"enum_value":  er.vals,
		})
	} else {
		enumTmpl.Execute(sb, map[string]any{
			"rule":        er.rule,
			"index":       er.idx,
			"struct_name": er.sName,
			"field_name":  er.fName,
			"field_type":  er.fType,
			"enum_value":  er.vals,
		})
	}

	return sb.String()
}

// NewEnumRule 创建枚举规则
//	[structName] : 结构体名
//	[fieldName]  : 字段名
//	[typeName]   : 字段类型
//	[rule]       : 规则信息
func NewEnumRule(structName, fieldName, typeName, rule string) *enumRule {
	index++
	return &enumRule{
		idx:   index,
		rule:  rule,
		sName: structName,
		fName: fieldName,
		fType: typeName,
		vals:  rule,

		Val: map[string]string{},
		Pkg: map[string]struct{}{"errors": {}},
	}
}
