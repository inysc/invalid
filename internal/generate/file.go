package generate

import (
	"go/ast"
	"invalid/internal/model"
)

type File struct {
	Pkg  *Package  // Package to which this file belongs.
	File *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	TypeName string          // Name of the constant type.
	Values   []*model.Struct // Accumulator for constant values of that type.

	TrimPrefix  string
	LineComment bool
}
