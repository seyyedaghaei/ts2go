// Code generated from /home/smortezasa/go/src/github.com/seyyedaghaei/ts2go/grammars/TypeScriptParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package ast // TypeScriptParser
import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/seyyedaghaei/ts2go/ast"
)

type BaseTypeScriptParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTypeScriptParserVisitor) VisitInitializer(ctx *ast.InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBindingPattern(ctx *ast.BindingPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeParameters(ctx *ast.TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeParameterList(ctx *ast.TypeParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeParameter(ctx *ast.TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstraint(ctx *ast.ConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeArguments(ctx *ast.TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeArgumentList(ctx *ast.TypeArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeArgument(ctx *ast.TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitType_(ctx *ast.Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIntersection(ctx *ast.IntersectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPrimary(ctx *ast.PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitUnion(ctx *ast.UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRedefinitionOfType(ctx *ast.RedefinitionOfTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPredefinedPrimType(ctx *ast.PredefinedPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayPrimType(ctx *ast.ArrayPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParenthesizedPrimType(ctx *ast.ParenthesizedPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitThisPrimType(ctx *ast.ThisPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTuplePrimType(ctx *ast.TuplePrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectPrimType(ctx *ast.ObjectPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitReferencePrimType(ctx *ast.ReferencePrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitQueryPrimType(ctx *ast.QueryPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPredefinedType(ctx *ast.PredefinedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeReference(ctx *ast.TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNestedTypeGeneric(ctx *ast.NestedTypeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeGeneric(ctx *ast.TypeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeIncludeGeneric(ctx *ast.TypeIncludeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeName(ctx *ast.TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectType(ctx *ast.ObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeBody(ctx *ast.TypeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeMemberList(ctx *ast.TypeMemberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeMember(ctx *ast.TypeMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayType(ctx *ast.ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTupleType(ctx *ast.TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTupleElementTypes(ctx *ast.TupleElementTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionType(ctx *ast.FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstructorType(ctx *ast.ConstructorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeQuery(ctx *ast.TypeQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeQueryExpression(ctx *ast.TypeQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertySignatur(ctx *ast.PropertySignaturContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeAnnotation(ctx *ast.TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCallSignature(ctx *ast.CallSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParameterList(ctx *ast.ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRequiredParameterList(ctx *ast.RequiredParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParameter(ctx *ast.ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitOptionalParameter(ctx *ast.OptionalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRestParameter(ctx *ast.RestParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRequiredParameter(ctx *ast.RequiredParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAccessibilityModifier(ctx *ast.AccessibilityModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierOrPattern(ctx *ast.IdentifierOrPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstructSignature(ctx *ast.ConstructSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIndexSignature(ctx *ast.IndexSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMethodSignature(ctx *ast.MethodSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeAliasDeclaration(ctx *ast.TypeAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstructorDeclaration(ctx *ast.ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInterfaceDeclaration(ctx *ast.InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInterfaceExtendsClause(ctx *ast.InterfaceExtendsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassOrInterfaceTypeList(ctx *ast.ClassOrInterfaceTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumDeclaration(ctx *ast.EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumBody(ctx *ast.EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumMemberList(ctx *ast.EnumMemberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumMember(ctx *ast.EnumMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNamespaceDeclaration(ctx *ast.NamespaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNamespaceName(ctx *ast.NamespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitImportAliasDeclaration(ctx *ast.ImportAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecoratorList(ctx *ast.DecoratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecorator(ctx *ast.DecoratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecoratorMemberExpression(ctx *ast.DecoratorMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecoratorCallExpression(ctx *ast.DecoratorCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitProgram(ctx *ast.ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitStatement(ctx *ast.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBlock(ctx *ast.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitStatementList(ctx *ast.StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAbstractDeclaration(ctx *ast.AbstractDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitImportStatement(ctx *ast.ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFromBlock(ctx *ast.FromBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMultipleImportStatement(ctx *ast.MultipleImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitExportStatement(ctx *ast.ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVariableStatement(ctx *ast.VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVariableDeclarationList(ctx *ast.VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVariableDeclaration(ctx *ast.VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEmptyStatement(ctx *ast.EmptyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitExpressionStatement(ctx *ast.ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIfStatement(ctx *ast.IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDoStatement(ctx *ast.DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitWhileStatement(ctx *ast.WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForStatement(ctx *ast.ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForVarStatement(ctx *ast.ForVarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForInStatement(ctx *ast.ForInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForVarInStatement(ctx *ast.ForVarInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVarModifier(ctx *ast.VarModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitContinueStatement(ctx *ast.ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBreakStatement(ctx *ast.BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitReturnStatement(ctx *ast.ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitYieldStatement(ctx *ast.YieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitWithStatement(ctx *ast.WithStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSwitchStatement(ctx *ast.SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCaseBlock(ctx *ast.CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCaseClauses(ctx *ast.CaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCaseClause(ctx *ast.CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDefaultClause(ctx *ast.DefaultClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLabelledStatement(ctx *ast.LabelledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitThrowStatement(ctx *ast.ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTryStatement(ctx *ast.TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCatchProduction(ctx *ast.CatchProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFinallyProduction(ctx *ast.FinallyProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDebuggerStatement(ctx *ast.DebuggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionDeclaration(ctx *ast.FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassDeclaration(ctx *ast.ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassHeritage(ctx *ast.ClassHeritageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassTail(ctx *ast.ClassTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassExtendsClause(ctx *ast.ClassExtendsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitImplementsClause(ctx *ast.ImplementsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassElement(ctx *ast.ClassElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyDeclarationExpression(ctx *ast.PropertyDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMethodDeclarationExpression(ctx *ast.MethodDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGetterSetterDeclarationExpression(ctx *ast.GetterSetterDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAbstractMemberDeclaration(ctx *ast.AbstractMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyMemberBase(ctx *ast.PropertyMemberBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIndexMemberDeclaration(ctx *ast.IndexMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorMethod(ctx *ast.GeneratorMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorFunctionDeclaration(ctx *ast.GeneratorFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorBlock(ctx *ast.GeneratorBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorDefinition(ctx *ast.GeneratorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIteratorBlock(ctx *ast.IteratorBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIteratorDefinition(ctx *ast.IteratorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFormalParameterList(ctx *ast.FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFormalParameterArg(ctx *ast.FormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLastFormalParameterArg(ctx *ast.LastFormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionBody(ctx *ast.FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayLiteral(ctx *ast.ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitElementList(ctx *ast.ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayElement(ctx *ast.ArrayElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectLiteral(ctx *ast.ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyExpressionAssignment(ctx *ast.PropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitComputedPropertyExpressionAssignment(ctx *ast.ComputedPropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyGetter(ctx *ast.PropertyGetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertySetter(ctx *ast.PropertySetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMethodProperty(ctx *ast.MethodPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyShorthand(ctx *ast.PropertyShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRestParameterInObject(ctx *ast.RestParameterInObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGetAccessor(ctx *ast.GetAccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSetAccessor(ctx *ast.SetAccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyName(ctx *ast.PropertyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArguments(ctx *ast.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArgumentList(ctx *ast.ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArgument(ctx *ast.ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitExpressionSequence(ctx *ast.ExpressionSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionExpressionDeclaration(ctx *ast.FunctionExpressionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTemplateStringExpression(ctx *ast.TemplateStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTernaryExpression(ctx *ast.TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLogicalAndExpression(ctx *ast.LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorsExpression(ctx *ast.GeneratorsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPreIncrementExpression(ctx *ast.PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectLiteralExpression(ctx *ast.ObjectLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInExpression(ctx *ast.InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLogicalOrExpression(ctx *ast.LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGenericTypes(ctx *ast.GenericTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNotExpression(ctx *ast.NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPreDecreaseExpression(ctx *ast.PreDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArgumentsExpression(ctx *ast.ArgumentsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitThisExpression(ctx *ast.ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionExpression(ctx *ast.FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitUnaryMinusExpression(ctx *ast.UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAssignmentExpression(ctx *ast.AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPostDecreaseExpression(ctx *ast.PostDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeofExpression(ctx *ast.TypeofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInstanceofExpression(ctx *ast.InstanceofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitUnaryPlusExpression(ctx *ast.UnaryPlusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDeleteExpression(ctx *ast.DeleteExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorsFunctionExpression(ctx *ast.GeneratorsFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionExpression(ctx *ast.ArrowFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIteratorsExpression(ctx *ast.IteratorsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEqualityExpression(ctx *ast.EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitXOrExpression(ctx *ast.BitXOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCastAsExpression(ctx *ast.CastAsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSuperExpression(ctx *ast.SuperExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMultiplicativeExpression(ctx *ast.MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitShiftExpression(ctx *ast.BitShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParenthesizedExpression(ctx *ast.ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAdditiveExpression(ctx *ast.AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRelationalExpression(ctx *ast.RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPostIncrementExpression(ctx *ast.PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitYieldExpression(ctx *ast.YieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitNotExpression(ctx *ast.BitNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNewExpression(ctx *ast.NewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLiteralExpression(ctx *ast.LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayLiteralExpression(ctx *ast.ArrayLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMemberDotExpression(ctx *ast.MemberDotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassExpression(ctx *ast.ClassExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMemberIndexExpression(ctx *ast.MemberIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierExpression(ctx *ast.IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitAndExpression(ctx *ast.BitAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitOrExpression(ctx *ast.BitOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAssignmentOperatorExpression(ctx *ast.AssignmentOperatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVoidExpression(ctx *ast.VoidExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAsExpression(ctx *ast.AsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionDeclaration(ctx *ast.ArrowFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionParameters(ctx *ast.ArrowFunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionBody(ctx *ast.ArrowFunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAssignmentOperator(ctx *ast.AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLiteral(ctx *ast.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTemplateStringLiteral(ctx *ast.TemplateStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTemplateStringAtom(ctx *ast.TemplateStringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNumericLiteral(ctx *ast.NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierName(ctx *ast.IdentifierNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierOrKeyWord(ctx *ast.IdentifierOrKeyWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitReservedWord(ctx *ast.ReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitKeyword(ctx *ast.KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGetter(ctx *ast.GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSetter(ctx *ast.SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEos(ctx *ast.EosContext) interface{} {
	return v.VisitChildren(ctx)
}
