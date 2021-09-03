package elements

type TypeParameter struct {

}

type TypeAlias struct {
	Identifier string
	Parameters []*TypeParameter
	Type string // TODO: Change this
}
