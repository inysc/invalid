package rule

import "strings"

func remove_ptr(typ string) string {
	return strings.ReplaceAll(typ, "*", "")
}
