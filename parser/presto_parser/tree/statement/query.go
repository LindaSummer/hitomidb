package statement

import (
	"hitomidb/parser/presto_parser/tree"
	"hitomidb/parser/presto_parser/tree/relation/query_body"
)

type Query struct {
	*Statement
	// TODO add with
	queryBody *query_body.QueryBody
	// TODO add order by
	limit *string
	// TODO add offset
}

func (q *Query) QueryBody() *query_body.QueryBody {
	return q.queryBody
}

func (q *Query) Limit() *string {
	return q.limit
}

// NewQuery TODO fulfill constructor
func NewQuery(queryBody *query_body.QueryBody, limit *string, location ...*tree.NodeLocation) *Query {
	return &Query{Statement: NewStatement(location...), queryBody: queryBody, limit: limit}
}

func (q *Query) Children() []tree.Node {
	return []tree.Node{
		q.queryBody,
		// TODO fill rest children
	}
}

func (q *Query) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitQuery(q)
}
