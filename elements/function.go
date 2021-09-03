package elements

type Function struct {
	Identifier string
	Signature *CallSignature
	Body *Block
}

type CallSignature struct {
	// TODO: TypeParameters
	Parameters []*Parameter
	Type string // TODO: Change this
}

type Parameter struct {
	// TODO: Decorators
	Identifier string
	Required bool
	Rest bool
	Initializer *Initializer
	Accessibility Accessibility
	Type string // TODO: Change this
}

type Initializer struct {
	Expression Expression
}

type Argument struct {
	Rest bool
	Expression Expression
	Identifier string
}