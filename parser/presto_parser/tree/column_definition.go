package tree

type IColumnDefinition interface {
	ITableElement
	Name() IIdentifier
	ColumnType() string
	Nullable() bool
	Properties() []IProperty
	Comment() *string
}

type ColumnDefinition struct {
	*TableElement
	name       IIdentifier
	columnType string
	nullable   bool
	properties []IProperty
	comment    *string
}

func NewColumnDefinition(
	name IIdentifier,
	columnType string,
	nullable bool,
	properties []IProperty,
	comment *string,
	location ...*NodeLocation,
) *ColumnDefinition {
	return &ColumnDefinition{
		TableElement: NewTableElement(location...),
		name:         name,
		columnType:   columnType,
		nullable:     nullable,
		properties:   properties,
		comment:      comment,
	}
}

func (c *ColumnDefinition) Name() IIdentifier {
	return c.name
}

func (c *ColumnDefinition) ColumnType() string {
	return c.columnType
}

func (c *ColumnDefinition) Nullable() bool {
	return c.nullable
}

func (c *ColumnDefinition) Properties() []IProperty {
	return c.properties
}

func (c *ColumnDefinition) Comment() *string {
	return c.comment
}

func (c *ColumnDefinition) Children() []Node {
	return EmptyChildren
}

func (c *ColumnDefinition) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitColumnDefinition(c)
}
