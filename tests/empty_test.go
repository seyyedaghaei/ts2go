package tests

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	parser := readParser("empty.ts")
	if parser == nil {
		t.Fail()
	} else {
		tree := parser.Program()
		if tree == nil {
			t.Fail()
		}
	}
}
