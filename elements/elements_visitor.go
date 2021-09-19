package elements

type Visitor interface {
	VisitProgram(program *Program) string
	VisitFunction(function *Function) string
	VisitParameter(param *Parameter) string
	VisitParameters(params []*Parameter) string
	VisitBlock(block *Block) string
	VisitExpressions(expressions []Expression) string
	VisitExpression(expression Expression) string
	VisitStatements(statements []Statement) string
	VisitStatement(statement Statement) string
	VisitVariable(variable *Variable) string
	VisitVariables(variables []*Variable) string
	VisitArguments(arguments *Arguments) string
	VisitArgument(argument *Argument) string
	VisitParenthesized(parenthesized *Parenthesized) string
	VisitMemberDot(memberDot *MemberDot) string
	VisitImport(imp *Import) string
}
