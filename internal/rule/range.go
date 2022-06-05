package rule

import (
	"fmt"
	"log"
	"strings"
)

// 对于字段为 nil 的 不做范围校验
type rangeRule struct {
	sName  string
	fName  string
	num    int
	left   string
	lVal   *string
	right  string
	layout string
	rVal   *string
	isNum  bool
	isPtr  bool

	Val map[string]string
	Pkg map[string]struct{}
}

var _ Ruler = &rangeRule{}

func NewRange(structName, fieldType, fieldName, rule string) *rangeRule {
	index++
	rr := &rangeRule{
		sName:  structName,
		fName:  fieldName,
		num:    index,
		left:   "",
		lVal:   nil,
		right:  "",
		layout: "",
		rVal:   nil,
		isNum:  true,
		isPtr:  fieldType[0] == '*',
		Val:    map[string]string{},
		Pkg:    map[string]struct{}{},
	}

	if rule[0] == '[' {
		rr.left = ">="
	} else if rule[0] == '(' {
		rr.left = ">"
	}

	if rule[len(rule)-1] == ']' {
		rr.right = "<="
	} else if rule[len(rule)-1] == ')' {
		rr.right = "<"
	}

	if strings.Count(rule, ",") != 1 {
		log.Panic("not supported<range time>")
	}

	idx := strings.IndexByte(rule, ',')
	if strings.TrimSpace(rule[1:idx]) != "" {
		rr.lVal = getr(rule[1:idx])
	}
	if strings.TrimSpace(rule[idx+1:len(rule)-1]) != "" {
		rr.rVal = getr(rule[idx+1 : len(rule)-1])
	}

	return rr
}

func (rr *rangeRule) Name() string {
	return fmt.Sprintf("_%s_invalid_range_%d_", rr.fName, rr.num)
}

func (rr *rangeRule) Meth() string {
	ret := fmt.Sprintf("func (i *%s)_%s_invalid_range_%d_() error {\n", rr.sName, rr.fName, rr.num)

	ret += rr.f1()
	ret += rr.f2("left", rRO(rr.left), rr.lVal)
	ret += rr.f2("right", rRO(rr.right), rr.rVal)

	ret += "return nil\n}\n\n"

	return ret
}

func (rr *rangeRule) f1() string {
	if rr.isPtr {
		return fmt.Sprintf(`if i.%s == nil {
			return nil
		}
		`, rr.fName)
	}
	return ""
}

func (rr *rangeRule) f2(loc, opt string, val *string) string {
	star := If(rr.isPtr, "*", "")

	stmt := ""
	if val != nil {
		if rr.isNum {
			stmt += fmt.Sprintf("if %si.%s %s %s {\n", star, rr.fName, opt, *val)
		} else {
			log.Panic("not supported<time range>")
			// vName := fmt.Sprintf("_time_%s_%d time.Time", loc, rr.num)

			// rr.Val[vName] = fmt.Sprintf(
			// 	`%s, err = time.Parse(%s, *i.%s)
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// `, vName, rr.layout, star)

			// rr.Pkg["time"] = struct{}{}

			// stmt += fmt.Sprintf(
			// 	`tm_%s, err := time.Parse(%s, %si.%s)`,
			// 	loc, rr.layout, star, rr.fName,
			// )
			// stmt += `if err != nil {
			// 	return err
			// `

			// tm, err := time.Parse("", "")
			// if err != nil {
			// 	return ""
			// }
			// if tm.Sub(tm) > 0 {
			// 	return ""
			// }
			// stmt += ""
		}
		stmt += fmt.Sprintf(
			`return errors.New("invalid<range>: %s.%s must not be %s %s")`,
			rr.sName, rr.fName, tRO(opt), *val,
		)
		stmt += "\n}\n\n"

	}

	return stmt
}
