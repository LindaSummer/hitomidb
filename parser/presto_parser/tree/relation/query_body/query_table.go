package query_body

import "hitomidb/parser/presto_parser/tree"

type QueryTable struct {
	*QueryBody
	name tree.QualifiedName
}

func (q *QueryTable) Name() tree.QualifiedName {
	return q.name
}

func NewQueryTable(name tree.QualifiedName, location ...*tree.NodeLocation) *QueryTable {
	return &QueryTable{
		QueryBody: NewQueryBody(location...),
		name:      name,
	}
}

func (q *QueryTable) Children() []tree.Node {
	return tree.EmptyChildren
}

func (q *QueryTable) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitQueryTable(q)
}
