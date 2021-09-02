package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// TypeScriptLexerBase state
type TypeScriptLexerBase struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	stackLength      int
	stackIx          int

	lastToken        antlr.Token
	useStrictDefault bool
	useStrictCurrent bool
	/**
	 * Keeps track of the the current depth of nested template string backticks.
	 * E.g. after the X in:
	 *
	 * `${a ? `${X
	 *
	 * templateDepth will be 2. This variable is needed to determine if a `}` is a
	 * plain CloseBrace, or one that closes an expression inside a template string.
	 */
	templateDepth int
}

func (l *TypeScriptLexerBase) IsInTemplateString() bool {
	return l.templateDepth > 0
}

func (l *TypeScriptLexerBase) IncreaseTemplateDepth() {
	l.templateDepth += 1
}

func (l *TypeScriptLexerBase) DecreaseTemplateDepth() {
	l.templateDepth -= 1
}

func (l *TypeScriptLexerBase) IsStartOfFile() bool {
	return l.lastToken == nil
}

func (l *TypeScriptLexerBase) pushStrictModeScope(v bool) {
	if l.stackIx == l.stackLength {
		l.scopeStrictModes = append(l.scopeStrictModes, v)
		l.stackLength++
	} else {
		l.scopeStrictModes[l.stackIx] = v
	}
	l.stackIx++
}

func (l *TypeScriptLexerBase) popStrictModeScope() bool {
	l.stackIx--
	v := l.scopeStrictModes[l.stackIx]
	l.scopeStrictModes[l.stackIx] = false
	return v
}

// IsStrictMode is self explanatory.
func (l *TypeScriptLexerBase) IsStrictMode() bool {
	return l.useStrictCurrent
}

// NextToken from the character stream.
func (l *TypeScriptLexerBase) NextToken() antlr.Token {
	next := l.BaseLexer.NextToken() // Get next token
	if next.GetChannel() == antlr.TokenDefaultChannel {
		// Keep track of the last token on default channel
		l.lastToken = next
	}
	return next
}

// ProcessOpenBrace is called when a { is encountered during
// lexing, we push a new scope everytime.
func (l *TypeScriptLexerBase) ProcessOpenBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 && l.scopeStrictModes[l.stackIx-1] {
		l.useStrictCurrent = true
	}
	l.pushStrictModeScope(l.useStrictCurrent)
}

// ProcessCloseBrace is called when a } is encountered during
// lexing, we pop a scope unless we're inside global scope.
func (l *TypeScriptLexerBase) ProcessCloseBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 {
		l.useStrictCurrent = l.popStrictModeScope()
	}
}

// ProcessStringLiteral is called when lexing a string literal.
func (l *TypeScriptLexerBase) ProcessStringLiteral() {
	if l.lastToken == nil || l.lastToken.GetTokenType() == TypeScriptLexerOpenBrace {
		if l.GetText() == `"use strict"` || l.GetText() == "'use strict'" {
			if l.stackIx > 0 {
				l.popStrictModeScope()
			}
			l.useStrictCurrent = true
			l.pushStrictModeScope(l.useStrictCurrent)
		}
	}
}

// IsRegexPossible returns true if the lexer can match a
// regex literal.
func (l *TypeScriptLexerBase) IsRegexPossible() bool {
	if l.lastToken == nil {
		return true
	}
	switch l.lastToken.GetTokenType() {
	case TypeScriptLexerIdentifier, TypeScriptLexerNullLiteral,
		TypeScriptLexerBooleanLiteral, TypeScriptLexerThis,
		TypeScriptLexerCloseBracket, TypeScriptLexerCloseParen,
		TypeScriptLexerOctalIntegerLiteral, TypeScriptLexerDecimalLiteral,
		TypeScriptLexerHexIntegerLiteral, TypeScriptLexerStringLiteral,
		TypeScriptLexerPlusPlus, TypeScriptLexerMinusMinus:
		return false
	default:
		return true
	}
}
