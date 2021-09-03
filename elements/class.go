package elements

type Class struct {
	Abstract   bool
	Identifier string
	// TODO: TypeParameters
	// TODO: ClassHeritage
	Elements []ClassElement
}

type ClassElement interface{}

type Constructor struct {
	Accessibility Accessibility
	Parameters    []*FormalParameter
	Body          *Block
}

type FormalParameter struct {
	// TODO: Decorator
	Accessibility Accessibility
	Name          string // TODO: Maybe you should change this
	// TODO: TypeAnnotation
	Type       string // TODO: You must change this
	Expression Expression
	Rest       bool
}

type Getter struct {
	Name string
	Type string // TODO: Change this
	Body *Block
}

type Setter struct {
	Name     string
	Type     string // TODO: Change this
	Body     *Block
	Argument string // TODO: Absolutely change this
}

type IndexSignature struct {
	Identifier string
	IsNumber   bool
	Type       string // TODO: Change this
}
