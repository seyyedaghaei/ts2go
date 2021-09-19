package transpile

import (
	"encoding/json"
	"fmt"
	"github.com/seyyedaghaei/ts2go/elements"
	"reflect"
	"strings"
)

type ElementsVisitor struct{}

func (v *ElementsVisitor) VisitImport(imp *elements.Import) string {
	if imp.FromBlock != nil {
		return fmt.Sprintf("import %s %s", imp.FromBlock.Alias, imp.FromBlock.Package)
	}
	return fmt.Sprintf("import %s %s")
}

func (v *ElementsVisitor) VisitMemberDot(memberDot *elements.MemberDot) string {
	return fmt.Sprintf("%s.%s", v.VisitExpression(memberDot.Expression), memberDot.Identifier)
}

func (v *ElementsVisitor) VisitParenthesized(parenthesized *elements.Parenthesized) string {
	return "(" + v.VisitExpression(parenthesized.Expressions) + ")"
}

func (v *ElementsVisitor) VisitArguments(arguments *elements.Arguments) string {
	args := make([]string, len(arguments.Arguments))
	for i, argument := range arguments.Arguments {
		args[i] = v.VisitArgument(argument)
	}
	return fmt.Sprintf("%s(%s)", v.VisitExpression(arguments.Expression), strings.Join(args, ", "))
}

func (v *ElementsVisitor) VisitArgument(argument *elements.Argument) string {
	res := ""
	if argument.Expression != nil {
		res = v.VisitExpression(argument.Expression)
	} else {
		res = argument.Identifier
	}
	if argument.Rest {
		res += "..."
	}
	return res
}

func (v *ElementsVisitor) VisitStatements(statements []elements.Statement) string {
	p := make([]string, len(statements))
	for i, s := range statements {
		p[i] = v.VisitStatement(s)
	}
	return strings.Join(p, "\n")
}

func (v *ElementsVisitor) VisitStatement(statement elements.Statement) string {
	switch s := statement.(type) {
	case *elements.Function:
		return v.VisitFunction(s)
	case []*elements.Variable:
		return v.VisitVariables(s)
	case []elements.Expression:
		return v.VisitExpressions(s)
	case *elements.Import:
		return v.VisitImport(s)
	default:
		panic(fmt.Errorf("you should implement %s statement", reflect.TypeOf(statement)))
	}
}

func (v *ElementsVisitor) VisitExpressions(expressions []elements.Expression) string {
	p := make([]string, len(expressions))
	for i, exp := range expressions {
		p[i] = v.VisitExpression(exp)
	}
	return strings.Join(p, "\n")
}

func (v *ElementsVisitor) VisitExpression(expression elements.Expression) string {
	switch exp := expression.(type) {
	case *elements.Arguments:
		return v.VisitArguments(exp)
	case *elements.Parenthesized:
		return v.VisitParenthesized(exp)
	case string:
		return exp
	case []elements.Expression:
		return v.VisitExpressions(exp)
	case *elements.MemberDot:
		return v.VisitMemberDot(exp)
	default:
		panic(fmt.Errorf("you should implement %s expression", reflect.TypeOf(exp)))
	}
}

func (v *ElementsVisitor) VisitVariables(variables []*elements.Variable) string {
	p := make([]string, len(variables))
	for i, va := range variables {
		p[i] = v.VisitVariable(va)
	}
	return strings.Join(p, "\n")
}

func (v *ElementsVisitor) VisitVariable(variable *elements.Variable) string {
	bytes, _ := json.Marshal(variable)
	fmt.Println(string(bytes))
	return ""
}

func (v *ElementsVisitor) VisitParameter(param *elements.Parameter) string {
	if len(param.Type) == 0 {
		param.Type = "interface{}"
	}
	return fmt.Sprintf("%s %s", param.Identifier, param.Type)
}

func (v *ElementsVisitor) VisitParameters(params []*elements.Parameter) string {
	p := make([]string, len(params))
	for i, param := range params {
		p[i] = v.VisitParameter(param)
	}
	return strings.Join(p, ", ")
}

func (v *ElementsVisitor) VisitBlock(block *elements.Block) string {
	if block == nil {
		return ""
	}
	return v.VisitStatements(block.Statements)
}

func (v *ElementsVisitor) VisitFunction(function *elements.Function) string {
	if len(function.Signature.Type) > 0 {
		function.Signature.Type = " " + function.Signature.Type
	}
	return fmt.Sprintf("func %s(%s)%s {\n%s\n}",
		function.Identifier,
		v.VisitParameters(function.Signature.Parameters),
		function.Signature.Type,
		v.VisitBlock(function.Body),
	)
}

func (v *ElementsVisitor) VisitProgram(program *elements.Program) string {
	return v.VisitStatements(program.Statements)
}

var _ elements.Visitor = &ElementsVisitor{}
