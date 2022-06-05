package model

import (
	"fmt"
	"go/ast"
	"invalid/internal/rule"
	"strings"
)

type Field struct {
	Pkg   string       // 结构体所在包
	sName string       // 结构体名
	fName string       // 字段名
	fType string       // 类型信息
	tag   string       // iv tag 信息
	duck  bool         // 是否实现了 interface { Invalid() error }
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

	f.Val, f.Pkgs, f.duck, f.Rule = rule.NewRules(structName, f.fType, f.fName, f.tag)

	if len(f.Rule) == 0 {
		return nil
	}

	return
}

func (f *Field) Meths() string {

	callStr := ""
	methStr := ""
	deferStr := ""
	for _, v := range f.Rule {
		methStr += v.Meth()
		name := v.Name()
		if !strings.Contains(name, "_invalid_default_") {
			callStr += fmt.Sprintf(callMethStr, v.Name())
		} else {
			deferStr = fmt.Sprintf(deferMethStr, f.fName)
		}
	}

	if f.duck {
		callStr += fmt.Sprintf(callMethStr, f.fName+".Invalid")
	}

	return fmt.Sprintf(fieldFuncStr, f.sName, f.fName, deferStr, callStr, methStr)
}

const fieldFuncStr = `
func (i *%s)_%s_invalid_() (err error) {
	%s

	%s

	return
}

%s

`

const callMethStr = `
err = i.%s()
if err != nil {
	return err
}

`

const deferMethStr = `
defer func ()  {
	if err != nil {
		err = i._%s_default_()
	}
}()

`
