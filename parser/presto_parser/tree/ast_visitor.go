package tree

import (
	expr "hitomidb/parser/presto_parser/tree/expression"
	"hitomidb/parser/presto_parser/tree/relation"
	"hitomidb/parser/presto_parser/tree/relation/query_body"
	"hitomidb/parser/presto_parser/tree/select_item"
	"hitomidb/parser/presto_parser/tree/table_element"
)

type AstVisitor interface {
	VisitSelect(stmt *SelectStmt) interface{}

	VisitTableElement(tableElement *table_element.TableElement) interface{}
	VisitSingleColumn(singleColumn *select_item.SingleColumn) interface{}

	VisitRelation(relation *relation.Relation) interface{}
	VisitQueryBody(queryBody *query_body.QueryBody) interface{}
	VisitQueryTable(queryTable *query_body.QueryTable) interface{}
	VisitQuerySpecification(querySpecification *query_body.QuerySpecification) interface{}

	VisitExpression(expression *expr.Expression) interface{}
	VisitIdentifier(identifier *expr.Identifier) interface{}
}
