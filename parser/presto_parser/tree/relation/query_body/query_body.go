package query_body

import (
	"hitomidb/parser/presto_parser/tree"
	"hitomidb/parser/presto_parser/tree/relation"
)

type QueryBody struct {
	*relation.Relation
}

func NewQueryBody(location ...*tree.NodeLocation) *QueryBody {
	return &QueryBody{Relation: relation.NewRelation(location...)}
}

func (q *QueryBody) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitQueryBody(q)
}
