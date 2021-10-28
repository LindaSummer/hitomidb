package parser

import (
	"hitomidb/parser/presto_parser/tree"

	funk "github.com/thoas/go-funk"

	"github.com/pkg/errors"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type AstBuilder struct {
	*BaseSqlBaseVisitor
	parserOptions ParserOptions
}

func (a *AstBuilder) VisitValueExp(ctx *ValueExpContext) interface{} {
	return a.Visit(ctx.ValueExpression())
}

func (a *AstBuilder) VisitBooleanExp(ctx *BooleanExpContext) interface{} {
	return a.Visit(ctx.BooleanExpression())
}

func NewAstBuilder(parserOptions ParserOptions) *AstBuilder {
	return &AstBuilder{
		BaseSqlBaseVisitor: &BaseSqlBaseVisitor{},
		parserOptions:      parserOptions,
	}
}

func (a *AstBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(a)
}

//func (a *AstBuilder) VisitChildren(node antlr.RuleNode) interface{} {
//	res := make([]interface{}, node.GetChildCount())
//	for i := 0; i < node.GetChildCount(); i++ {
//		res[i] = a.Visit(node.GetChild(i).(antlr.ParseTree))
//	}
//	return res
//}
//
//func (a *AstBuilder) VisitTerminal(node antlr.TerminalNode) interface{} {
//	return node.Accept(a)
//}
//
//func (a *AstBuilder) VisitErrorNode(node antlr.ErrorNode) interface{} {
//	return node.Accept(a)
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
	body := a.Visit(ctx.QueryNoWith()).(tree.IQuery)

	var with tree.IWith

	if ctx.With() != nil {
		with = a.Visit(ctx.With()).(tree.IWith)
	}

	return tree.NewQuery(
		body.QueryBody(),
		with,
		body.OrderBy(),
		body.Limit(),
		body.Offset(),
		a.getParserRuleContextLocation(ctx),
	)
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
	queryTerm := a.Visit(ctx.QueryTerm()).(tree.IQueryBody)
	if queryTerm == nil {
		panic(errors.WithStack(errors.New("queryTerm is nil")))
	}
	var limit *string
	if ctx.GetLimit() != nil {
		l := ctx.GetLimit().GetText()
		limit = &l
	}

	var orderBy tree.IOrderBy
	if len(ctx.AllSortItem()) > 0 {
		orderBy = tree.NewOrderBy(
			funk.Map(ctx.AllSortItem(), func(s ISortItemContext) tree.ISortItem {
				return a.Visit(s).(tree.ISortItem)
			}).([]tree.ISortItem),
			a.getParserRuleContextLocation(ctx),
		)
	}

	var offset tree.IOffset
	if ctx.GetOffset() != nil {
		offset = tree.NewOffset(ctx.GetOffset().GetText(), a.getTokenLocation(ctx.GetOffset()))
	}

	return tree.NewQuery(queryTerm, nil, orderBy, limit, offset, a.getParserRuleContextLocation(ctx))
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
	var from tree.IRelation
	selectItems := funk.Map(ctx.AllSelectItem(), func(s ISelectItemContext) tree.ISelectItem {
		return a.Visit(s).(tree.ISelectItem)
	}).([]tree.ISelectItem)

	relations := funk.Map(ctx.AllRelation(), func(r IRelationContext) tree.IRelation {
		return a.Visit(r).(tree.IRelation)
	}).([]tree.IRelation)

	if len(relations) > 0 {
		from = relations[0]
	}

	for i := 1; i < len(relations); i++ {
		from = tree.NewJoin(tree.IMPLICIT, from, relations[i], nil, a.getParserRuleContextLocation(ctx))
	}

	var where tree.IExpression
	var groupBy tree.IGroupBy
	var having tree.IExpression

	if w, ok := a.visitIfExist(ctx.where).(tree.IExpression); ok && w != nil {
		where = w
	}

	if g, ok := a.visitIfExist(ctx.GroupBy()).(tree.IGroupBy); ok && g != nil {
		groupBy = g
	}

	if h, ok := a.visitIfExist(ctx.having).(tree.IExpression); ok && h != nil {
		having = h
	}

	return tree.NewQuerySpecification(
		tree.NewSelectStmt(
			a.isDistinct(ctx.SetQuantifier()),
			selectItems,
			a.getParserRuleContextLocation(ctx),
			a.getTerminalNodeLocation(ctx.SELECT()),
		),
		from,
		where,
		groupBy,
		having,
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
	var identifier tree.IIdentifier
	if ctx.Identifier() != nil {
		identifier = a.Visit(ctx.Identifier()).(tree.IIdentifier)
	}
	expression := a.Visit(ctx.Expression()).(tree.IExpression)
	return tree.NewSingleColumn(identifier, expression, a.getParserRuleContextLocation(ctx))
}

func (a *AstBuilder) VisitSelectAll(ctx *SelectAllContext) interface{} {
	if ctx, ok := ctx.QualifiedName().(*QualifiedNameContext); ok && ctx != nil {
		return tree.NewAllColumns(a.VisitQualifiedName(ctx).(*tree.QualifiedName), a.getParserRuleContextLocation(ctx))
	}

	return tree.NewAllColumns(nil, a.getParserRuleContextLocation(ctx))
}

func (a *AstBuilder) VisitRelationDefault(ctx *RelationDefaultContext) interface{} {
	return a.Visit(ctx.SampledRelation())
}

func (a *AstBuilder) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	left := a.Visit(ctx.GetLeft()).(tree.IRelation)
	var right tree.IRelation

	if ctx.CROSS() != nil {
		right = a.Visit(ctx.GetRight()).(tree.IRelation)
		return tree.NewJoin(tree.CROSS, left, right, nil, a.getParserRuleContextLocation(ctx))
	}

	var criteria tree.IJoinCriteria
	if ctx.NATURAL() != nil {
		right = a.Visit(ctx.GetRight()).(tree.IRelation)
		criteria = tree.NewNaturalJoin(a.getTerminalNodeLocation(ctx.NATURAL()))
	} else {
		right = a.Visit(ctx.GetRightRelation()).(tree.IRelation)
		criteria = a.Visit(ctx.JoinCriteria()).(tree.IJoinCriteria)
	}

	joinType := a.Visit(ctx.JoinType()).(tree.JoinType)
	return tree.NewJoin(joinType, left, right, criteria, a.getParserRuleContextLocation(ctx))
}

