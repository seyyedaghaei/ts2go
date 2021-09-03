package visitor // TypeScriptParser
import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/seyyedaghaei/ts2go/ast"
	"github.com/seyyedaghaei/ts2go/elements"
	"reflect"
)

type TypeScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *TypeScriptVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	panic(fmt.Errorf("you should implement %s type", reflect.TypeOf(node)))
}

func (v *TypeScriptVisitor) VisitInitializer(ctx *ast.InitializerContext) interface{} {
	return &elements.Initializer{Expression: ctx.SingleExpression().Accept(v).(elements.Expression)}
}

func (v *TypeScriptVisitor) VisitBindingPattern(ctx *ast.BindingPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeParameters(ctx *ast.TypeParametersContext) interface{} {
	if ctx.TypeParameterList() != nil {
		return ctx.TypeParameterList().Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeParameterList(ctx *ast.TypeParameterListContext) interface{} {
	params := make([]*elements.TypeParameter, 0)
	for _, param := range ctx.AllTypeParameter() {
		params = append(params, param.Accept(v).(*elements.TypeParameter))
	}
	return params
}

func (v *TypeScriptVisitor) VisitTypeParameter(ctx *ast.TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitConstraint(ctx *ast.ConstraintContext) interface{} {
	return ctx.Type_().Accept(v) // TODO: Maybe you should change this
}

func (v *TypeScriptVisitor) VisitTypeArguments(ctx *ast.TypeArgumentsContext) interface{} {
	if ctx.TypeArgumentList() != nil {
		return ctx.TypeArgumentList().Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeArgumentList(ctx *ast.TypeArgumentListContext) interface{} {
	typeArgs := make([]string, 0)
	for _, arg := range ctx.AllTypeArgument() {
		typeArgs = append(typeArgs, arg.Accept(v).(string))
	}
	return typeArgs
}

func (v *TypeScriptVisitor) VisitTypeArgument(ctx *ast.TypeArgumentContext) interface{} {
	return ctx.Type_().Accept(v)
}

func (v *TypeScriptVisitor) VisitType_(ctx *ast.Type_Context) interface{} {
	return ctx.GetText() // TODO: Change this
}

func (v *TypeScriptVisitor) VisitIntersection(ctx *ast.IntersectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPrimary(ctx *ast.PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitUnion(ctx *ast.UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitRedefinitionOfType(ctx *ast.RedefinitionOfTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPredefinedPrimType(ctx *ast.PredefinedPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrayPrimType(ctx *ast.ArrayPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitParenthesizedPrimType(ctx *ast.ParenthesizedPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitThisPrimType(ctx *ast.ThisPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTuplePrimType(ctx *ast.TuplePrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitObjectPrimType(ctx *ast.ObjectPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitReferencePrimType(ctx *ast.ReferencePrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitQueryPrimType(ctx *ast.QueryPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPredefinedType(ctx *ast.PredefinedTypeContext) interface{} {
	return v.returnNonNilText(
		ctx.Any(),
		ctx.Number(),
		ctx.Boolean(),
		ctx.Str(),
		ctx.Symbol(),
		ctx.Void(),
	) // TODO: Change this
}

func (v *TypeScriptVisitor) VisitTypeReference(ctx *ast.TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitNestedTypeGeneric(ctx *ast.NestedTypeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeGeneric(ctx *ast.TypeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeIncludeGeneric(ctx *ast.TypeIncludeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeName(ctx *ast.TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitObjectType(ctx *ast.ObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeBody(ctx *ast.TypeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeMemberList(ctx *ast.TypeMemberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeMember(ctx *ast.TypeMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrayType(ctx *ast.ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTupleType(ctx *ast.TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTupleElementTypes(ctx *ast.TupleElementTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFunctionType(ctx *ast.FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitConstructorType(ctx *ast.ConstructorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeQuery(ctx *ast.TypeQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeQueryExpression(ctx *ast.TypeQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertySignatur(ctx *ast.PropertySignaturContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeAnnotation(ctx *ast.TypeAnnotationContext) interface{} {
	return ctx.Type_().Accept(v) // TODO: Change this
}

func (v *TypeScriptVisitor) VisitCallSignature(ctx *ast.CallSignatureContext) interface{} {
	parameters := make([]*elements.Parameter, 0)
	if p := ctx.ParameterList(); p != nil {
		parameters = p.Accept(v).([]*elements.Parameter)
	}
	typ := ""
	if ctx.TypeAnnotation() != nil {
		typ = ctx.TypeAnnotation().Accept(v).(string) // TODO: Change this
	}
	return &elements.CallSignature{
		Parameters: parameters,
		Type:       typ,
	}
}

func (v *TypeScriptVisitor) VisitParameterList(ctx *ast.ParameterListContext) interface{} {
	parameters := make([]*elements.Parameter, 0)
	for _, p := range ctx.AllParameter() {
		parameters = append(parameters, p.Accept(v).(*elements.Parameter))
	}
	if rest := ctx.RestParameter(); rest != nil {
		parameters = append(parameters, rest.Accept(v).(*elements.Parameter))
	}
	return parameters
}

func (v *TypeScriptVisitor) VisitRequiredParameterList(ctx *ast.RequiredParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitParameter(ctx *ast.ParameterContext) interface{} {
	return v.returnNonNil(ctx.RequiredParameter(), ctx.OptionalParameter())
}

func (v *TypeScriptVisitor) VisitOptionalParameter(ctx *ast.OptionalParameterContext) interface{} {
	parameter := &elements.Parameter{
		Identifier: ctx.IdentifierOrPattern().Accept(v).(string), // TODO: Change this
		Required:   false,
		Rest:       false,
	}
	if ctx.AccessibilityModifier() != nil {
		parameter.Accessibility = ctx.AccessibilityModifier().Accept(v).(elements.Accessibility)
	}
	if ctx.TypeAnnotation() != nil {
		parameter.Type = ctx.TypeAnnotation().Accept(v).(string) // TODO: Change this
	}
	if ctx.Initializer() != nil {
		parameter.Initializer = ctx.Initializer().Accept(v).(*elements.Initializer)
	}
	return parameter
}

func (v *TypeScriptVisitor) VisitRestParameter(ctx *ast.RestParameterContext) interface{} {
	typ := ""
	if ctx.TypeAnnotation() != nil {
		typ = ctx.TypeAnnotation().Accept(v).(string) // TODO: Change this
	}
	return &elements.Parameter{
		Identifier:    ctx.SingleExpression().Accept(v).(string), // TODO: Change this
		Required:      false,
		Rest:          true,
		Accessibility: elements.None,
		Type:          typ,
	}
}

func (v *TypeScriptVisitor) VisitRequiredParameter(ctx *ast.RequiredParameterContext) interface{} {
	a := elements.None
	if ctx.AccessibilityModifier() != nil {
		a = ctx.AccessibilityModifier().Accept(v).(elements.Accessibility)
	}
	typ := ""
	if ctx.TypeAnnotation() != nil {
		typ = ctx.TypeAnnotation().Accept(v).(string) // TODO: Change this
	}
	return &elements.Parameter{
		Identifier:    ctx.IdentifierOrPattern().Accept(v).(string), // TODO: Change this
		Required:      true,
		Rest:          false,
		Accessibility: a,
		Type:          typ,
	}
}

func (v *TypeScriptVisitor) VisitAccessibilityModifier(ctx *ast.AccessibilityModifierContext) interface{} {
	if ctx.Private() != nil {
		return elements.Private
	}
	if ctx.Public() != nil {
		return elements.Public
	}
	if ctx.Protected() != nil {
		return elements.Protected
	}
	return elements.None
}

func (v *TypeScriptVisitor) VisitIdentifierOrPattern(ctx *ast.IdentifierOrPatternContext) interface{} {
	return v.returnNonNil(
		ctx.IdentifierName(),
		ctx.BindingPattern(),
	)
}

func (v *TypeScriptVisitor) VisitConstructSignature(ctx *ast.ConstructSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIndexSignature(ctx *ast.IndexSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitMethodSignature(ctx *ast.MethodSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeAliasDeclaration(ctx *ast.TypeAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitConstructorDeclaration(ctx *ast.ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitInterfaceDeclaration(ctx *ast.InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitInterfaceExtendsClause(ctx *ast.InterfaceExtendsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassOrInterfaceTypeList(ctx *ast.ClassOrInterfaceTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitEnumDeclaration(ctx *ast.EnumDeclarationContext) interface{} {
	return &elements.Enum{
		Const:      ctx.Const() != nil,
		Identifier: ctx.Identifier().GetText(),
		Members:    ctx.EnumBody().Accept(v).([]*elements.EnumMember),
	}
}

func (v *TypeScriptVisitor) VisitEnumBody(ctx *ast.EnumBodyContext) interface{} {
	return ctx.EnumMemberList().Accept(v)
}

func (v *TypeScriptVisitor) VisitEnumMemberList(ctx *ast.EnumMemberListContext) interface{} {
	members := make([]elements.EnumMember, 0)
	for _, member := range ctx.AllEnumMember() {
		members = append(members, member.Accept(v).(elements.EnumMember))
	}
	return members
}

func (v *TypeScriptVisitor) VisitEnumMember(ctx *ast.EnumMemberContext) interface{} {
	return &elements.EnumMember{
		Name:       ctx.PropertyName().Accept(v).(string), // TODO: Maybe you should change this
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitNamespaceDeclaration(ctx *ast.NamespaceDeclarationContext) interface{} {
	return &elements.NameSpace{
		Names:      ctx.NamespaceName().Accept(v).([]string),
		Statements: ctx.StatementList().Accept(v).([]elements.Statement),
	}
}

func (v *TypeScriptVisitor) VisitNamespaceName(ctx *ast.NamespaceNameContext) interface{} {
	identifiers := make([]string, 0)
	for _, id := range ctx.AllIdentifier() {
		identifiers = append(identifiers, id.GetText())
	}
	return identifiers
}

func (v *TypeScriptVisitor) VisitImportAliasDeclaration(ctx *ast.ImportAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDecoratorList(ctx *ast.DecoratorListContext) interface{} {
	decorators := make([]*elements.Decorator, 0)
	for _, d := range ctx.AllDecorator() {
		decorators = append(decorators, d.Accept(v).(*elements.Decorator))
	}
	return decorators
}

func (v *TypeScriptVisitor) VisitDecorator(ctx *ast.DecoratorContext) interface{} {
	return &elements.Decorator{} // TODO: Complete this
}

func (v *TypeScriptVisitor) VisitDecoratorMemberExpression(ctx *ast.DecoratorMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDecoratorCallExpression(ctx *ast.DecoratorCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitProgram(ctx *ast.ProgramContext) interface{} {
	statements := make([]elements.Statement, 0)
	if statementList := ctx.StatementList(); statementList != nil {
		statements = statementList.Accept(v).([]elements.Statement)
	}
	return &elements.Program{Statements: statements}
}

func (v *TypeScriptVisitor) VisitStatement(ctx *ast.StatementContext) interface{} {
	return v.returnNonNil(
		ctx.Block(),
		ctx.ImportStatement(),
		ctx.ExportStatement(),
		ctx.EmptyStatement(),
		ctx.AbstractDeclaration(),
		ctx.DecoratorList(),
		ctx.ClassDeclaration(),
		ctx.InterfaceDeclaration(),
		ctx.NamespaceDeclaration(),
		ctx.IfStatement(),
		ctx.IterationStatement(),
		ctx.ContinueStatement(),
		ctx.BreakStatement(),
		ctx.ReturnStatement(),
		ctx.YieldStatement(),
		ctx.WithStatement(),
		ctx.LabelledStatement(),
		ctx.SwitchStatement(),
		ctx.ThrowStatement(),
		ctx.TryStatement(),
		ctx.DebuggerStatement(),
		ctx.FunctionDeclaration(),
		ctx.ArrowFunctionDeclaration(),
		ctx.GeneratorFunctionDeclaration(),
		ctx.VariableStatement(),
		ctx.TypeAliasDeclaration(),
		ctx.EnumDeclaration(),
		ctx.ExpressionStatement(),
	)
}

func (v *TypeScriptVisitor) VisitBlock(ctx *ast.BlockContext) interface{} {
	statements := make([]elements.Statement, 0)
	if statementList := ctx.StatementList(); statementList != nil {
		statements = statementList.Accept(v).([]elements.Statement)
	}
	return &elements.Block{Statements: statements}
}

func (v *TypeScriptVisitor) VisitStatementList(ctx *ast.StatementListContext) interface{} {
	statements := make([]elements.Statement, 0)
	for _, statement := range ctx.AllStatement() {
		statements = append(statements, statement.Accept(v).(elements.Statement))
	}
	return statements
}

func (v *TypeScriptVisitor) VisitAbstractDeclaration(ctx *ast.AbstractDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitImportStatement(ctx *ast.ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFromBlock(ctx *ast.FromBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitMultipleImportStatement(ctx *ast.MultipleImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitExportStatement(ctx *ast.ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitVariableStatement(ctx *ast.VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitVariableDeclarationList(ctx *ast.VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitVariableDeclaration(ctx *ast.VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitEmptyStatement(_ *ast.EmptyStatementContext) interface{} {
	return &elements.Empty{}
}

func (v *TypeScriptVisitor) VisitExpressionStatement(ctx *ast.ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIfStatement(ctx *ast.IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDoStatement(ctx *ast.DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitWhileStatement(ctx *ast.WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitForStatement(ctx *ast.ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitForVarStatement(ctx *ast.ForVarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitForInStatement(ctx *ast.ForInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitForVarInStatement(ctx *ast.ForVarInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitVarModifier(ctx *ast.VarModifierContext) interface{} {
	return "var" // TODO: Change this
}

func (v *TypeScriptVisitor) VisitContinueStatement(ctx *ast.ContinueStatementContext) interface{} {
	c := &elements.Continue{}
	if ctx.Identifier() != nil {
		c.Identifier = ctx.Identifier().GetText()
	}
	return c
}

func (v *TypeScriptVisitor) VisitBreakStatement(ctx *ast.BreakStatementContext) interface{} {
	b := &elements.Break{}
	if ctx.Identifier() != nil {
		b.Identifier = ctx.Identifier().GetText()
	}
	return b
}

func (v *TypeScriptVisitor) VisitReturnStatement(ctx *ast.ReturnStatementContext) interface{} {
	expressions := make([]elements.Expression, 0)
	if sequence := ctx.ExpressionSequence(); sequence != nil {
		expressions = sequence.Accept(v).([]elements.Expression)
	}
	return &elements.Return{
		Expressions: expressions,
	}
}

func (v *TypeScriptVisitor) VisitYieldStatement(ctx *ast.YieldStatementContext) interface{} {
	expressions := make([]elements.Expression, 0)
	if sequence := ctx.ExpressionSequence(); sequence != nil {
		expressions = sequence.Accept(v).([]elements.Expression)
	}
	return &elements.Yield{
		Expressions: expressions,
	}
}

func (v *TypeScriptVisitor) VisitWithStatement(ctx *ast.WithStatementContext) interface{} {
	return &elements.With{
		Expressions: ctx.ExpressionSequence().Accept(v).([]elements.Expression),
		Statement:   ctx.Statement().(elements.Statement),
	}
}

func (v *TypeScriptVisitor) VisitSwitchStatement(ctx *ast.SwitchStatementContext) interface{} {
	return &elements.Switch{
		Expressions: ctx.ExpressionSequence().Accept(v).([]elements.Expression),
		CaseClauses: ctx.CaseBlock().Accept(v).([]*elements.CaseClause),
	}
}

func (v *TypeScriptVisitor) VisitCaseBlock(ctx *ast.CaseBlockContext) interface{} {
	clauses := make([]*elements.CaseClause, 0)
	for _, clause := range ctx.AllCaseClauses() {
		clauses = append(clauses, clause.Accept(v).(*elements.CaseClause))
	}
	if ctx.DefaultClause() != nil {
		clauses = append(clauses, ctx.DefaultClause().Accept(v).(*elements.CaseClause))
	}
	return clauses
}

func (v *TypeScriptVisitor) VisitCaseClauses(ctx *ast.CaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitCaseClause(ctx *ast.CaseClauseContext) interface{} {
	return &elements.CaseClause{
		Expressions: ctx.ExpressionSequence().Accept(v).([]elements.Expression),
		Statements:  ctx.StatementList().Accept(v).([]elements.Statement),
	}
}

func (v *TypeScriptVisitor) VisitDefaultClause(ctx *ast.DefaultClauseContext) interface{} {
	return &elements.CaseClause{
		Statements: ctx.StatementList().Accept(v).([]elements.Statement),
	}
}

func (v *TypeScriptVisitor) VisitLabelledStatement(ctx *ast.LabelledStatementContext) interface{} {
	return &elements.LabeledStatement{
		Identifier: ctx.Identifier().GetText(),
		Statement:  ctx.Statement().Accept(v).(elements.Statement),
	}
}

func (v *TypeScriptVisitor) VisitThrowStatement(ctx *ast.ThrowStatementContext) interface{} {
	return &elements.Throw{Expressions: ctx.Throw().Accept(v).([]elements.Expression)}
}

func (v *TypeScriptVisitor) VisitTryStatement(ctx *ast.TryStatementContext) interface{} {
	try := &elements.Try{
		Block: ctx.Block().Accept(v).(*elements.Block),
	}
	if catch := ctx.CatchProduction(); catch != nil {
		try.Catch = catch.Accept(v).(*elements.Catch)
	}
	if finally := ctx.CatchProduction(); finally != nil {
		try.Finally = finally.Accept(v).(*elements.Finally)
	}
	return try
}

func (v *TypeScriptVisitor) VisitCatchProduction(ctx *ast.CatchProductionContext) interface{} {
	return &elements.Catch{
		Identifier: ctx.Identifier().GetText(),
		Block:      ctx.Block().Accept(v).(*elements.Block),
	}
}

func (v *TypeScriptVisitor) VisitFinallyProduction(ctx *ast.FinallyProductionContext) interface{} {
	return &elements.Finally{Block: ctx.Block().Accept(v).(*elements.Block)}
}

func (v *TypeScriptVisitor) VisitDebuggerStatement(ctx *ast.DebuggerStatementContext) interface{} {
	return &elements.Debugger{}
}

func (v *TypeScriptVisitor) VisitFunctionDeclaration(ctx *ast.FunctionDeclarationContext) interface{} {
	return &elements.Function{
		Identifier: ctx.Identifier().GetText(),
		Signature:  ctx.CallSignature().Accept(v).(*elements.CallSignature),
	}
}

func (v *TypeScriptVisitor) VisitClassDeclaration(ctx *ast.ClassDeclarationContext) interface{} {
	return &elements.Class{
		Abstract:   ctx.Abstract() != nil,
		Identifier: ctx.Identifier().GetText(),
		Elements:   ctx.ClassTail().Accept(v).([]elements.ClassElement),
	}
}

func (v *TypeScriptVisitor) VisitClassHeritage(ctx *ast.ClassHeritageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassTail(ctx *ast.ClassTailContext) interface{} {
	classElements := make([]elements.ClassElement, 0)
	for _, element := range ctx.AllClassElement() {
		classElements = append(classElements, element.Accept(v).(elements.ClassElement))
	}
	return classElements
}

func (v *TypeScriptVisitor) VisitClassExtendsClause(ctx *ast.ClassExtendsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitImplementsClause(ctx *ast.ImplementsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassElement(ctx *ast.ClassElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertyDeclarationExpression(ctx *ast.PropertyDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitMethodDeclarationExpression(ctx *ast.MethodDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGetterSetterDeclarationExpression(ctx *ast.GetterSetterDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitAbstractMemberDeclaration(ctx *ast.AbstractMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertyMemberBase(ctx *ast.PropertyMemberBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIndexMemberDeclaration(ctx *ast.IndexMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGeneratorMethod(ctx *ast.GeneratorMethodContext) interface{} {
	g := &elements.Generator{
		Identifier: ctx.Identifier().GetText(),
		Body:       ctx.FunctionBody().Accept(v).(*elements.Block),
	}
	if ctx.FormalParameterList() != nil {
		g.Parameters = ctx.FormalParameterList().Accept(v).([]*elements.FormalParameter)
	}
	return g
}

func (v *TypeScriptVisitor) VisitGeneratorFunctionDeclaration(ctx *ast.GeneratorFunctionDeclarationContext) interface{} {
	g := &elements.Generator{
		Identifier: ctx.Identifier().GetText(),
		Body:       ctx.FunctionBody().Accept(v).(*elements.Block),
	}
	if ctx.FormalParameterList() != nil {
		g.Parameters = ctx.FormalParameterList().Accept(v).([]*elements.FormalParameter)
	}
	return g
}

func (v *TypeScriptVisitor) VisitGeneratorBlock(ctx *ast.GeneratorBlockContext) interface{} {
	generators := make([]*elements.Iterator, 0)
	for _, g := range ctx.AllGeneratorDefinition() {
		generators = append(generators, g.Accept(v).(*elements.Iterator))
	}
	return generators
}

func (v *TypeScriptVisitor) VisitGeneratorDefinition(ctx *ast.GeneratorDefinitionContext) interface{} {
	return ctx.IteratorDefinition().Accept(v)
}

func (v *TypeScriptVisitor) VisitIteratorBlock(ctx *ast.IteratorBlockContext) interface{} {
	iterators := make([]*elements.Iterator, 0)
	for _, i := range ctx.AllIteratorDefinition() {
		iterators = append(iterators, i.Accept(v).(*elements.Iterator))
	}
	return iterators
}

func (v *TypeScriptVisitor) VisitIteratorDefinition(ctx *ast.IteratorDefinitionContext) interface{} {
	i := &elements.Iterator{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
		Body:       ctx.FunctionBody().Accept(v).(*elements.Block),
	}
	if ctx.FormalParameterList() != nil {
		i.Parameters = ctx.FormalParameterList().Accept(v).([]*elements.FormalParameter)
	}
	return i
}

func (v *TypeScriptVisitor) VisitFormalParameterList(ctx *ast.FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFormalParameterArg(ctx *ast.FormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLastFormalParameterArg(ctx *ast.LastFormalParameterArgContext) interface{} {
	return ctx.Identifier().GetText() // TODO: Maybe you should change this
}

func (v *TypeScriptVisitor) VisitFunctionBody(ctx *ast.FunctionBodyContext) interface{} {
	statements := make([]elements.Statement, 0)
	if statementList := ctx.StatementList(); statementList != nil {
		statements = statementList.Accept(v).([]elements.Statement)
	}
	return &elements.Block{Statements: statements}
}

func (v *TypeScriptVisitor) VisitArrayLiteral(ctx *ast.ArrayLiteralContext) interface{} {
	elems := make([]*elements.ArrayElement, 0)
	if e := ctx.ElementList(); e != nil {
		elems = e.Accept(v).([]*elements.ArrayElement)
	}
	return &elements.Array{Elements: elems}
}

func (v *TypeScriptVisitor) VisitElementList(ctx *ast.ElementListContext) interface{} {
	elems := make([]*elements.ArrayElement, 0)
	for _, e := range ctx.AllArrayElement() {
		elems = append(elems, e.Accept(v).(*elements.ArrayElement))
	}
	return elems
}

func (v *TypeScriptVisitor) VisitArrayElement(ctx *ast.ArrayElementContext) interface{} {
	element := &elements.ArrayElement{Ellipsis: ctx.Ellipsis() != nil}
	if ctx.Identifier() != nil {
		element.Identifier = ctx.Identifier().GetText()
	}
	if ctx.SingleExpression() != nil {
		element.Expression = ctx.SingleExpression().Accept(v).(elements.Expression)
	}
	return element
}

func (v *TypeScriptVisitor) VisitObjectLiteral(ctx *ast.ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertyExpressionAssignment(ctx *ast.PropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitComputedPropertyExpressionAssignment(ctx *ast.ComputedPropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertyGetter(ctx *ast.PropertyGetterContext) interface{} {
	return ctx.GetAccessor().Accept(v)
}

func (v *TypeScriptVisitor) VisitPropertySetter(ctx *ast.PropertySetterContext) interface{} {
	return ctx.SetAccessor().Accept(v)
}

func (v *TypeScriptVisitor) VisitMethodProperty(ctx *ast.MethodPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertyShorthand(ctx *ast.PropertyShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitRestParameterInObject(ctx *ast.RestParameterInObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGetAccessor(ctx *ast.GetAccessorContext) interface{} {
	return &elements.Getter{
		Name: ctx.Getter().Accept(v).(string),         // TODO: Change this
		Type: ctx.TypeAnnotation().Accept(v).(string), // TODO: Change this
		Body: ctx.FunctionBody().Accept(v).(*elements.Block),
	}
}

func (v *TypeScriptVisitor) VisitSetAccessor(ctx *ast.SetAccessorContext) interface{} {
	return &elements.Setter{
		Name:     ctx.Setter().Accept(v).(string),         // TODO: Change this
		Type:     ctx.TypeAnnotation().Accept(v).(string), // TODO: Change this
		Argument: ctx.Identifier().Accept(v).(string),     // TODO: Absolutely change this
		Body:     ctx.FunctionBody().Accept(v).(*elements.Block),
	}
}

func (v *TypeScriptVisitor) VisitPropertyName(ctx *ast.PropertyNameContext) interface{} {
	return ctx.GetText() // TODO: Change this
}

func (v *TypeScriptVisitor) VisitArguments(ctx *ast.ArgumentsContext) interface{} {
	return ctx.ArgumentList().Accept(v)
}

func (v *TypeScriptVisitor) VisitArgumentList(ctx *ast.ArgumentListContext) interface{} {
	arguments := make([]elements.Statement, 0)
	for _, arg := range ctx.AllArgument() {
		arguments = append(arguments, arg.Accept(v).(elements.Statement))
	}
	return arguments
}

func (v *TypeScriptVisitor) VisitArgument(ctx *ast.ArgumentContext) interface{} {
	return &elements.Argument{
		Rest:       ctx.Ellipsis() != nil,
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
		Identifier: ctx.Identifier().GetText(),
	}
}

func (v *TypeScriptVisitor) VisitExpressionSequence(ctx *ast.ExpressionSequenceContext) interface{} {
	expressions := make([]elements.Expression, 0)
	for _, expression := range ctx.AllSingleExpression() {
		expressions = append(expressions, expression.Accept(v).(elements.Expression))
	}
	return expressions
}

func (v *TypeScriptVisitor) VisitFunctionExpressionDeclaration(ctx *ast.FunctionExpressionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTemplateStringExpression(ctx *ast.TemplateStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTernaryExpression(ctx *ast.TernaryExpressionContext) interface{} {
	return &elements.Ternary{
		Question:        ctx.SingleExpression(0).Accept(v).(elements.Expression),
		TrueExpression:  ctx.SingleExpression(1).Accept(v).(elements.Expression),
		FalseExpression: ctx.SingleExpression(2).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitLogicalAndExpression(ctx *ast.LogicalAndExpressionContext) interface{} {
	return &elements.LogicalAnd{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitGeneratorsExpression(ctx *ast.GeneratorsExpressionContext) interface{} {
	return ctx.GeneratorBlock().Accept(v)
}

func (v *TypeScriptVisitor) VisitPreIncrementExpression(ctx *ast.PreIncrementExpressionContext) interface{} {
	return &elements.PreIncrement{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitObjectLiteralExpression(ctx *ast.ObjectLiteralExpressionContext) interface{} {
	return ctx.ObjectLiteral().Accept(v)
}

func (v *TypeScriptVisitor) VisitInExpression(ctx *ast.InExpressionContext) interface{} {
	return &elements.In{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitLogicalOrExpression(ctx *ast.LogicalOrExpressionContext) interface{} {
	return &elements.LogicalOr{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitGenericTypes(ctx *ast.GenericTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitNotExpression(ctx *ast.NotExpressionContext) interface{} {
	return &elements.Not{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitPreDecreaseExpression(ctx *ast.PreDecreaseExpressionContext) interface{} {
	return &elements.PreDecrease{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitArgumentsExpression(ctx *ast.ArgumentsExpressionContext) interface{} {
	return &elements.Arguments{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
		Arguments:  ctx.Arguments().Accept(v).([]*elements.Argument),
	}
}

func (v *TypeScriptVisitor) VisitThisExpression(ctx *ast.ThisExpressionContext) interface{} {
	return &elements.This{}
}

func (v *TypeScriptVisitor) VisitFunctionExpression(ctx *ast.FunctionExpressionContext) interface{} {
	return ctx.FunctionExpressionDeclaration().Accept(v)
}

func (v *TypeScriptVisitor) VisitUnaryMinusExpression(ctx *ast.UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitAssignmentExpression(ctx *ast.AssignmentExpressionContext) interface{} {
	return &elements.Assignment{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitPostDecreaseExpression(ctx *ast.PostDecreaseExpressionContext) interface{} {
	return &elements.PostDecrease{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitTypeofExpression(ctx *ast.TypeofExpressionContext) interface{} {
	return &elements.TypeOf{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitInstanceofExpression(ctx *ast.InstanceofExpressionContext) interface{} {
	return &elements.InstanceOf{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitUnaryPlusExpression(ctx *ast.UnaryPlusExpressionContext) interface{} {
	return &elements.UnaryPlus{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitDeleteExpression(ctx *ast.DeleteExpressionContext) interface{} {
	return &elements.Delete{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitGeneratorsFunctionExpression(ctx *ast.GeneratorsFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrowFunctionExpression(ctx *ast.ArrowFunctionExpressionContext) interface{} {
	return ctx.ArrowFunctionDeclaration().Accept(v)
}

func (v *TypeScriptVisitor) VisitIteratorsExpression(ctx *ast.IteratorsExpressionContext) interface{} {
	return ctx.Accept(v)
}

func (v *TypeScriptVisitor) VisitEqualityExpression(ctx *ast.EqualityExpressionContext) interface{} {
	return &elements.Equality{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
		Sign:            elements.Sign(v.returnNonNilText(ctx.Equals_(), ctx.NotEquals(), ctx.IdentityEquals(), ctx.IdentityNotEquals()).(string)),
	}
}

func (v *TypeScriptVisitor) VisitBitXOrExpression(ctx *ast.BitXOrExpressionContext) interface{} {
	return &elements.BitXOr{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitCastAsExpression(ctx *ast.CastAsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitSuperExpression(ctx *ast.SuperExpressionContext) interface{} {
	return &elements.Super{}
}

func (v *TypeScriptVisitor) VisitMultiplicativeExpression(ctx *ast.MultiplicativeExpressionContext) interface{} {
	return &elements.Multiplicative{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
		Sign:            elements.Sign(v.returnNonNilText(ctx.Multiply(), ctx.Divide(), ctx.Modulus()).(string)),
	}
}

func (v *TypeScriptVisitor) VisitBitShiftExpression(ctx *ast.BitShiftExpressionContext) interface{} {
	return &elements.BitShift{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
		Sign:            elements.Sign(v.returnNonNilText(ctx.LeftShiftArithmetic(), ctx.RightShiftArithmetic(), ctx.RightShiftLogical()).(string)),
	}
}

func (v *TypeScriptVisitor) VisitParenthesizedExpression(ctx *ast.ParenthesizedExpressionContext) interface{} {
	return &elements.Parenthesized{Expressions: ctx.ExpressionSequence().Accept(v).([]elements.Expression)}
}

func (v *TypeScriptVisitor) VisitAdditiveExpression(ctx *ast.AdditiveExpressionContext) interface{} {
	return &elements.Additive{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
		Sign:            elements.Sign(v.returnNonNilText(ctx.Plus(), ctx.Minus()).(string)),
	}
}

func (v *TypeScriptVisitor) VisitRelationalExpression(ctx *ast.RelationalExpressionContext) interface{} {
	return &elements.Relational{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
		Sign:            elements.Sign(v.returnNonNilText(ctx.LessThan(), ctx.MoreThan(), ctx.LessThanEquals(), ctx.GreaterThanEquals()).(string)),
	}
}

func (v *TypeScriptVisitor) VisitPostIncrementExpression(ctx *ast.PostIncrementExpressionContext) interface{} {
	return &elements.PostIncrement{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitYieldExpression(ctx *ast.YieldExpressionContext) interface{} {
	return ctx.YieldStatement().Accept(v)
}

func (v *TypeScriptVisitor) VisitBitNotExpression(ctx *ast.BitNotExpressionContext) interface{} {
	return &elements.BitNot{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitNewExpression(ctx *ast.NewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLiteralExpression(ctx *ast.LiteralExpressionContext) interface{} {
	return ctx.Literal().Accept(v)
}

func (v *TypeScriptVisitor) VisitArrayLiteralExpression(ctx *ast.ArrayLiteralExpressionContext) interface{} {
	return ctx.ArrayLiteral().Accept(v)
}

func (v *TypeScriptVisitor) VisitMemberDotExpression(ctx *ast.MemberDotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassExpression(ctx *ast.ClassExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitMemberIndexExpression(ctx *ast.MemberIndexExpressionContext) interface{} {
	return &elements.MemberIndex{
		Expression: ctx.SingleExpression().Accept(v).(elements.Expression),
		Index:      ctx.ExpressionSequence().Accept(v).([]elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitIdentifierExpression(ctx *ast.IdentifierExpressionContext) interface{} {
	// TODO: You must change and complete this
	return ctx.IdentifierName().Accept(v)
}

func (v *TypeScriptVisitor) VisitBitAndExpression(ctx *ast.BitAndExpressionContext) interface{} {
	return &elements.BitAnd{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitBitOrExpression(ctx *ast.BitOrExpressionContext) interface{} {
	return &elements.BitOr{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
	}
}

func (v *TypeScriptVisitor) VisitAssignmentOperatorExpression(ctx *ast.AssignmentOperatorExpressionContext) interface{} {
	return &elements.AssignmentOperator{
		LeftExpression:  ctx.SingleExpression(0).Accept(v).(elements.Expression),
		RightExpression: ctx.SingleExpression(1).Accept(v).(elements.Expression),
		Sign:            ctx.AssignmentOperator().Accept(v).(elements.Sign),
	}
}

func (v *TypeScriptVisitor) VisitVoidExpression(ctx *ast.VoidExpressionContext) interface{} {
	return &elements.Void{Expression: ctx.SingleExpression().Accept(v).([]elements.Expression)}
}

func (v *TypeScriptVisitor) VisitAsExpression(ctx *ast.AsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrowFunctionDeclaration(ctx *ast.ArrowFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrowFunctionParameters(ctx *ast.ArrowFunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrowFunctionBody(ctx *ast.ArrowFunctionBodyContext) interface{} {
	if ctx.SingleExpression() != nil {
		return &elements.Block{Statements: []elements.Statement{ctx.SingleExpression().Accept(v).(elements.Statement)}}
	}
	return ctx.FunctionBody().Accept(v)
}

func (v *TypeScriptVisitor) VisitAssignmentOperator(ctx *ast.AssignmentOperatorContext) interface{} {
	return elements.Sign(ctx.GetText()) // TODO: Change this
}

func (v *TypeScriptVisitor) VisitLiteral(ctx *ast.LiteralContext) interface{} {
	// TODO: Complete this
	if ctx.NullLiteral() != nil {
		return &elements.Null{}
	}
	if ctx.BooleanLiteral() != nil {
		return ctx.BooleanLiteral().GetText() // TODO: You must change this
	}
	if ctx.TemplateStringLiteral() != nil {
		ctx.TemplateStringLiteral().Accept(v)
	}
	if ctx.NumericLiteral() != nil {
		return ctx.NumericLiteral().Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTemplateStringLiteral(ctx *ast.TemplateStringLiteralContext) interface{} {
	// TODO: Maybe you should change this
	atoms := make([]interface{}, 0) // TODO: Type
	for _, atom := range ctx.AllTemplateStringAtom() {
		atoms = append(atoms, atom.Accept(v)) // TODO: cast
	}
	return atoms
}

func (v *TypeScriptVisitor) VisitTemplateStringAtom(ctx *ast.TemplateStringAtomContext) interface{} {
	// TODO: You should change this
	if ctx.TemplateStringAtom() != nil {
		return ctx.TemplateStringAtom().GetText()
	}
	return ctx.SingleExpression().Accept(v)
}

func (v *TypeScriptVisitor) VisitNumericLiteral(ctx *ast.NumericLiteralContext) interface{} {
	return ctx.GetText() // TODO: Change this
}

func (v *TypeScriptVisitor) VisitIdentifierName(ctx *ast.IdentifierNameContext) interface{} {
	if ctx.Identifier() != nil {
		return ctx.Identifier().GetText()
	}
	return ctx.ReservedWord().Accept(v).(string) // TODO: Change this
}

func (v *TypeScriptVisitor) VisitIdentifierOrKeyWord(ctx *ast.IdentifierOrKeyWordContext) interface{} {
	return v.returnNonNilText(
		ctx.Identifier(),
		ctx.TypeAlias(),
		ctx.Require(),
	) // TODO: You probably wanna change this
}

func (v *TypeScriptVisitor) VisitReservedWord(ctx *ast.ReservedWordContext) interface{} {
	text := v.returnNonNilText(ctx.NullLiteral(), ctx.BooleanLiteral())
	if text != nil {
		return text
	}
	return ctx.Keyword().Accept(v) // TODO: You probably wanna change this
}

func (v *TypeScriptVisitor) VisitKeyword(ctx *ast.KeywordContext) interface{} {
	return v.returnNonNilText(
		ctx.Break(),
		ctx.Do(),
		ctx.Instanceof(),
		ctx.Typeof(),
		ctx.Case(),
		ctx.Else(),
		ctx.New(),
		ctx.Var(),
		ctx.Catch(),
		ctx.Finally(),
		ctx.Return(),
		ctx.Void(),
		ctx.Continue(),
		ctx.For(),
		ctx.Switch(),
		ctx.While(),
		ctx.Debugger(),
		ctx.Function_(),
		ctx.This(),
		ctx.With(),
		ctx.Default(),
		ctx.If(),
		ctx.Throw(),
		ctx.Delete(),
		ctx.In(),
		ctx.Try(),
		ctx.ReadOnly(),
		ctx.Async(),
		ctx.From(),
		ctx.Class(),
		ctx.Enum(),
		ctx.Extends(),
		ctx.Super(),
		ctx.Const(),
		ctx.Export(),
		ctx.Import(),
		ctx.Implements(),
		ctx.Let(),
		ctx.Private(),
		ctx.Public(),
		ctx.Interface(),
		ctx.Package(),
		ctx.Protected(),
		ctx.Static(),
		ctx.Yield(),
		ctx.Get(),
		ctx.Set(),
		ctx.Require(),
		ctx.TypeAlias(),
		ctx.Str(),
	) // TODO: You probably wanna change this
}

func (v *TypeScriptVisitor) VisitGetter(ctx *ast.GetterContext) interface{} {
	return ctx.PropertyName().Accept(v)
}

func (v *TypeScriptVisitor) VisitSetter(ctx *ast.SetterContext) interface{} {
	return ctx.PropertyName().Accept(v)
}

func (v *TypeScriptVisitor) VisitEos(ctx *ast.EosContext) interface{} {
	return &elements.EOS{}
}
