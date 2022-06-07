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
