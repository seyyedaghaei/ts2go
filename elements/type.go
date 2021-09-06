package elements

type TypeParameter struct {
	Identifier     string
	Constraint     string // TODO: Maybe you should change this
	TypeParameters []*TypeParameter
}

type TypeAlias struct {
	Identifier string
	Parameters []*TypeParameter
	Type       string // TODO: Change this
}

type TypeMember interface{}

type FunctionType struct {
	TypeParameters []*TypeParameter
	Parameters     []*Parameter
	Type           string // TODO: Change this
}

type ConstructorType struct {
	TypeParameters []*TypeParameter
	Parameters     []*Parameter
	Type           string // TODO: Change this
}
