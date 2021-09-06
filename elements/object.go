package elements

type Object struct {
	Properties []ObjectProperty
}

type ObjectProperty interface{}

type PropertyAssignment struct {
	Name string
	Expression Expression
}

type ComputedPropertyAssignment struct {
	Name Expression
	Expression Expression
}

type BindingPattern struct {
	Array *Array
	Object *Object
}