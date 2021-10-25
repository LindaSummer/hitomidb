package tree

type IQuery interface {
	IStatement
	QueryBody() IQueryBody
	Limit() *string
	With() IWith
	OrderBy() IOrderBy
	Offset() IOffset
}

type Query struct {
	*Statement
	with      IWith
	queryBody IQueryBody
	orderBy   IOrderBy
	limit     *string
	offset    IOffset
}

func (q *Query) OrderBy() IOrderBy {
	return q.orderBy
}

func (q *Query) Offset() IOffset {
	return q.offset
}

func (q *Query) With() IWith {
	return q.with
}

func (q *Query) QueryBody() IQueryBody {
	return q.queryBody
}

func (q *Query) Limit() *string {
	return q.limit
}

func NewQuery(
	queryBody IQueryBody,
	with IWith,
	orderBy IOrderBy,
	limit *string,
	offset IOffset,
	location ...*NodeLocation,
) *Query {
	return &Query{
		Statement: NewStatement(location...),
		with:      with,
		queryBody: queryBody,
		orderBy:   orderBy,
		limit:     limit,
		offset:    offset,
	}
}

func (q *Query) Children() []Node {
	var children = make([]Node, 0)
	if q.with != nil {
		children = append(children, q.with)
	}

	children = append(children, q.queryBody)

	if q.orderBy != nil {
		children = append(children, q.orderBy)
	}

	if q.offset != nil {
		children = append(children, q.offset)
	}

	return children
}

func (q *Query) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitQuery(q)
}
