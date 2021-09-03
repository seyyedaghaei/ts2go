package elements

type TypeParameter struct {

}

type TypeAlias struct {
	Identifier string
	Parameters []*TypeParameter
	Type string // TODO: Change this
}

type TypeMember interface {}

type FunctionType struct {
	TypeParameters []*TypeParameter
	Parameters []*Parameter
	Type string // TODO: Change this
}

type ConstructorType struct {
	TypeParameters []*TypeParameter
	Parameters []*Parameter
	Type string // TODO: Change this
}