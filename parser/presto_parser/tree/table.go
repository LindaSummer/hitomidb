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
