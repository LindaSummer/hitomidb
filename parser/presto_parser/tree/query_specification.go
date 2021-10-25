package tree

type IQuerySpecification interface {
	IQueryBody
	SelectStmt() ISelectStmt
	From() IRelation
	Where() IExpression
	Limit() *string
}

type QuerySpecification struct {
	*QueryBody
	selectStmt ISelectStmt
	from       IRelation
	where      IExpression
	groupBy    IGroupBy
	having     IExpression
	orderBy    IOrderBy
	offset     IOffset
	limit      *string
}

func (q *QuerySpecification) SelectStmt() ISelectStmt {
	return q.selectStmt
}

func (q *QuerySpecification) From() IRelation {
	return q.from
}

func (q *QuerySpecification) Where() IExpression {
	return q.where
}

func (q *QuerySpecification) Limit() *string {
	return q.limit
}

func NewQuerySpecification(
	selectStmt ISelectStmt,
	from IRelation,
	where IExpression,
	groupBy IGroupBy,
	having IExpression,
	orderBy IOrderBy,
	limit *string,
	location ...*NodeLocation,
) *QuerySpecification {
	return &QuerySpecification{
		QueryBody:  NewQueryBody(location...),
		selectStmt: selectStmt,
		from:       from,
		where:      where,
		groupBy:    groupBy,
		having:     having,
		orderBy:    orderBy,
		limit:      limit,
	}
}

func (q *QuerySpecification) Children() []Node {
	var children []Node
	children = AppendNodeIfNotNull(children, q.selectStmt)
	children = AppendNodeIfNotNull(children, q.from)
	children = AppendNodeIfNotNull(children, q.where)
	children = AppendNodeIfNotNull(children, q.groupBy)
	children = AppendNodeIfNotNull(children, q.having)
	children = AppendNodeIfNotNull(children, q.orderBy)
	children = AppendNodeIfNotNull(children, q.offset)

	return children
}

func (q *QuerySpecification) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitQuerySpecification(q)
}
