package rule

import (
	"strings"
)

func splitTag(tag string) (ts []string) {

	if len(tag) == 0 {
		return []string{}
	}

	idxs := []int{-1}

	isTrans := false
	for i, v := range tag {
		if v == ';' && !isTrans {
			idxs = append(idxs, i)
		} else if v == '\\' {
			isTrans = true
		} else {
			isTrans = false
		}
	}
	idxs = append(idxs, len(tag))

	l := len(idxs)
	for i := 1; i < l; i++ {
		ts = append(ts, strings.TrimSpace(tag[idxs[i-1]+1:idxs[i]]))
	}

	return
}
