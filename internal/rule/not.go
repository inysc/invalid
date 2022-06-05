package rule

import (
	"fmt"
	"log"
)

type notRule struct {
	idx   int
	sName string // 结构体名
	fName string // 字段名
	val   string
	isPtr bool
}

var _ Ruler = &notRule{}

func (nr *notRule) Name() string {
	return fmt.Sprintf("_%s_invalid_not_%d_", nr.fName, nr.idx)
}

func (nr *notRule) Meth() string {
	ret := fmt.Sprintf(
		"errors.New(`invalid<not>: %s.%s must not be %s`)",
		nr.sName, nr.fName, nr.val,
	)
	if nr.isPtr && nr.val != "nil" {
		return fmt.Sprintf(notRulePtrFunc,
			nr.sName, nr.fName, nr.idx, nr.val, ret)
	}
	return fmt.Sprintf(notRuleNoPtrFunc, nr.sName, nr.fName, nr.idx, nr.val, ret)
}

func NewNotRule(structName, fieldType, fieldName, rule string) *notRule {
	if rule[0] != '!' {
		log.Panicf("invalid not rule<%s>", rule)
	}

	index++
	return &notRule{
		idx:   index,
		sName: structName,
		fName: fieldName,
		val:   rule[1:],
		isPtr: fieldType[0] == '*',
	}
}

// NotRuleFuncStr
//	[1] : 结构体名
//	[2] : 字段名
//	[3] : 随机编号
//	[4] : 禁止值
const notRuleNoPtrFunc = `
func (i *%[1]s) _%[2]s_invalid_not_%[3]d_() error {
	if i.%[2]s == %[4]s {
		return %[5]s
	}
	return nil
}
`

const notRulePtrFunc = `
func (i *%[1]s) _%[2]s_invalid_not_%[3]d_() error {
	if i.%[2]s != nil && *i.%[2]s == %[4]s {
		return %[5]s
	}
	return nil
}
`
