package tree

type AstVisitor interface {
	VisitSelect(stmt ISelectStmt) interface{}

	VisitTableElement(tableElement ITableElement) interface{}
	VisitSingleColumn(singleColumn ISingleColumn) interface{}

	VisitStatement(statement IStatement) interface{}
	VisitQuery(query IQuery) interface{}

	VisitRelation(relation IRelation) interface{}
	VisitSampledRelation(relation ISampledRelation) interface{}
	VisitQueryBody(queryBody IQueryBody) interface{}
	VisitSelectItem(item ISelectItem) interface{}
	VisitQueryTable(queryTable IQueryTable) interface{}
	VisitQuerySpecification(querySpecification IQuerySpecification) interface{}
	VisitWith(with IWith) interface{}
	VisitWithQuery(withQuery IWithQuery) interface{}

	VisitOrderBy(orderBy IOrderBy) interface{}
	VisitGroupBy(groupBy IGroupBy) interface{}
	VisitOffset(offset IOffset) interface{}

	VisitTable(table ITable) interface{}
	VisitProperty(property IProperty) interface{}

	VisitAliasRelation(alias IAliasRelation) interface{}

	VisitExpression(expression IExpression) interface{}
	VisitIdentifier(identifier IIdentifier) interface{}
	VisitDereferenceExpression(expression IDereferenceExpression) interface{}
	VisitComparisonExpression(expression IComparisonExpression) interface{}

	VisitSortItem(sortItem ISortItem) interface{}
}
