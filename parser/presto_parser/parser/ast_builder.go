package parser

import (
	stmt "hitomidb/parser/presto_parser/tree"
	"reflect"

	funk "github.com/thoas/go-funk"

	"github.com/pkg/errors"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type AstBuilder struct {
	*BaseSqlBaseVisitor
	parserOptions ParserOptions
}

func NewAstBuilder(parserOptions ParserOptions) *AstBuilder {
	return &AstBuilder{
		BaseSqlBaseVisitor: &BaseSqlBaseVisitor{},
		parserOptions:      parserOptions,
	}
}

func (a *AstBuilder) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *SingleStatementContext:
		return a.VisitSingleStatement(ctx)
	case *StandaloneExpressionContext:
		return a.VisitStandaloneExpression(ctx)
	case *StandaloneRoutineBodyContext:
		return a.VisitStandaloneRoutineBody(ctx)
	case *StatementDefaultContext:
		return a.VisitStatementDefault(ctx)
	case *QueryContext:
		return a.VisitQuery(ctx)
	case *QueryNoWithContext:
		return a.VisitQueryNoWith(ctx)
	case *QueryTermDefaultContext:
		return a.VisitQueryTermDefault(ctx)
	case *QueryPrimaryDefaultContext:
		return a.VisitQueryPrimaryDefault(ctx)
	case *QuerySpecificationContext:
		return a.VisitQuerySpecification(ctx)
	case *SelectAllContext:
		return a.VisitSelectAll(ctx)
	default:
		panic(errors.WithStack(errors.Errorf("not implemented type: %v", reflect.TypeOf(ctx))))
	}
}

//func (a *AstBuilder) VisitChildren(node antlr.RuleNode) interface{} {
//	panic("implement me")
//}
//
//func (a *AstBuilder) VisitTerminal(node antlr.TerminalNode) interface{} {
//	panic("implement me")
//}
//
//func (a *AstBuilder) VisitErrorNode(node antlr.ErrorNode) interface{} {
//	panic("implement me")
//}

func (a *AstBuilder) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return a.Visit(ctx.Statement())
}

func (a *AstBuilder) VisitStandaloneExpression(ctx *StandaloneExpressionContext) interface{} {
	return a.Visit(ctx.Expression())
}

func (a *AstBuilder) VisitStandaloneRoutineBody(ctx *StandaloneRoutineBodyContext) interface{} {
	return a.Visit(ctx.RoutineBody())
}

func (a *AstBuilder) VisitStatementDefault(ctx *StatementDefaultContext) interface{} {
	return a.Visit(ctx.Query())
}

