package model

import (
	"fmt"
	"go/ast"
	"invalid/internal/rule"
	"sort"
)

type Field struct {
	Pkg   string       // 结构体所在包
	sName string       // 结构体名
	fName string       // 字段名
	fType string       // 类型信息
	tag   string       // iv tag 信息
	Rule  []rule.Ruler // 规则集合

	Val  map[string]string   // 需要的被提前声明的全局数据
	Pkgs map[string]struct{} // 需要导入的包
}

func NewField(pkg, structName string, field *ast.Field) (f *Field) {
	f = &Field{
		Pkg:   pkg,
		sName: structName,
		fName: "",
		fType: "",
		tag:   "",
		Rule:  []rule.Ruler{},
		Val:   map[string]string{},
		Pkgs:  map[string]struct{}{},
	}

	if field.Tag == nil || taget(field.Tag.Value, "iv") == "" {
		return nil
	}

	f.fName = field.Names[0].Name
	f.fType = GetFiledType(field.Type)
	f.tag = taget(field.Tag.Value, "iv")

	f.Val, f.Pkgs, f.Rule = rule.NewRules(structName, f.fType, f.fName, f.tag)

	if len(f.Rule) == 0 {
		return nil
	}

	sort.Slice(f.Rule, func(i, j int) bool {
		return f.Rule[i].Prio() < f.Rule[j].Prio()
	})

	return
}

func (f *Field) Meths() string {
	checkStr := ""
	for i := range f.Rule {
		checkStr += f.Rule[i].Check()
	}

	return fmt.Sprintf(fieldFuncStr, f.tag, f.sName, f.fName, checkStr)
}

const fieldFuncStr = `
 // 规则：<%s>
func (i *%s)_%s_invalid_() (err error) {
	%s

	return
}
`
