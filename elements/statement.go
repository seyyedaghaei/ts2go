package elements

type Statement interface {
}

type Empty struct {
}

type Continue struct {
}

type Break struct {
}

type Return struct {
	Expressions []Expression
}

type Yield struct {
	Expressions []Expression
}

type With struct {
	Expressions []Expression
	Statement   Statement
}

type LabeledStatement struct {
	Identifier string
	Statement Statement
}

type Try struct {
	Block *Block
	Catch *Catch
	Finally *Finally
}

type Catch struct {
	Identifier string
	Block *Block
}

type Finally struct {
	Block *Block
}