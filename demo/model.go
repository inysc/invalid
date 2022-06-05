package demo

import (
	"invalid/internal/generate"
	"time"
)

type PInt int64

// Pill 这是一个demo
// 尝试以他为基础设计 Gen
//go:generate invalid -type=Pill
type Pill struct {
	ID     *int64 `iv:"!nil; [1, 20); @3"`
	Socre  *PInt  `iv:"@12; :>"`
	Name   string `iv:"!\"\""`
	Demo0  generate.Generator
	Demo1  []generate.Generator `iv:""`
	Expire time.Time            `iv:""`
	Roles  []string             `iv:""`
	Errs   []error
	Demo2  struct{ A string }
	// Errors [5]error // 占时没兴趣做数组，先做做切片 // TODO:
}
