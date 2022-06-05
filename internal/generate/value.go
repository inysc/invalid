package generate

// Struct represents a declared constant.
type Struct struct {
	OriginalName string // The name of the constant.
	Name         string // The name with trimmed prefix.
	Value        uint64 // Will be converted to int64 when needed.
	Signed       bool   // Whether the constant is a signed type.
	Str          string // The string representation given by the "go/constant" package.
}
