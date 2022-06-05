package model

import (
	"go/ast"
	"log"
	"sort"
	"strconv"
)

func taget(tag, key string) (value string) {
	if tag != "" && tag[0] == '`' {
		tag = tag[1 : len(tag)-1]
	}

	for tag != "" {
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		if key == name {
			value, err := strconv.Unquote(qvalue)
			if err != nil {
				break
			}
			return value
		}
	}
	return ""
}

func GetFiledType(expr ast.Expr) (ret string) {
	var ok bool
	var vTyp ast.Expr
	vTyp, ok = expr.(*ast.StarExpr) // 取出指针的基本类型
	if ok {
		ret += "*"
		vTyp = expr.(*ast.StarExpr).X
	} else {
		vTyp = expr
	}

	switch vt := vTyp.(type) {
	case *ast.ArrayType: // 切片 不支持指针切片
		ret += "[]"
		switch elt := vt.Elt.(type) {
		case *ast.Ident:
			ret += elt.Name
		case *ast.SelectorExpr:
			vtp, ok := elt.X.(*ast.Ident)
			if ok {
				ret += vtp.Name
			}
			ret += "." + elt.Sel.Name
		default:
			log.Panicf("%+#v", elt)
		}
	case *ast.Ident: // 基础类型或包内类型
		ret += vt.Name
	case *ast.SelectorExpr: // 选择符 形如：time.Time
		vtp, ok := vt.X.(*ast.Ident)
		if ok {
			ret += vtp.Name
		}
		ret += "." + vt.Sel.Name
	default: // 未知
		log.Panicf("<%#v>", vt)
	}
	return
}

func SortMap[V any](mp map[string]V, f func(k string, v V)) {
	keys := []string{}

	for k := range mp {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		f(keys[i], mp[keys[i]])
	}
}
