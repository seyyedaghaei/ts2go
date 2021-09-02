package tests

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/seyyedaghaei/ts2go/ast"
)

func readParser(file string) *ast.TypeScriptParser {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		panic(err)
	}
	lexer := ast.NewTypeScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := ast.NewTypeScriptParser(stream)
	p.BuildParseTrees = true
	return p
}