func (a *AstBuilder) VisitUse(ctx *UseContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateSchema(ctx *CreateSchemaContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDropSchema(ctx *DropSchemaContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRenameSchema(ctx *RenameSchemaContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateTableAsSelect(ctx *CreateTableAsSelectContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateTable(ctx *CreateTableContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDropTable(ctx *DropTableContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitInsertInto(ctx *InsertIntoContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDelete(ctx *DeleteContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRenameTable(ctx *RenameTableContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRenameColumn(ctx *RenameColumnContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDropColumn(ctx *DropColumnContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAddColumn(ctx *AddColumnContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAnalyze(ctx *AnalyzeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateType(ctx *CreateTypeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateView(ctx *CreateViewContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDropView(ctx *DropViewContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateMaterializedView(ctx *CreateMaterializedViewContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDropMaterializedView(ctx *DropMaterializedViewContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRefreshMaterializedView(ctx *RefreshMaterializedViewContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateFunction(ctx *CreateFunctionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAlterFunction(ctx *AlterFunctionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDropFunction(ctx *DropFunctionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCall(ctx *CallContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCreateRole(ctx *CreateRoleContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDropRole(ctx *DropRoleContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitGrantRoles(ctx *GrantRolesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRevokeRoles(ctx *RevokeRolesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSetRole(ctx *SetRoleContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitGrant(ctx *GrantContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRevoke(ctx *RevokeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowGrants(ctx *ShowGrantsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExplain(ctx *ExplainContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowCreateTable(ctx *ShowCreateTableContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowCreateView(ctx *ShowCreateViewContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowCreateMaterializedView(ctx *ShowCreateMaterializedViewContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowCreateFunction(ctx *ShowCreateFunctionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowTables(ctx *ShowTablesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowSchemas(ctx *ShowSchemasContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowCatalogs(ctx *ShowCatalogsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowColumns(ctx *ShowColumnsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowStats(ctx *ShowStatsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowStatsForQuery(ctx *ShowStatsForQueryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowRoles(ctx *ShowRolesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowRoleGrants(ctx *ShowRoleGrantsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowFunctions(ctx *ShowFunctionsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitShowSession(ctx *ShowSessionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSetSession(ctx *SetSessionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitResetSession(ctx *ResetSessionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitStartTransaction(ctx *StartTransactionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCommit(ctx *CommitContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRollback(ctx *RollbackContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitPrepare(ctx *PrepareContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDeallocate(ctx *DeallocateContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExecute(ctx *ExecuteContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDescribeInput(ctx *DescribeInputContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDescribeOutput(ctx *DescribeOutputContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitQuery(ctx *QueryContext) interface{} {
	body := a.Visit(ctx.QueryNoWith()).(*stmt.Query)

	// TODO fulfill with
	//if ctx.With() != nil {
	//	a.Visit(ctx.With())
	//}

	return stmt.NewQuery(body.QueryBody(), body.Limit())
}

func (a *AstBuilder) VisitWith(ctx *WithContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTableElement(ctx *TableElementContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLikeClause(ctx *LikeClauseContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitProperties(ctx *PropertiesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitProperty(ctx *PropertyContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSqlParameterDeclaration(ctx *SqlParameterDeclarationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRoutineCharacteristics(ctx *RoutineCharacteristicsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRoutineCharacteristic(ctx *RoutineCharacteristicContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAlterRoutineCharacteristics(ctx *AlterRoutineCharacteristicsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAlterRoutineCharacteristic(ctx *AlterRoutineCharacteristicContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRoutineBody(ctx *RoutineBodyContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExternalBodyReference(ctx *ExternalBodyReferenceContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLanguage(ctx *LanguageContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDeterminism(ctx *DeterminismContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNullCallClause(ctx *NullCallClauseContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExternalRoutineName(ctx *ExternalRoutineNameContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitQueryNoWith(ctx *QueryNoWithContext) interface{} {
	queryTerm := a.Visit(ctx.QueryTerm()).(*stmt.QueryBody)
	if queryTerm == nil {
		panic(errors.WithStack(errors.New("queryTerm is nil")))
	}
	limit := ctx.limit.GetText()
	// TODO implement rest
	return stmt.NewQuery(queryTerm, &limit)
}

func (a *AstBuilder) VisitQueryTermDefault(ctx *QueryTermDefaultContext) interface{} {
	return a.Visit(ctx.QueryPrimary())
}

func (a *AstBuilder) VisitSetOperation(ctx *SetOperationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return a.Visit(ctx.QuerySpecification())
}

func (a *AstBuilder) VisitTable(ctx *TableContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitInlineTable(ctx *InlineTableContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSubquery(ctx *SubqueryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSortItem(ctx *SortItemContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	var from *stmt.Relation
	selectItems := a.visitList(ctx.AllSelectItem()).([]*stmt.SelectItem)

	relations := a.visitList(ctx.AllRelation()).([]*stmt.Relation)
	if len(relations) > 0 {
		// TODO finish JOIN
		from = relations[0]
	}

	return stmt.NewQuerySpecification(
		stmt.NewSelectStmt(
			a.isDistinct(ctx.SetQuantifier().(*SetQuantifierContext)),
			selectItems,
			a.getParserRuleContextLocation(ctx),
			a.getTerminalNodeLocation(ctx.SELECT()),
		),
		from,
		a.visitIfExist(ctx.where).(*stmt.Expression),
		a.visitIfExist(ctx.GroupBy()).(*stmt.GroupBy),
		a.visitIfExist(ctx.having).(*stmt.Expression),
		nil,
		nil,
	)
}

func (a *AstBuilder) VisitGroupBy(ctx *GroupByContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRollup(ctx *RollupContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCube(ctx *CubeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitGroupingSet(ctx *GroupingSetContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNamedQuery(ctx *NamedQueryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSelectSingle(ctx *SelectSingleContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSelectAll(ctx *SelectAllContext) interface{} {
	if ctx := ctx.QualifiedName().(*QualifiedNameContext); ctx != nil {
		return stmt.NewAllColumns(a.VisitQualifiedName(ctx).(*stmt.QualifiedName), a.getParserRuleContextLocation(ctx))
	}

	return stmt.NewAllColumns(nil, a.getParserRuleContextLocation(ctx))
}

func (a *AstBuilder) VisitRelationDefault(ctx *RelationDefaultContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitJoinType(ctx *JoinTypeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSampledRelation(ctx *SampledRelationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSampleType(ctx *SampleTypeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAliasedRelation(ctx *AliasedRelationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTableName(ctx *TableNameContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSubqueryRelation(ctx *SubqueryRelationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitUnnest(ctx *UnnestContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLateral(ctx *LateralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExpression(ctx *ExpressionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitPredicated(ctx *PredicatedContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitComparison(ctx *ComparisonContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitQuantifiedComparison(ctx *QuantifiedComparisonContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBetween(ctx *BetweenContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitInList(ctx *InListContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitInSubquery(ctx *InSubqueryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLike(ctx *LikeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNullPredicate(ctx *NullPredicateContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDistinctFrom(ctx *DistinctFromContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitConcatenation(ctx *ConcatenationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAtTimeZone(ctx *AtTimeZoneContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDereference(ctx *DereferenceContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTypeConstructor(ctx *TypeConstructorContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSpecialDateTimeFunction(ctx *SpecialDateTimeFunctionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSubstring(ctx *SubstringContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCast(ctx *CastContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLambda(ctx *LambdaContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitParameter(ctx *ParameterContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNormalize(ctx *NormalizeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRowConstructor(ctx *RowConstructorContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSubscript(ctx *SubscriptContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCurrentUser(ctx *CurrentUserContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExtract(ctx *ExtractContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExists(ctx *ExistsContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitPosition(ctx *PositionContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBasicStringLiteral(ctx *BasicStringLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNullTreatment(ctx *NullTreatmentContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTimeZoneInterval(ctx *TimeZoneIntervalContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTimeZoneString(ctx *TimeZoneStringContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitComparisonQuantifier(ctx *ComparisonQuantifierContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitInterval(ctx *IntervalContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitIntervalField(ctx *IntervalFieldContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNormalForm(ctx *NormalFormContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTypes(ctx *TypesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitType_(ctx *Type_Context) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBaseType(ctx *BaseTypeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitFilter(ctx *FilterContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitOver(ctx *OverContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBoundedFrame(ctx *BoundedFrameContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExplainFormat(ctx *ExplainFormatContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitExplainType(ctx *ExplainTypeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitIsolationLevel(ctx *IsolationLevelContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitReadUncommitted(ctx *ReadUncommittedContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitReadCommitted(ctx *ReadCommittedContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRepeatableRead(ctx *RepeatableReadContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSerializable(ctx *SerializableContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitPositionalArgument(ctx *PositionalArgumentContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNamedArgument(ctx *NamedArgumentContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	origin := funk.Chain(ctx.AllIdentifier()).Map(func(ctx IIdentifierContext) string {
		return ctx.GetText()
	}).Value().([]string)
	return stmt.NewQualifiedNameWithOrigin(origin[0], origin[1:]...)
}

func (a *AstBuilder) VisitCurrentUserGrantor(ctx *CurrentUserGrantorContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitCurrentRoleGrantor(ctx *CurrentRoleGrantorContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitSpecifiedPrincipal(ctx *SpecifiedPrincipalContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitUserPrincipal(ctx *UserPrincipalContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRolePrincipal(ctx *RolePrincipalContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitUnspecifiedPrincipal(ctx *UnspecifiedPrincipalContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitRoles(ctx *RolesContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitDoubleLiteral(ctx *DoubleLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitNonReserved(ctx *NonReservedContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) visitList(list interface{}) interface{} {
	return funk.Map(list, func(tree antlr.ParseTree) interface{} {
		return a.Visit(tree)
	})
}

func (a *AstBuilder) visitIfExist(tree antlr.ParseTree) interface{} {
	if tree != nil {
		return a.Visit(tree)
	}
	return nil
}

func (a *AstBuilder) isDistinct(ctx *SetQuantifierContext) bool {
	return ctx != nil && ctx.DISTINCT() != nil
}

func (a *AstBuilder) getTerminalNodeLocation(node antlr.TerminalNode) *stmt.NodeLocation {
	if node == nil {
		return nil
	}

	return a.getTokenLocation(node.GetSymbol())
}

func (a *AstBuilder) getParserRuleContextLocation(ctx antlr.ParserRuleContext) *stmt.NodeLocation {
	if ctx == nil {
		return nil
	}
	return a.getTokenLocation(ctx.GetStart())
}

func (a *AstBuilder) getTokenLocation(token antlr.Token) *stmt.NodeLocation {
	if token != nil {
		return nil
	}

	return stmt.NewNodeLocation(token.GetLine(), token.GetColumn())
}
