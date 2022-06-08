package rule

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

// 对于字段为 nil 的 不做范围校验
type rangeRule struct {
	idx    int
	sName  string
	fName  string
	fType  string
	left   string
	lVal   *string
	right  string
	layout string
	rVal   *string
	isNum  bool
	isPtr  bool
	isLen  bool

	Val map[string]string
	Pkg map[string]struct{}
}

var _ Ruler = &rangeRule{}

func NewRange(structName, fieldType, fieldName, rule string) *rangeRule {
	index++
	isLength := rule[0] == 'l'
	if isLength {
		rule = rule[1:]
	}
	rr := &rangeRule{
		sName:  structName,
		fName:  fieldName,
		fType:  fieldType,
		idx:    index,
		left:   "",
		lVal:   nil,
		right:  "",
		layout: "",
		rVal:   nil,
		isNum:  true,
		isPtr:  fieldType[0] == '*',
		isLen:  isLength,
		Val:    map[string]string{},
		Pkg:    map[string]struct{}{"errors": {}},
	}

	if rule[0] == '[' {
		rr.left = ">="
	} else if rule[0] == '(' {
		rr.left = ">"
	} else {
		log.Panicf("the unsupport rule<%s>", rule)
	}

	if rule[len(rule)-1] == ']' {
		rr.right = "<="
	} else if rule[len(rule)-1] == ')' {
		rr.right = "<"
	} else {
		log.Panicf("the unsupport rule<%s>", rule)
	}

	if strings.Count(rule, ",") != 1 {
		rr.isNum = false
		cont := strings.TrimSpace(rule[1:len(rule)-1]) + ","
		lastPos := 0
		vals := make([]*string, 0, 3)
		for idx, v := range cont {
			if v == ',' {
				val := strings.TrimSpace(cont[lastPos:idx])
				if val != "" {
					vals = append(vals, &val)
				} else {
					vals = append(vals, nil)
				}
				lastPos = idx + 1
			}
		}
		rr.lVal = vals[0]
		rr.layout = *vals[1]
		rr.rVal = vals[2]
		if rr.lVal != nil {
			vname := fmt.Sprintf("time_%s_left_%d time.Time", strings.ToLower(structName), index)
			val := fmt.Sprintf(`time_%s_left_%d, err = time.Parse(%s, %s)
					if err != nil {
						panic(err)
					}
				`, strings.ToLower(structName), index, rr.layout, *rr.lVal)
			rr.Val[vname] = val
			rr.Pkg["time"] = struct{}{}
		}
		if rr.rVal != nil {
			vname := fmt.Sprintf("time_%s_right_%d time.Time", strings.ToLower(structName), index)
			val := fmt.Sprintf(`time_%s_right_%d, err = time.Parse(%s, %s)
					if err != nil {
						panic(err)
					}
				`, strings.ToLower(structName), index, rr.layout, *rr.rVal)
			rr.Val[vname] = val
			rr.Pkg["time"] = struct{}{}
		}
	} else {
		rr.isNum = true
		idx := strings.IndexByte(rule, ',')
		lVal := strings.TrimSpace(rule[1:idx])
		if lVal != "" {
			rr.lVal = &lVal
		}
		rVal := strings.TrimSpace(rule[idx+1 : len(rule)-1])
		if rVal != "" {
			rr.rVal = &rVal
		}
	}

	return rr
}

func (*rangeRule) Prio() int {
	return PrioOther
}

func (rr *rangeRule) Check() string {
	sb := &bytes.Buffer{}
	isPtr := rr.fType[0] == '*'
	if rr.isLen {
		if isPtr {
			if rr.lVal != nil {
				rangeLengthPtrTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.left,
					"limit_value": *rr.lVal,
				})
			}
			if rr.rVal != nil {
				rangeLengthPtrTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"limit_value": *rr.rVal,
				})
			}
		} else {
			if rr.lVal != nil {
				rangeLengthTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.left,
					"limit_value": *rr.lVal,
				})
			}
			if rr.rVal != nil {
				rangeLengthTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"limit_value": *rr.rVal,
				})
			}
		}
	} else if strings.Contains(rr.fType, "[]") {
		if isPtr {
			if rr.lVal != nil {
				rangeSlicePtrTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.left,
					"limit_value": *rr.lVal,
				})
			}
			if rr.rVal != nil {
				rangeSlicePtrTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"limit_value": *rr.rVal,
				})
			}
		} else {
			if rr.lVal != nil {
				rangeSliceTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.left,
					"limit_value": *rr.lVal,
				})
			}
			if rr.rVal != nil {
				rangeSliceTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"limit_value": *rr.rVal,
				})
			}
		}
	} else if rr.isNum {
		if isPtr {
			if rr.lVal != nil {
				rangePtrTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.left,
					"limit_value": *rr.lVal,
				})
			}
			if rr.rVal != nil {
				rangePtrTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"limit_value": *rr.rVal,
				})
			}
		} else {
			if rr.lVal != nil {
				rangeTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.left,
					"limit_value": *rr.lVal,
				})
			}
			if rr.rVal != nil {
				rangeTmpl.Execute(sb, map[string]any{
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"limit_value": *rr.rVal,
				})
			}
		}
	} else {
		if isPtr {
			if rr.lVal != nil {
				rangeTimePtrTmpl.Execute(sb, map[string]any{
					"index":       rr.idx,
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"pos":         "left",
					"layout":      rr.layout,
					"limit_value": *rr.rVal,
				})
			}
			if rr.rVal != nil {
				rangeTimePtrTmpl.Execute(sb, map[string]any{
					"index":       rr.idx,
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"pos":         "right",
					"layout":      rr.layout,
					"limit_value": *rr.rVal,
				})
			}
		} else {
			if rr.lVal != nil {
				rangeTimeTmpl.Execute(sb, map[string]any{
					"index":       rr.idx,
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"pos":         "left",
					"layout":      rr.layout,
					"limit_value": *rr.rVal,
				})
			}
			if rr.rVal != nil {
				rangeTimeTmpl.Execute(sb, map[string]any{
					"index":       rr.idx,
					"struct_name": rr.sName,
					"field_name":  rr.fName,
					"field_type":  rr.fType,
					"opt":         rr.right,
					"pos":         "right",
					"layout":      rr.layout,
					"limit_value": *rr.rVal,
				})
			}
		}
	}

	return sb.String()
}
