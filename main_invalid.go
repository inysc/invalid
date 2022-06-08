package main

import (
	"errors"
	"invalid/demo"
	"regexp"
	"time"
)

func main() {
	p := new(demo.Pill)
	p.Invalid()
}

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
	Expire        time.Time  `iv:""`
	Roles         []string   `iv:""`
	Errs          []error
	Demo2         struct{ A string }
	// Errors [5]error // 占时没兴趣做数组，先做做切片 // TODO:
}

var (
	regex_19 *regexp.Regexp
)

func init() {
	var err error

	regex_19, err = regexp.Compile("^A")
	if err != nil {
		panic(err)
	}

}

func (i *Pill) Invalid() (err error) {

	for _, v := range []func() error{i._ID_invalid_} {
		err = v()
		if err != nil {
			return err
		}
	}

	err = i._ID_invalid_()
	if err != nil {
		return err
	}

	err = i._Socre_invalid_()
	if err != nil {
		return err
	}

	err = i._SocreSlice_invalid_()
	if err != nil {
		return err
	}

	err = i._SocreSlicePtr_invalid_()
	if err != nil {
		return err
	}

	err = i._IDSlice_invalid_()
	if err != nil {
		return err
	}

	err = i._IDSlicePtr_invalid_()
	if err != nil {
		return err
	}

	err = i._Not_invalid_()
	if err != nil {
		return err
	}

	err = i._NotPtr_invalid_()
	if err != nil {
		return err
	}

	err = i._NotSlice_invalid_()
	if err != nil {
		return err
	}

	err = i._NotSlicePtr_invalid_()
	if err != nil {
		return err
	}

	err = i._Name_invalid_()
	if err != nil {
		return err
	}

	return
}

func (i *Pill) _ID_invalid_() (err error) {

	defer func() {
		if err != nil {
			err = i._ID_default_()
		}
	}()

	err = i._ID_invalid_not_1_()
	if err != nil {
		return err
	}

	err = i._ID_invalid_not_2_()
	if err != nil {
		return err
	}

	err = i._ID_invalid_range_3_()
	if err != nil {
		return err
	}

	return
}

// !nil
func (i *Pill) _ID_invalid_not_1_() error {
	if i.ID == nil {
		return errors.New(`invalid<not>: Pill.ID must not be nil`)
	}
	return nil
}

// !2
func (i *Pill) _ID_invalid_not_2_() error {
	if i.ID != nil && *i.ID == 2 {
		return errors.New(`invalid<not>: Pill.ID must not be 2`)
	}
	return nil
}
func (i *Pill) _ID_invalid_range_3_() error {
	if i.ID == nil {
		return nil
	}
	if *i.ID < 1 {
		return errors.New("invalid<range>: Pill.ID must not be less than 1")
	}

	if *i.ID >= 20 {
		return errors.New("invalid<range>: Pill.ID must not be greater than or equal to  20")
	}

	return nil
}

// d3
func (i *Pill) _ID_default_() error {
	var tmp int64 = 3
	i.ID = &tmp
	return nil
}

func (i *Pill) _Socre_invalid_() (err error) {

	defer func() {
		if err != nil {
			err = i._Socre_default_()
		}
	}()

	err = i._Socre_invalid_duck_4_()
	if err != nil {
		return err
	}

	return
}

// d12
func (i *Pill) _Socre_default_() error {
	var tmp PInt = 12
	i.Socre = &tmp
	return nil
}

// v
func (i *Pill) _Socre_invalid_duck_4_() error {
	return i.Socre.Invalid()
}

func (i *Pill) _SocreSlice_invalid_() (err error) {

	err = i._SocreSlice_invalid_duck_5_()
	if err != nil {
		return err
	}

	return
}

// v
func (i *Pill) _SocreSlice_invalid_duck_5_() (err error) {
	for idx := range i.SocreSlice {
		err = i.SocreSlice[idx].Invalid()
		if err != nil {
			return
		}
	}
	return
}

func (i *Pill) _SocreSlicePtr_invalid_() (err error) {

	err = i._SocreSlicePtr_invalid_duck_6_()
	if err != nil {
		return err
	}

	return
}

// v
func (i *Pill) _SocreSlicePtr_invalid_duck_6_() (err error) {
	for idx := range *(i.SocreSlicePtr) {
		err = (*(i.SocreSlicePtr))[idx].Invalid()
		if err != nil {
			return
		}
	}
	return
}

func (i *Pill) _IDSlice_invalid_() (err error) {

	defer func() {
		if err != nil {
			err = i._IDSlice_default_()
		}
	}()

	err = i._IDSlice_invalid_enum_8_()
	if err != nil {
		return err
	}

	return
}

