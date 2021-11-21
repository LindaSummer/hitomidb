package tree

type ITable interface {
	IQueryBody
	Name() *QualifiedName
}

type Table struct {
	*QueryBody
	name *QualifiedName
}

func NewTable(
	name *QualifiedName,
	location ...*NodeLocation,
) *Table {
	return &Table{
		QueryBody: NewQueryBody(location...),
		name:      name,
	}
}

func (t Table) Name() *QualifiedName {
	return t.name
}

func (t *Table) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitTable(t)
}

func (t *Table) Children() []Node {
	return EmptyChildren
}
