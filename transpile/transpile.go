package transpile

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/seyyedaghaei/ts2go/ast"
	"github.com/seyyedaghaei/ts2go/elements"
	"github.com/seyyedaghaei/ts2go/visitor"
	"log"
)

func readParser(data string) *ast.TypeScriptParser {
	input := antlr.NewInputStream(data)
	lexer := ast.NewTypeScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := ast.NewTypeScriptParser(stream)
	p.BuildParseTrees = true
	return p
}

func Transpile(text string) string {
	programCtx := readParser(text).Program()
	if programCtx == nil {
		log.Fatal("ProgramContext is nil")
	}
	program := programCtx.Accept(&visitor.TypeScriptVisitor{}).(*elements.Program)
	if program == nil {
		log.Fatal("Program is nil")
	}
	return fmt.Sprintf("package main\n%s", (&ElementsVisitor{}).VisitProgram(program))
}
