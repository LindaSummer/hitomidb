package query_body

import (
	"hitomidb/parser/presto_parser/tree"
	expr "hitomidb/parser/presto_parser/tree/expression"
	"hitomidb/parser/presto_parser/tree/relation"
	"hitomidb/parser/presto_parser/util"
)

type QuerySpecification struct {
	*QueryBody
	selectStmt *tree.SelectStmt
	from       *relation.Relation
	where      *expr.Expression
	// TODO add `group by`
	//groupBy	*GroupBy
	//TODO add having
	//having *Having
	// TODO add `order by`
	//orderBy *OrderBy
	//TODO add offset
	//offset *Offset
	limit *string
}

func (q *QuerySpecification) SelectStmt() *tree.SelectStmt {
	return q.selectStmt
}

func (q *QuerySpecification) From() *relation.Relation {
	return q.from
}

func (q *QuerySpecification) Where() *expr.Expression {
	return q.where
}

func (q *QuerySpecification) Limit() *string {
	return q.limit
}

func NewQuerySpecification(
	selectStmt *tree.SelectStmt,
	from *relation.Relation,
	where *expr.Expression,
	limit *string,
	location ...*tree.NodeLocation,
) *QuerySpecification {
	return &QuerySpecification{
		QueryBody:  NewQueryBody(location...),
		selectStmt: selectStmt,
		from:       from,
		where:      where,
		limit:      limit,
	}
}

func (q *QuerySpecification) Children() []tree.Node {
	var children []tree.Node
	// TODO add all children
	children = util.AppendNodeIfNotNull(children, q.selectStmt)
	children = util.AppendNodeIfNotNull(children, q.from)
	children = util.AppendNodeIfNotNull(children, q.where)

	return children
}

func (q *QuerySpecification) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitQuerySpecification(q)
}
