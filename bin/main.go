package main

import (
	"invalid/internal/generate"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	g := generate.Generator{}

	g.ParsePackage([]string{"."}, []string{})

	// 解析指定类型
	for _, name := range types {
		g.Parse(name)
	}

	g.Generate()

	g.Save(types[0])
}
