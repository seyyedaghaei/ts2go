package elements

type Variable struct {
	// TODO: array and object literal
	Identifier string
	Type string // TODO: Change this
	Expression Expression
	TypeParameters []*TypeParameter
	Initializer Expression
}
