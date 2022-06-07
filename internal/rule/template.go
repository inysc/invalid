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
	//go:embed template/not_slice.tmpl
	notSliceFuncStr string
	//go:embed template/not_slice_ptr.tmpl
	notSlicePtrFuncStr string
	notTmpl            *template.Template
	notPtrTmpl         *template.Template
	notSliceTmpl       *template.Template
	notSlicePtrTmpl    *template.Template

	//go:embed template/enum.tmpl
	enumFuncStr string
	//go:embed template/enum_ptr.tmpl
	enumPtrFuncStr string
	//go:embed template/enum_slice.tmpl
	enumSliceFuncStr string
	//go:embed template/enum_slice_ptr.tmpl
	enumSlicePtrFuncStr string
	enumTmpl            *template.Template
	enumPtrTmpl         *template.Template
	enumSliceTmpl       *template.Template
	enumSlicePtrTmpl    *template.Template

	//go:embed template/duck.tmpl
	duckFuncStr string
	//go:embed template/duck_slice.tmpl
	duckSliceFuncStr string
	//go:embed template/duck_slice_ptr.tmpl
	duckSlicePtrFuncStr string
	duckTmpl            *template.Template
	duckSliceTmpl       *template.Template
	duckSlicePtrTmpl    *template.Template

	//go:embed template/regex.tmpl
	regexFuncStr string
	//go:embed template/regex_ptr.tmpl
	regexPtrFuncStr string
	//go:embed template/regex_slice.tmpl
	regexSliceFuncStr string
	//go:embed template/regex_slice_ptr.tmpl
	regexSlicePtrFuncStr string
	regexTmpl            *template.Template
	regexPtrTmpl         *template.Template
	regexSliceTmpl       *template.Template
	regexSlicePtrTmpl    *template.Template
)

func init() {

	funcMap := template.FuncMap{
		"remove_ptr":       remove_ptr,
		"remove_slice":     remove_slice,
		"remove_slice_ptr": remove_slice_ptr,
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

		notSliceTmpl, err = template.
			New("invalid_not_slice").
			Funcs(funcMap).
			Parse(notSliceFuncStr)
		if err != nil {
			panic(err)
		}

		notSlicePtrTmpl, err = template.
			New("invalid_not_slice_ptr").
			Funcs(funcMap).
			Parse(notSlicePtrFuncStr)
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

		enumSliceTmpl, err = template.
			New("invalid_enum_slice").
			Funcs(funcMap).
			Parse(enumSliceFuncStr)
		if err != nil {
			panic(err)
		}

		enumSlicePtrTmpl, err = template.
			New("invalid_enum_slice_ptr").
			Funcs(funcMap).
			Parse(enumSlicePtrFuncStr)
		if err != nil {
			panic(err)
		}
	}

	// duck
	{
		duckTmpl, err = template.
			New("invalid_duck").
			Funcs(funcMap).
			Parse(duckFuncStr)
		if err != nil {
			panic(err)
		}

		duckSliceTmpl, err = template.
			New("invalid_duck_slice").
			Funcs(funcMap).
			Parse(duckSliceFuncStr)
		if err != nil {
			panic(err)
		}

		duckSlicePtrTmpl, err = template.
			New("invalid_duck_slice_ptr").
			Funcs(funcMap).
			Parse(duckSlicePtrFuncStr)
		if err != nil {
			panic(err)
		}
	}

	// regex
	{
		regexTmpl, err = template.
			New("invalid_regex").
			Funcs(funcMap).
			Parse(regexFuncStr)
		if err != nil {
			panic(err)
		}

		regexPtrTmpl, err = template.
			New("invalid_regex_ptr").
			Funcs(funcMap).
			Parse(regexPtrFuncStr)
		if err != nil {
			panic(err)
		}

		regexSliceTmpl, err = template.
			New("invalid_regex_slice").
			Funcs(funcMap).
			Parse(regexSliceFuncStr)
		if err != nil {
			panic(err)
		}

		regexSlicePtrTmpl, err = template.
			New("invalid_regex_slice_ptr").
			Funcs(funcMap).
			Parse(regexSlicePtrFuncStr)
		if err != nil {
			panic(err)
		}
	}
}
