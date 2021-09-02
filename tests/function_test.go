package tests

import (
	"github.com/seyyedaghaei/ts2go/elements"
	"testing"
)


func TestFunction(t *testing.T) {
	program := readProgram("function.ts")
	if len(program.Statements) != 1 {
		t.Fatal("program must have exactly one statement")
	}

	switch statement := program.Statements[0].(type) {
	case *elements.Function:
		if statement.Identifier != "test" {
			t.Fatal("wrong identifier name")
		}
		if statement.Signature == nil {
			t.Fatal("function must have signature")
		} else {
			if statement.Signature.Type != "any" {
				t.Fatal("wrong return type")
			}
			parameters := statement.Signature.Parameters
			if parameters == nil || len(parameters) != 4{
				t.Fatal("program must have exactly four parameters")
			}
			a := parameters[0]
			if a.Type != "boolean" || a.Identifier != "a" || !a.Required || a.Rest || a.Initializer != nil || a.Accessibility != elements.None {
				t.Fail()
			}
			b := parameters[1]
			if b.Type != "unknown" || b.Identifier != "b" || b.Required || b.Rest || b.Initializer != nil || b.Accessibility != elements.None {
				t.Fail()
			}
			c := parameters[2]
			if c.Type != "number" || c.Identifier != "c" || c.Required || c.Rest || c.Initializer == nil || c.Accessibility != elements.None {
				t.Fail()
			}
			s := parameters[3]
			if s.Type != "string[]" || s.Identifier != "s" || s.Required || !s.Rest || s.Initializer != nil || s.Accessibility != elements.None {
				t.Fail()
			}
		}
		if statement.Body != nil && statement.Body.Statements != nil && len(statement.Body.Statements) > 0 {
			t.Fatal("function must not have any statement in body")
		}
	default:
		t.Fatal("type of statement must be *elements.Function")
	}
}
