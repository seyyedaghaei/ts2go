package tests

import (
	"github.com/seyyedaghaei/ts2go/elements"
	"github.com/seyyedaghaei/ts2go/visitor"
	"testing"
)

func readProgram(file string) *elements.Program {
	return readParser(file).Program().Accept(&visitor.TypeScriptVisitor{}).(*elements.Program)
}

func TestProgram(t *testing.T) {
	programCtx := readParser("empty.ts").Program()
	if programCtx == nil {
		t.Fatal("ProgramContext is nil")
	}
	program := programCtx.Accept(&visitor.TypeScriptVisitor{}).(*elements.Program)
	if program == nil {
		t.Fatal("Program is nil")
	}
}