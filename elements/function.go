package elements

type Function struct {
	Identifier string
	Signature  *CallSignature
	Body       *Block
}

type CallSignature struct {
	// TODO: TypeParameters
	Parameters []*Parameter
	Type       string // TODO: Change this
}

type Parameter struct {
	// TODO: Decorators
	Identifier    string
	Required      bool
	Rest          bool
	Initializer   *Initializer
	Accessibility Accessibility
	Type          string // TODO: Change this
}

type Initializer struct {
	Expression Expression
}

type Argument struct {
	Rest       bool
	Expression Expression
	Identifier string
}

type MemberDot struct {
	Expression Expression
	Identifier string
	// TODO: nestedTypeGeneric
}

type Generator struct {
	Identifier string
	Parameters []*FormalParameter
	Body       *Block
}

type Iterator struct {
	Expression Expression
	Parameters []*FormalParameter
	Body       *Block
}

type MethodSignature struct {
	Name          string
	CallSignature *CallSignature
}

type PropertySignature struct {
	ReadOnly   bool
	Name       string
	Type       string // TODO: Change this
	ReturnType string // TODO: Change this
}

type PrimaryType interface{}

type ArrayType struct {
	Type PrimaryType
}
