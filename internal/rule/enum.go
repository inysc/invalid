package rule

import "fmt"

type EnumRule struct {
	idx        int
	StructName string
	FieldName  string
	Type       string // 必须是值类型，而非指针类型
	IsPtr      bool
	Vals       string
}

func (er *EnumRule) Meth() string {
	if er.IsPtr {
		return fmt.Sprintf(enumPtrFuncStr, er.StructName, er.FieldName, er.idx, er.Type, er.Vals, "%v")
	}
	return fmt.Sprintf(enumNoPtrFuncStr, er.StructName, er.FieldName, er.idx, er.Type, er.Vals, "%v")
}

func (er *EnumRule) Name() string {
	return fmt.Sprintf("_%s_invalid_enum_%d_", er.FieldName, er.idx)
}

// NewEnumRule 创建枚举规则
//	[structName] : 结构体名
//	[fieldName]  : 字段名
//	[typeName]   : 字段类型
//	[rule]       : 规则信息
func NewEnumRule(structName, fieldName, typeName, rule string) *EnumRule {
	index++
	return &EnumRule{
		idx:        index,
		StructName: structName,
		FieldName:  fieldName,
		Type:       If(typeName[0] == '*', typeName[1:], typeName),
		IsPtr:      typeName[0] == '*',
		Vals:       rule,
	}
}

const enumNoPtrFuncStr = `
func (i *%[1]s) _%[2]s_invalid_enum_%[3]d_() error {
	for _, v := range []%[4]s%[5]s {
		if i.%[2]s == v {
			return nil
		}
	}
	return fmt.Errorf("invalid<enum>: %[1]s.%[2]s must not be %[6]s", i.%[2]s)
}
`

const enumPtrFuncStr = `
func (i *%[1]s) _%[2]s_invalid_enum_%[3]d_() error {
	if i.%[2]s == nil {
		return errors.New("invalid<enum>: %[1]s.%[2]s must not be nil")
	}
	for _, v := range []%[4]s%[5]s {
		if *i.%[2]s == v {
			return nil
		}
	}
	return fmt.Errorf("invalid<enum>: %[1]s.%[2]s must not be %[6]s", i.%[2]s)
}
`
