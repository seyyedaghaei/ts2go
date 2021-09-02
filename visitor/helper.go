package visitor

import "github.com/antlr/antlr4/runtime/Go/antlr"

func (v *TypeScriptVisitor) returnNonNil(trees ...antlr.ParseTree) interface{} {
	for _, tree := range trees {
		if tree != nil {
			return tree.Accept(v)
		}
	}
	return nil
}
