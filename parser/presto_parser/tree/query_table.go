package tree

type IQueryTable interface {
	IQueryBody
	Name() QualifiedName
}

type QueryTable struct {
	*QueryBody
	name QualifiedName
}

func (q *QueryTable) Name() QualifiedName {
	return q.name
}

func NewQueryTable(name QualifiedName, location ...*NodeLocation) *QueryTable {
	return &QueryTable{
		QueryBody: NewQueryBody(location...),
		name:      name,
	}
}

func (q *QueryTable) Children() []Node {
	return EmptyChildren
}

func (q *QueryTable) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitQueryTable(q)
}