func (a *AstBuilder) VisitJoinType(ctx *JoinTypeContext) interface{} {
	if ctx.LEFT() != nil {
		return tree.LEFT
	}

	if ctx.RIGHT() != nil {
		return tree.RIGHT
	}

	if ctx.FULL() != nil {
		return tree.FULL
	}

	return tree.INNER
}

func (a *AstBuilder) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	if ctx.ON() != nil {
		return tree.NewJoinOn(a.Visit(ctx.BooleanExpression()).(tree.IExpression), a.getParserRuleContextLocation(ctx))
	}

	if ctx.USING() != nil {
		identifiers := funk.Map(ctx.AllIdentifier(), func(i IIdentifierContext) tree.IIdentifier {
			return a.Visit(i).(tree.IIdentifier)
		}).([]tree.IIdentifier)
		return tree.NewJoinUsing(identifiers, a.getParserRuleContextLocation(ctx))
	}
	return nil
}

func (a *AstBuilder) VisitSampledRelation(ctx *SampledRelationContext) interface{} {
	relation := a.Visit(ctx.AliasedRelation()).(tree.IRelation)
	if ctx.TABLESAMPLE() == nil {
		return relation
	}

	return tree.NewSampledRelation(
		relation,
		a.Visit(ctx.SampleType()).(tree.SampledRelationType),
		a.Visit(ctx.GetPercentage()).(tree.IExpression),
		a.getParserRuleContextLocation(ctx),
	)
}

func (a *AstBuilder) VisitSampleType(ctx *SampleTypeContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitAliasedRelation(ctx *AliasedRelationContext) interface{} {
	relation := a.Visit(ctx.RelationPrimary()).(tree.IRelation)

	if ctx.Identifier() == nil {
		return relation
	}

	return tree.NewAliasRelation(
		relation,
		a.Visit(ctx.Identifier()).(tree.IIdentifier),
		a.Visit(ctx.ColumnAliases()).([]tree.IIdentifier),
	)
}

func (a *AstBuilder) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return funk.Map(ctx.AllIdentifier(), func(i IIdentifierContext) tree.IIdentifier {
		return a.Visit(i).(tree.IIdentifier)
	}).([]tree.IIdentifier)
}

func (a *AstBuilder) VisitTableName(ctx *TableNameContext) interface{} {
	return tree.NewTable(a.Visit(ctx.QualifiedName()).(*tree.QualifiedName))
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

func (a *AstBuilder) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	panic("implement me")
}

func (a *AstBuilder) VisitComparison(ctx *ComparisonContext) interface{} {
	return tree.NewComparisonExpression(
		tree.ExpressionOperator(a.Visit(ctx.ComparisonOperator()).(string)),
		a.Visit(ctx.GetLeft()).(tree.IExpression),
		a.Visit(ctx.GetRight()).(tree.IExpression),
		a.getParserRuleContextLocation(ctx),
	)
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
	return a.Visit(ctx.PrimaryExpression())
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
	return tree.NewDereferenceExpression(
		a.Visit(ctx.GetBase()).(tree.IExpression),
		a.Visit(ctx.GetFieldName()).(tree.IIdentifier),
		a.getParserRuleContextLocation(ctx),
	)
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
	return a.Visit(ctx.Identifier())
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
	if ctx.EQ() != nil {
		return tree.EQUAL
	}

	if ctx.NEQ() != nil {
		return tree.NOT_EQUAL
	}

	if ctx.LT() != nil {
		return tree.LESS_THAN
	}

	if ctx.LTE() != nil {
		return tree.LESS_THAN_OR_EQUAL
	}

	if ctx.GT() != nil {
		return tree.GREATER_THAN
	}

	if ctx.GTE() != nil {
		return tree.GREATER_THAN_OR_EQUAL
	}

	return nil
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
	return tree.NewQualifiedNameWithOrigin(origin[0], origin[1:]...)
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
	return tree.NewIdentifier(ctx.IDENTIFIER().GetText(), a.getParserRuleContextLocation(ctx))
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

func (a *AstBuilder) isDistinct(ctx ISetQuantifierContext) bool {
	if ctx, ok := ctx.(*SetQuantifierContext); ok && ctx != nil && ctx.DISTINCT() != nil {
		return true
	}
	return false
}

func (a *AstBuilder) getTerminalNodeLocation(node antlr.TerminalNode) *tree.NodeLocation {
	if node == nil {
		return nil
	}

	return a.getTokenLocation(node.GetSymbol())
}

func (a *AstBuilder) getParserRuleContextLocation(ctx antlr.ParserRuleContext) *tree.NodeLocation {
	if ctx == nil {
		return nil
	}
	return a.getTokenLocation(ctx.GetStart())
}

func (a *AstBuilder) getTokenLocation(token antlr.Token) *tree.NodeLocation {
	if token != nil {
		return nil
	}

	return tree.NewNodeLocation(token.GetLine(), token.GetColumn())
}
