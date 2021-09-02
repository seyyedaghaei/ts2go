package ast // TypeScriptParser
import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/seyyedaghaei/ts2go/ast"
	"github.com/seyyedaghaei/ts2go/elements"
)

type TypeScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *TypeScriptVisitor) VisitInitializer(ctx *ast.InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBindingPattern(ctx *ast.BindingPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeParameters(ctx *ast.TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeParameterList(ctx *ast.TypeParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeParameter(ctx *ast.TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitConstraint(ctx *ast.ConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeArguments(ctx *ast.TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeArgumentList(ctx *ast.TypeArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeArgument(ctx *ast.TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitType_(ctx *ast.Type_Context) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitCallSignature(ctx *ast.CallSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitParameterList(ctx *ast.ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitRequiredParameterList(ctx *ast.RequiredParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitParameter(ctx *ast.ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitOptionalParameter(ctx *ast.OptionalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitRestParameter(ctx *ast.RestParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitRequiredParameter(ctx *ast.RequiredParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitAccessibilityModifier(ctx *ast.AccessibilityModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIdentifierOrPattern(ctx *ast.IdentifierOrPatternContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitEnumBody(ctx *ast.EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitEnumMemberList(ctx *ast.EnumMemberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitEnumMember(ctx *ast.EnumMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitNamespaceDeclaration(ctx *ast.NamespaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitNamespaceName(ctx *ast.NamespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitImportAliasDeclaration(ctx *ast.ImportAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDecoratorList(ctx *ast.DecoratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDecorator(ctx *ast.DecoratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDecoratorMemberExpression(ctx *ast.DecoratorMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDecoratorCallExpression(ctx *ast.DecoratorCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitProgram(ctx *ast.ProgramContext) interface{} {
	return &elements.Program{Statements: ctx.StatementList().Accept(v).([]elements.Statement)}
}

func (v *TypeScriptVisitor) VisitStatement(ctx *ast.StatementContext) interface{} {
	statement := v.returnNonNil(
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBlock(ctx *ast.BlockContext) interface{} {
	return v.VisitChildren(ctx)
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

func (v *TypeScriptVisitor) VisitEmptyStatement(ctx *ast.EmptyStatementContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitContinueStatement(ctx *ast.ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBreakStatement(ctx *ast.BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitReturnStatement(ctx *ast.ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitYieldStatement(ctx *ast.YieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitWithStatement(ctx *ast.WithStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitSwitchStatement(ctx *ast.SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitCaseBlock(ctx *ast.CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitCaseClauses(ctx *ast.CaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitCaseClause(ctx *ast.CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDefaultClause(ctx *ast.DefaultClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLabelledStatement(ctx *ast.LabelledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitThrowStatement(ctx *ast.ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTryStatement(ctx *ast.TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitCatchProduction(ctx *ast.CatchProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFinallyProduction(ctx *ast.FinallyProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDebuggerStatement(ctx *ast.DebuggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFunctionDeclaration(ctx *ast.FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassDeclaration(ctx *ast.ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassHeritage(ctx *ast.ClassHeritageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassTail(ctx *ast.ClassTailContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGeneratorFunctionDeclaration(ctx *ast.GeneratorFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGeneratorBlock(ctx *ast.GeneratorBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGeneratorDefinition(ctx *ast.GeneratorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIteratorBlock(ctx *ast.IteratorBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIteratorDefinition(ctx *ast.IteratorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFormalParameterList(ctx *ast.FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFormalParameterArg(ctx *ast.FormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLastFormalParameterArg(ctx *ast.LastFormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFunctionBody(ctx *ast.FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrayLiteral(ctx *ast.ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitElementList(ctx *ast.ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrayElement(ctx *ast.ArrayElementContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertySetter(ctx *ast.PropertySetterContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitSetAccessor(ctx *ast.SetAccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPropertyName(ctx *ast.PropertyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArguments(ctx *ast.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArgumentList(ctx *ast.ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArgument(ctx *ast.ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitExpressionSequence(ctx *ast.ExpressionSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFunctionExpressionDeclaration(ctx *ast.FunctionExpressionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTemplateStringExpression(ctx *ast.TemplateStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTernaryExpression(ctx *ast.TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLogicalAndExpression(ctx *ast.LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGeneratorsExpression(ctx *ast.GeneratorsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPreIncrementExpression(ctx *ast.PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitObjectLiteralExpression(ctx *ast.ObjectLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitInExpression(ctx *ast.InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLogicalOrExpression(ctx *ast.LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGenericTypes(ctx *ast.GenericTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitNotExpression(ctx *ast.NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPreDecreaseExpression(ctx *ast.PreDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArgumentsExpression(ctx *ast.ArgumentsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitThisExpression(ctx *ast.ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitFunctionExpression(ctx *ast.FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitUnaryMinusExpression(ctx *ast.UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitAssignmentExpression(ctx *ast.AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPostDecreaseExpression(ctx *ast.PostDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTypeofExpression(ctx *ast.TypeofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitInstanceofExpression(ctx *ast.InstanceofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitUnaryPlusExpression(ctx *ast.UnaryPlusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitDeleteExpression(ctx *ast.DeleteExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGeneratorsFunctionExpression(ctx *ast.GeneratorsFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrowFunctionExpression(ctx *ast.ArrowFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIteratorsExpression(ctx *ast.IteratorsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitEqualityExpression(ctx *ast.EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBitXOrExpression(ctx *ast.BitXOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitCastAsExpression(ctx *ast.CastAsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitSuperExpression(ctx *ast.SuperExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitMultiplicativeExpression(ctx *ast.MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBitShiftExpression(ctx *ast.BitShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitParenthesizedExpression(ctx *ast.ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitAdditiveExpression(ctx *ast.AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitRelationalExpression(ctx *ast.RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitPostIncrementExpression(ctx *ast.PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitYieldExpression(ctx *ast.YieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBitNotExpression(ctx *ast.BitNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitNewExpression(ctx *ast.NewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLiteralExpression(ctx *ast.LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitArrayLiteralExpression(ctx *ast.ArrayLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitMemberDotExpression(ctx *ast.MemberDotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitClassExpression(ctx *ast.ClassExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitMemberIndexExpression(ctx *ast.MemberIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIdentifierExpression(ctx *ast.IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBitAndExpression(ctx *ast.BitAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitBitOrExpression(ctx *ast.BitOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitAssignmentOperatorExpression(ctx *ast.AssignmentOperatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitVoidExpression(ctx *ast.VoidExpressionContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitAssignmentOperator(ctx *ast.AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitLiteral(ctx *ast.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTemplateStringLiteral(ctx *ast.TemplateStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitTemplateStringAtom(ctx *ast.TemplateStringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitNumericLiteral(ctx *ast.NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIdentifierName(ctx *ast.IdentifierNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitIdentifierOrKeyWord(ctx *ast.IdentifierOrKeyWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitReservedWord(ctx *ast.ReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitKeyword(ctx *ast.KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitGetter(ctx *ast.GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitSetter(ctx *ast.SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeScriptVisitor) VisitEos(ctx *ast.EosContext) interface{} {
	return v.VisitChildren(ctx)
}
