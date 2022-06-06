package rule

import (
	_ "embed"
	"text/template"
)

var (
	err error

	//go:embed template/default.tmpl
	defaultFuncStr string
	//go:embed template/default_ptr.tmpl
	defaultPtrFuncStr string
	defaultTmpl       *template.Template
	defaultPtrTmpl    *template.Template

	//go:embed template/not.tmpl
	notFuncStr string
	//go:embed template/not_ptr.tmpl
	notPtrFuncStr string
	notTmpl       *template.Template
	notPtrTmpl    *template.Template

	//go:embed template/enum.tmpl
	enumFuncStr string
	//go:embed template/enum_ptr.tmpl
	enumPtrFuncStr string
	enumTmpl       *template.Template
	enumPtrTmpl    *template.Template
)

func init() {

	funcMap := template.FuncMap{
		"remove_ptr": remove_ptr,
	}

	// default
	{
		defaultTmpl, err = template.
			New("invalid_default").
			Funcs(funcMap).
			Parse(defaultFuncStr)
		if err != nil {
			panic(err)
		}

		defaultPtrTmpl, err = template.
			New("invalid_default_ptr").
			Funcs(funcMap).
			Parse(defaultPtrFuncStr)
		if err != nil {
			panic(err)
		}
	}

	// not
	{
		notTmpl, err = template.
			New("invalid_not").
			Funcs(funcMap).
			Parse(notFuncStr)
		if err != nil {
			panic(err)
		}

		notPtrTmpl, err = template.
			New("invalid_not_ptr").
			Funcs(funcMap).
			Parse(notPtrFuncStr)
		if err != nil {
			panic(err)
		}
	}

	// enum
	{
		enumTmpl, err = template.
			New("invalid_enum").
			Funcs(funcMap).
			Parse(enumFuncStr)
		if err != nil {
			panic(err)
		}

		enumPtrTmpl, err = template.
			New("invalid_enum_ptr").
			Funcs(funcMap).
			Parse(enumPtrFuncStr)
		if err != nil {
			panic(err)
		}
	}
}
