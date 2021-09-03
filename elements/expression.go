package elements

type Expression interface {
}

type Delete struct {
	Expression Expression
}

type Void struct {
	Expression Expression
}

type TypeOf struct {
	Expression Expression
}

type PreIncrement struct {
	Expression Expression
}

type PreDecrease struct {
	Expression Expression
}

type PostIncrement struct {
	Expression Expression
}

type PostDecrease struct {
	Expression Expression
}

type UnaryPlus struct {
	Expression Expression
}

type UnaryMinus struct {
	Expression Expression
}

type BitNot struct {
	Expression Expression
}

type Not struct {
	Expression Expression
}

type Sign string // TODO: Change this

type Multiplicative struct {
	LeftExpression  Expression
	RightExpression Expression
	Sign            Sign
}

type Additive struct {
	LeftExpression  Expression
	RightExpression Expression
	Sign            Sign
}

type BitShift struct {
	LeftExpression  Expression
	RightExpression Expression
	Sign            Sign
}

type Relational struct {
	LeftExpression  Expression
	RightExpression Expression
	Sign            Sign
}

type AssignmentOperator struct {
	LeftExpression  Expression
	RightExpression Expression
	Sign            Sign
}

type In struct {
	LeftExpression  Expression
	RightExpression Expression
}

type InstanceOf struct {
	LeftExpression  Expression
	RightExpression Expression
}

type BitAnd struct {
	LeftExpression  Expression
	RightExpression Expression
}

type BitXOr struct {
	LeftExpression  Expression
	RightExpression Expression
}

type BitOr struct {
	LeftExpression  Expression
	RightExpression Expression
}

type LogicalAnd struct {
	LeftExpression  Expression
	RightExpression Expression
}

type LogicalOr struct {
	LeftExpression  Expression
	RightExpression Expression
}

type Assignment struct {
	LeftExpression  Expression
	RightExpression Expression
	Sign Sign
}

type Ternary struct {
	Question        Expression
	TrueExpression  Expression
	FalseExpression Expression
}

type This struct{}

type Super struct{}

type EOS struct {}

type Null struct {}

type Arguments struct {
	Expression Expression
	Arguments []*Argument
}

type MemberIndex struct {
	Expression Expression
	Index []Expression
}