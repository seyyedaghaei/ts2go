package elements

type Array struct {
	Elements []*ArrayElement
}

type ArrayElement struct {
	Ellipsis bool
	Expression Expression
	Identifier string
}
