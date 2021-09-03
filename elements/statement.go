package elements

type Statement interface{}

type Empty struct{}

type Continue struct{
	Identifier string
}

type Break struct {
	Identifier string
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
	Statement  Statement
}

type Try struct {
	Block   *Block
	Catch   *Catch
	Finally *Finally
}

type Catch struct {
	Identifier string
	Block      *Block
}

type Finally struct {
	Block *Block
}

type Throw struct {
	Expressions []Expression
}

type NameSpace struct {
	Names      []string
	Statements []Statement
}

type Enum struct {
	Const      bool
	Identifier string
	Members    []*EnumMember
}

type EnumMember struct {
	Name       string
	Expression Expression
}

type Switch struct {
	Expressions []Expression
	CaseClauses []*CaseClause
}

type CaseClause struct {
	Expressions []Expression
	Statements  []Statement
}

type Debugger struct {} // Maybe useless
