package tree

type AstVisitor interface {
	VisitSelect(stmt *SelectStmt) interface{}

	VisitTableElement(tableElement *TableElement) interface{}
	VisitSingleColumn(singleColumn *SingleColumn) interface{}

	VisitStatement(statement *Statement) interface{}
	VisitQuery(query *Query) interface{}

	VisitRelation(relation *Relation) interface{}
	VisitQueryBody(queryBody *QueryBody) interface{}
	VisitQueryTable(queryTable *QueryTable) interface{}
	VisitQuerySpecification(querySpecification *QuerySpecification) interface{}

	VisitExpression(expression *Expression) interface{}
	VisitIdentifier(identifier *Identifier) interface{}
}
