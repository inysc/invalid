package demo

import (
	"invalid/internal/generate"
	"time"
)

type PInt int64

func (p *PInt) Invalid() error {
	return nil
}

// Pill 这是一个demo
// 尝试以他为基础设计 Gen
//go:generate invalid -types=Pill
type Pill struct {
	ID            *int64  `iv:"!nil; !2; [1, 20); d3"`
	Socre         *PInt   `iv:"d12; v"`
	SocreSlice    []PInt  `iv:"v"`
	SocreSlicePtr *[]PInt `iv:"v"`
	IDInt         int
	IDFloat64     float64
	IDSlice       []float64  `iv:"{1, 3, 4, i.IDFloat64}; dmake([]float64, 0, 10)"`
	IDSlicePtr    *[]float32 `iv:"{1, 3, 4}"`
	Not           int        `iv:"!2; !i.IDInt"`
	NotPtr        *int32     `iv:"!2; !nil"`
	NotSlice      []rune     `iv:"!2"`
	NotSlicePtr   *[]int8    `iv:"!nil; !4"`
	Name          string     `iv:"!\"\"; r\"^A\""`
	Demo0         generate.Generator
	Demo1         []generate.Generator `iv:""`
	Expire        time.Time            `iv:""`
	Roles         []string             `iv:""`
	Errs          []error
	Demo2         struct{ A string }
	// Errors [5]error // 占时没兴趣做数组，先做做切片 // TODO:
}

//go:generate invalid -types=NewReq
type NewReq struct {
	ID    int64
	Score *float64 `iv:"!nil; !20; [1,)"`
	Name  string   `iv:""`
	Phone string
	Role  struct{ Type string }
}
