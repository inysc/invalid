package main

import (
	"flag"
	"strings"
)

var (
	types_input = flag.String("types", "Pill", "待生成校验方法的函数")
)

var (
	types []string
)

func init() {
	flag.Parse()

	types = strings.Split(*types_input, ",")
}
