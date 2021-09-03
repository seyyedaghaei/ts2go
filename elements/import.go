package elements

type FromBlock struct {
	Members []string
	Alias   string
	Package string
}

type Export struct {
	Default bool
	FromBlock *FromBlock
	Statement Statement
}

type ImportAlias struct {
	Name string
	NameSpace []string
}

type Import struct {
	FromBlock *FromBlock
	ImportAlias *ImportAlias
}