// {1, 3, 4, i.IDFloat64}
func (i *Pill) _IDSlice_invalid_enum_8_() error {
	tot := 0 // 符合枚举值的总数
	for idx := range i.IDSlice {
		for _, v := range []float64{1, 3, 4, i.IDFloat64} {
			if i.IDSlice[idx] == v {
				tot++
				break
			}
		}
	}
	if tot == len(i.IDSlice) {
		return nil
	}

	return errors.New(`invalid<enum>: each of Pill.IDSlice must be in the {1, 3, 4, i.IDFloat64}`)
}

// dmake([]float64, 0, 10)
func (i *Pill) _IDSlice_default_() error {
	i.IDSlice = make([]float64, 0, 10)
	return nil
}

func (i *Pill) _IDSlicePtr_invalid_() (err error) {

	err = i._IDSlicePtr_invalid_enum_10_()
	if err != nil {
		return err
	}

	return
}

// {1, 3, 4}
func (i *Pill) _IDSlicePtr_invalid_enum_10_() error {
	if i.IDSlicePtr == nil {
		return nil
	}

	tot := 0 // 符合枚举值的总数
	for _, val := range *(i.IDSlicePtr) {
		for _, v := range []float32{1, 3, 4} {
			if val == v {
				tot++
				break
			}
		}
	}
	if tot == len(*(i.IDSlicePtr)) {
		return nil
	}

	return errors.New(`invalid<enum>: each of Pill.IDSlicePtr must be in the {1, 3, 4}`)
}

func (i *Pill) _Not_invalid_() (err error) {

	err = i._Not_invalid_not_11_()
	if err != nil {
		return err
	}

	err = i._Not_invalid_not_12_()
	if err != nil {
		return err
	}

	return
}

// !2
func (i *Pill) _Not_invalid_not_11_() error {
	if i.Not == 2 {
		return errors.New(`invalid<not>: Pill.Not must not be 2`)
	}
	return nil
}

// !i.IDInt
func (i *Pill) _Not_invalid_not_12_() error {
	if i.Not == i.IDInt {
		return errors.New(`invalid<not>: Pill.Not must not be i.IDInt`)
	}
	return nil
}

func (i *Pill) _NotPtr_invalid_() (err error) {

	err = i._NotPtr_invalid_not_13_()
	if err != nil {
		return err
	}

	err = i._NotPtr_invalid_not_14_()
	if err != nil {
		return err
	}

	return
}

// !2
func (i *Pill) _NotPtr_invalid_not_13_() error {
	if i.NotPtr != nil && *i.NotPtr == 2 {
		return errors.New(`invalid<not>: Pill.NotPtr must not be 2`)
	}
	return nil
}

// !nil
func (i *Pill) _NotPtr_invalid_not_14_() error {
	if i.NotPtr == nil {
		return errors.New(`invalid<not>: Pill.NotPtr must not be nil`)
	}
	return nil
}

func (i *Pill) _NotSlice_invalid_() (err error) {

	err = i._NotSlice_invalid_not_15_()
	if err != nil {
		return err
	}

	return
}

// !2
func (i *Pill) _NotSlice_invalid_not_15_() error {
	for idx := range i.NotSlice {
		if i.NotSlice[idx] == 2 {
			return errors.New(`invalid<not>: each of Pill.NotSlice must not be 2`)
		}
	}
	return nil
}

func (i *Pill) _NotSlicePtr_invalid_() (err error) {

	err = i._NotSlicePtr_invalid_not_16_()
	if err != nil {
		return err
	}

	err = i._NotSlicePtr_invalid_not_17_()
	if err != nil {
		return err
	}

	return
}

// !nil
func (i *Pill) _NotSlicePtr_invalid_not_16_() error {
	if i.NotSlicePtr == nil {
		return errors.New(`invalid<not>: Pill.NotSlicePtr must not be nil`)
	}
	return nil
}

// !4
func (i *Pill) _NotSlicePtr_invalid_not_17_() error {
	if i.NotSlicePtr == nil {
		return nil
	}

	for _, val := range *(i.NotSlicePtr) {
		if val == 4 {
			return errors.New(`invalid<not>: each of Pill.NotSlicePtr must not be 4`)
		}
	}
	return nil
}

func (i *Pill) _Name_invalid_() (err error) {

	err = i._Name_invalid_not_18_()
	if err != nil {
		return err
	}

	err = i._Name_invalid_regex_19_()
	if err != nil {
		return err
	}

	return
}

// !""
func (i *Pill) _Name_invalid_not_18_() error {
	if i.Name == "" {
		return errors.New(`invalid<not>: Pill.Name must not be ""`)
	}
	return nil
}

// r"^A"
func (i *Pill) _Name_invalid_regex_19_() error {
	if !regex_19.MatchString(i.Name) {
		return errors.New(`invalid<regex>: Pill.Name must conform to the regex<"^A">`)
	}
	return nil
}
