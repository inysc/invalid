package rule

import "strings"

func remove_ptr(typ string) string {
	return strings.TrimPrefix(typ, "*")
}

func remove_slice(typ string) string {
	return strings.TrimPrefix(typ, "[]")
}

func remove_slice_ptr(typ string) string {
	return strings.TrimPrefix(typ, "*[]")
}

// 取反关系运算符
func rRo(l string) string {
	switch l {
	case ">":
		return "<="
	case "<":
		return ">="
	case ">=":
		return "<"
	case "<=":
		return ">"
	}
	return l
}

// 翻译关系运算符
func tRo(l string) string {
	switch l {
	case ">":
		return "greater than"
	case "<":
		return "less than"
	case ">=":
		return "greater than or equal to"
	case "<=":
		return "less than or equal to"
	}
	return l
}
