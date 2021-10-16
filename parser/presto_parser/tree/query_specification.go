package tree

type QuerySpecification struct {
	*QueryBody
	selectStmt *SelectStmt
	from       *Relation
	where      *Expression
	groupBy    *GroupBy
	having     *Expression
	orderBy    *OrderBy
	offset     *Offset
	limit      *string
}

func (q *QuerySpecification) SelectStmt() *SelectStmt {
	return q.selectStmt
}

func (q *QuerySpecification) From() *Relation {
	return q.from
}

func (q *QuerySpecification) Where() *Expression {
	return q.where
}

func (q *QuerySpecification) Limit() *string {
	return q.limit
}

func NewQuerySpecification(
	selectStmt *SelectStmt,
	from *Relation,
	where *Expression,
	groupBy *GroupBy,
	having *Expression,
	orderBy *OrderBy,
	limit *string,
	location ...*NodeLocation,
) *QuerySpecification {
	return &QuerySpecification{
		QueryBody:  NewQueryBody(location...),
		selectStmt: selectStmt,
		from:       from,
		where:      where,
		limit:      limit,
	}
}

func (q *QuerySpecification) Children() []Node {
	var children []Node
	// TODO add all children
	children = AppendNodeIfNotNull(children, q.selectStmt)
	children = AppendNodeIfNotNull(children, q.from)
	children = AppendNodeIfNotNull(children, q.where)

	return children
}

func (q *QuerySpecification) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitQuerySpecification(q)
}
