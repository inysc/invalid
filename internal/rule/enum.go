package rule

import (
	"bytes"
	"fmt"
)

type EnumRule struct {
	idx   int
	rule  string
	sName string
	fName string
	fType string
	IsPtr bool
	Vals  string
}

func (er *EnumRule) Meth() string {
	sb := &bytes.Buffer{}
	if er.IsPtr {
		enumPtrTmpl.Execute(sb, map[string]any{
			"rule":        "",
			"index":       2,
			"struct_name": "Pill",
			"field_name":  "ID",
			"field_type":  "*int64",
			"enum_value":  "{234, 123}",
		})
	} else {
		enumTmpl.Execute(sb, map[string]any{
			"rule":        er.rule,
			"index":       er.idx,
			"struct_name": er.sName,
			"field_name":  er.fName,
			"field_type":  er.fType,
			"enum_value":  er.Vals,
		})
	}

	return sb.String()
}

func (er *EnumRule) Name() string {
	return fmt.Sprintf("_%s_invalid_enum_%d_", er.fName, er.idx)
}

// NewEnumRule 创建枚举规则
//	[structName] : 结构体名
//	[fieldName]  : 字段名
//	[typeName]   : 字段类型
//	[rule]       : 规则信息
func NewEnumRule(structName, fieldName, typeName, rule string) *EnumRule {
	index++
	return &EnumRule{
		idx:   index,
		rule:  rule,
		sName: structName,
		fName: fieldName,
		fType: typeName,
		IsPtr: typeName[0] == '*',
		Vals:  rule,
	}
}
