package rule

import (
	"log"
)

var index = 0

type Ruler interface {
	Name() string // 函数名
	Meth() string // 对应的方法内容
}

// NewRules
//	st  结构体名
//	typ 字段类型名
//	tag tag内容
func NewRules(st, typ, name, tags string) (
	val map[string]string, pkgs map[string]struct{}, duck bool, rs []Ruler) {

	if len(tags) <= 2 {
		return
	}

	val = map[string]string{}
	pkgs = map[string]struct{}{}

	merge := func(vl map[string]string, pg map[string]struct{}) {
		for k, v := range vl {
			val[k] = v
		}
		for k := range pg {
			pkgs[k] = struct{}{}
		}
	}

	// isPtr := strings.HasPrefix(typ, "*")

	for _, tag := range splitTag(tags) {
		log.Printf("tag<%s>", tag)
		switch tag[0] {
		case '!': // 禁止值
			pkgs["errors"] = struct{}{}
			rs = append(rs, NewNotRule(st, typ, name, tag))
		case '{': // 枚举值
			pkgs["errors"] = struct{}{}
			pkgs["fmt"] = struct{}{}
			rs = append(rs, NewEnumRule(st, name, typ, tag))
		case '[', '(': // 区间限制
			pkgs["errors"] = struct{}{}
			r := NewRange(st, typ, name, tag)
			merge(r.Val, r.Pkg)
			rs = append(rs, r)
		case 'r': // 正则约束
			pkgs["errors"] = struct{}{}
		case '@': // 默认值
			rs = append(rs, NewDefaultRule(st, typ, name, tag))
		case ':': // 撒扽
			if tag == ":>" {
				duck = true
				continue
			}
			fallthrough
		default:
			log.Panicf("the rule<%s> that cannot be parsed (%s)", tags, tag)
		}
	}
	return
}
