package model

import (
	"fmt"
	"go/ast"
)

type Struct struct {
	Pkg    string   // 结构体所在的包
	Name   string   // 结构体名
	Fields []*Field // 字段信息

	// key<变量声明>	value<声明及其处理 error>
	Val  map[string]string   // 需要的被提前声明的全局数据
	Pkgs map[string]struct{} // 需要导入的包
}

func (s *Struct) AddField(field *Field) {
	for k, v := range field.Val {
		s.Val[k] = v
	}
	for k, v := range field.Pkgs {
		s.Pkgs[k] = v
	}
	s.Fields = append(s.Fields, field)
}

func NewStruct(pkg, name string, typ *ast.StructType) (st *Struct) {
	st = &Struct{
		Pkg:    pkg,
		Name:   name,
		Fields: []*Field{},
		Val:    map[string]string{},
		Pkgs:   map[string]struct{}{},
	}
	for _, v := range typ.Fields.List {
		f := NewField(pkg, name, v)
		if f != nil {
			st.AddField(f)
		}
	}

	return
}

func (s *Struct) Meths() string {

	callStr := ""
	methStr := ""

	for _, v := range s.Fields {
		name := fmt.Sprintf("_%s_invalid_", v.fName)
		callStr += fmt.Sprintf(callMethStr, name)
		methStr += v.Meths()
	}

	return fmt.Sprintf(structMethsStr, s.Name, callStr, methStr)
}

const structMethsStr = `
func (i *%s)Invalid() (err error) {

	%s

	return
}

%s
`
