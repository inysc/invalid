package rule

import "fmt"

type defaultRule struct {
	sName string // 结构体名
	fName string // 字段名
	fType string // 类型
	val   string // 默认值
	isPtr bool
}

var _ Ruler = &defaultRule{}

func NewDefaultRule(structName, fieldType, filedName, val string) *defaultRule {
	return &defaultRule{
		sName: structName,
		fName: filedName,
		fType: fieldType,
		val:   mvRefTag(val),
		isPtr: fieldType[0] == '*',
	}
}

func (dr *defaultRule) Meth() string {
	if dr.isPtr && dr.val != "nil" {
		return fmt.Sprintf(defaultPtrFuncStr, dr.sName, dr.fName, dr.fType[1:], dr.val)
	}
	return fmt.Sprintf(defaultNoPtrFuncStr, dr.sName, dr.fName, dr.val)
}

func (dr *defaultRule) Name() string {
	return fmt.Sprintf("_%s_invalid_default_", dr.fName)
}

const defaultNoPtrFuncStr = `
func (i *%[1]s) _%[2]s_default_() error {
	i.%[2]s = %[3]s
	return nil
}
`

const defaultPtrFuncStr = `
func (i *%[1]s) _%[2]s_default_() error {
	var tmp %[3]s = %[4]s
	i.%[2]s = &tmp
	return nil
}
`
