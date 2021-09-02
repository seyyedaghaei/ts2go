package tests

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/seyyedaghaei/ts2go/ast"
	"testing"
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

func TestParser(t *testing.T) {
	parser := readParser("empty.ts")
	if parser == nil {
		t.Fail()
	}
}
