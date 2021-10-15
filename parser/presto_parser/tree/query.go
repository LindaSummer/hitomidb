package tree

type Query struct {
	*Statement
	// TODO add with
	queryBody *QueryBody
	// TODO add order by
	limit *string
	// TODO add offset
}

func (q *Query) QueryBody() *QueryBody {
	return q.queryBody
}

func (q *Query) Limit() *string {
	return q.limit
}

// NewQuery TODO fulfill constructor
func NewQuery(queryBody *QueryBody, limit *string, location ...*NodeLocation) *Query {
	return &Query{Statement: NewStatement(location...), queryBody: queryBody, limit: limit}
}

func (q *Query) Children() []Node {
	return []Node{
		q.queryBody,
		// TODO fill rest children
	}
}

func (q *Query) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitQuery(q)
}
