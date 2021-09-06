package ast // TypeScriptParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TypeScriptParser.
type TypeScriptParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TypeScriptParser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#bindingPattern.
	VisitBindingPattern(ctx *BindingPatternContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeParameterList.
	VisitTypeParameterList(ctx *TypeParameterListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#constraint.
	VisitConstraint(ctx *ConstraintContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeArgumentList.
	VisitTypeArgumentList(ctx *TypeArgumentListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by TypeScriptParser#Intersection.
	VisitIntersection(ctx *IntersectionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#Primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#Union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#RedefinitionOfType.
	VisitRedefinitionOfType(ctx *RedefinitionOfTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PredefinedPrimType.
	VisitPredefinedPrimType(ctx *PredefinedPrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ArrayPrimType.
	VisitArrayPrimType(ctx *ArrayPrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ParenthesizedPrimType.
	VisitParenthesizedPrimType(ctx *ParenthesizedPrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ThisPrimType.
	VisitThisPrimType(ctx *ThisPrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#TuplePrimType.
	VisitTuplePrimType(ctx *TuplePrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ObjectPrimType.
	VisitObjectPrimType(ctx *ObjectPrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ReferencePrimType.
	VisitReferencePrimType(ctx *ReferencePrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#QueryPrimType.
	VisitQueryPrimType(ctx *QueryPrimTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#predefinedType.
	VisitPredefinedType(ctx *PredefinedTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeReference.
	VisitTypeReference(ctx *TypeReferenceContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#nestedTypeGeneric.
	VisitNestedTypeGeneric(ctx *NestedTypeGenericContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeGeneric.
	VisitTypeGeneric(ctx *TypeGenericContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeIncludeGeneric.
	VisitTypeIncludeGeneric(ctx *TypeIncludeGenericContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#objectType.
	VisitObjectType(ctx *ObjectTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeBody.
	VisitTypeBody(ctx *TypeBodyContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeMemberList.
	VisitTypeMemberList(ctx *TypeMemberListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeMember.
	VisitTypeMember(ctx *TypeMemberContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#tupleElementTypes.
	VisitTupleElementTypes(ctx *TupleElementTypesContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#constructorType.
	VisitConstructorType(ctx *ConstructorTypeContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeQuery.
	VisitTypeQuery(ctx *TypeQueryContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeQueryExpression.
	VisitTypeQueryExpression(ctx *TypeQueryExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#propertySignatur.
	VisitPropertySignatur(ctx *PropertySignaturContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#callSignature.
	VisitCallSignature(ctx *CallSignatureContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#requiredParameterList.
	VisitRequiredParameterList(ctx *RequiredParameterListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#optionalParameter.
	VisitOptionalParameter(ctx *OptionalParameterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#restParameter.
	VisitRestParameter(ctx *RestParameterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#requiredParameter.
	VisitRequiredParameter(ctx *RequiredParameterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#accessibilityModifier.
	VisitAccessibilityModifier(ctx *AccessibilityModifierContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#identifierOrPattern.
	VisitIdentifierOrPattern(ctx *IdentifierOrPatternContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#constructSignature.
	VisitConstructSignature(ctx *ConstructSignatureContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#indexSignature.
	VisitIndexSignature(ctx *IndexSignatureContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#methodSignature.
	VisitMethodSignature(ctx *MethodSignatureContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#typeAliasDeclaration.
	VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#interfaceExtendsClause.
	VisitInterfaceExtendsClause(ctx *InterfaceExtendsClauseContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#classOrInterfaceTypeList.
	VisitClassOrInterfaceTypeList(ctx *ClassOrInterfaceTypeListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#enumMemberList.
	VisitEnumMemberList(ctx *EnumMemberListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#enumMember.
	VisitEnumMember(ctx *EnumMemberContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#namespaceDeclaration.
	VisitNamespaceDeclaration(ctx *NamespaceDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#namespaceName.
	VisitNamespaceName(ctx *NamespaceNameContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#importAliasDeclaration.
	VisitImportAliasDeclaration(ctx *ImportAliasDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#decoratorList.
	VisitDecoratorList(ctx *DecoratorListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#decorator.
	VisitDecorator(ctx *DecoratorContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#decoratorMemberExpression.
	VisitDecoratorMemberExpression(ctx *DecoratorMemberExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#decoratorCallExpression.
	VisitDecoratorCallExpression(ctx *DecoratorCallExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#abstractDeclaration.
	VisitAbstractDeclaration(ctx *AbstractDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#fromBlock.
	VisitFromBlock(ctx *FromBlockContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#multipleImportStatement.
	VisitMultipleImportStatement(ctx *MultipleImportStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#DoStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#WhileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ForStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ForVarStatement.
	VisitForVarStatement(ctx *ForVarStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ForInStatement.
	VisitForInStatement(ctx *ForInStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ForVarInStatement.
	VisitForVarInStatement(ctx *ForVarInStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#varModifier.
	VisitVarModifier(ctx *VarModifierContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#yieldStatement.
	VisitYieldStatement(ctx *YieldStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#withStatement.
	VisitWithStatement(ctx *WithStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#caseClauses.
	VisitCaseClauses(ctx *CaseClausesContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#defaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#labelledStatement.
	VisitLabelledStatement(ctx *LabelledStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#catchProduction.
	VisitCatchProduction(ctx *CatchProductionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#finallyProduction.
	VisitFinallyProduction(ctx *FinallyProductionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#debuggerStatement.
	VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#classHeritage.
	VisitClassHeritage(ctx *ClassHeritageContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#classTail.
	VisitClassTail(ctx *ClassTailContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#classExtendsClause.
	VisitClassExtendsClause(ctx *ClassExtendsClauseContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#implementsClause.
	VisitImplementsClause(ctx *ImplementsClauseContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#classElement.
	VisitClassElement(ctx *ClassElementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PropertyDeclarationExpression.
	VisitPropertyDeclarationExpression(ctx *PropertyDeclarationExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#MethodDeclarationExpression.
	VisitMethodDeclarationExpression(ctx *MethodDeclarationExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#GetterSetterDeclarationExpression.
	VisitGetterSetterDeclarationExpression(ctx *GetterSetterDeclarationExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#AbstractMemberDeclaration.
	VisitAbstractMemberDeclaration(ctx *AbstractMemberDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#propertyMemberBase.
	VisitPropertyMemberBase(ctx *PropertyMemberBaseContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#indexMemberDeclaration.
	VisitIndexMemberDeclaration(ctx *IndexMemberDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#generatorMethod.
	VisitGeneratorMethod(ctx *GeneratorMethodContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#generatorFunctionDeclaration.
	VisitGeneratorFunctionDeclaration(ctx *GeneratorFunctionDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#generatorBlock.
	VisitGeneratorBlock(ctx *GeneratorBlockContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#generatorDefinition.
	VisitGeneratorDefinition(ctx *GeneratorDefinitionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#iteratorBlock.
	VisitIteratorBlock(ctx *IteratorBlockContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#iteratorDefinition.
	VisitIteratorDefinition(ctx *IteratorDefinitionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#formalParameterArg.
	VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#lastFormalParameterArg.
	VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#arrayElement.
	VisitArrayElement(ctx *ArrayElementContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PropertyExpressionAssignment.
	VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ComputedPropertyExpressionAssignment.
	VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PropertyGetter.
	VisitPropertyGetter(ctx *PropertyGetterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PropertySetter.
	VisitPropertySetter(ctx *PropertySetterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#MethodProperty.
	VisitMethodProperty(ctx *MethodPropertyContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PropertyShorthand.
	VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#RestParameterInObject.
	VisitRestParameterInObject(ctx *RestParameterInObjectContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#getAccessor.
	VisitGetAccessor(ctx *GetAccessorContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#setAccessor.
	VisitSetAccessor(ctx *SetAccessorContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#expressionSequence.
	VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#functionExpressionDeclaration.
	VisitFunctionExpressionDeclaration(ctx *FunctionExpressionDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#TemplateStringExpression.
	VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#TernaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#LogicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#GeneratorsExpression.
	VisitGeneratorsExpression(ctx *GeneratorsExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PreIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ObjectLiteralExpression.
	VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#InExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#LogicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#GenericTypes.
	VisitGenericTypes(ctx *GenericTypesContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#NotExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PreDecreaseExpression.
	VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ArgumentsExpression.
	VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ThisExpression.
	VisitThisExpression(ctx *ThisExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#FunctionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#UnaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#AssignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PostDecreaseExpression.
	VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#TypeofExpression.
	VisitTypeofExpression(ctx *TypeofExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#InstanceofExpression.
	VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#UnaryPlusExpression.
	VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#DeleteExpression.
	VisitDeleteExpression(ctx *DeleteExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#GeneratorsFunctionExpression.
	VisitGeneratorsFunctionExpression(ctx *GeneratorsFunctionExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ArrowFunctionExpression.
	VisitArrowFunctionExpression(ctx *ArrowFunctionExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#IteratorsExpression.
	VisitIteratorsExpression(ctx *IteratorsExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#EqualityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#BitXOrExpression.
	VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#CastAsExpression.
	VisitCastAsExpression(ctx *CastAsExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#SuperExpression.
	VisitSuperExpression(ctx *SuperExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#MultiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#BitShiftExpression.
	VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#AdditiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#RelationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#PostIncrementExpression.
	VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#YieldExpression.
	VisitYieldExpression(ctx *YieldExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#BitNotExpression.
	VisitBitNotExpression(ctx *BitNotExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#NewExpression.
	VisitNewExpression(ctx *NewExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ArrayLiteralExpression.
	VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#MemberDotExpression.
	VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#ClassExpression.
	VisitClassExpression(ctx *ClassExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#MemberIndexExpression.
	VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#BitAndExpression.
	VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#BitOrExpression.
	VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#AssignmentOperatorExpression.
	VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#VoidExpression.
	VisitVoidExpression(ctx *VoidExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#asExpression.
	VisitAsExpression(ctx *AsExpressionContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#arrowFunctionDeclaration.
	VisitArrowFunctionDeclaration(ctx *ArrowFunctionDeclarationContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#arrowFunctionParameters.
	VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#arrowFunctionBody.
	VisitArrowFunctionBody(ctx *ArrowFunctionBodyContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#templateStringLiteral.
	VisitTemplateStringLiteral(ctx *TemplateStringLiteralContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#templateStringAtom.
	VisitTemplateStringAtom(ctx *TemplateStringAtomContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#identifierName.
	VisitIdentifierName(ctx *IdentifierNameContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#identifierOrKeyWord.
	VisitIdentifierOrKeyWord(ctx *IdentifierOrKeyWordContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#getter.
	VisitGetter(ctx *GetterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#setter.
	VisitSetter(ctx *SetterContext) interface{}

	// Visit a parse tree produced by TypeScriptParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
