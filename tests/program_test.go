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
	program := readParser("empty.ts").Program()
	if program == nil || program.Accept(&visitor.TypeScriptVisitor{}) == nil {
		t.Fail()
	}
}