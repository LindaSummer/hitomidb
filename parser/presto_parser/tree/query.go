package tree

type IQuery interface {
	IStatement
	QueryBody() IQueryBody
	Limit() *string
}

type Query struct {
	*Statement
	// TODO add with
	queryBody IQueryBody
	// TODO add order by
	limit *string
	// TODO add offset
}

func (q *Query) QueryBody() IQueryBody {
	return q.queryBody
}

func (q *Query) Limit() *string {
	return q.limit
}

// NewQuery TODO fulfill constructor
func NewQuery(queryBody IQueryBody, limit *string, location ...*NodeLocation) *Query {
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
