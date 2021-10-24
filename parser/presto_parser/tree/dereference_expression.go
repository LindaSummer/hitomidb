package tree

type IDereferenceExpression interface {
	IExpression
	Base() IExpression
	Field() IIdentifier
}

type DereferenceExpression struct {
	*Expression
	base  IExpression
	field IIdentifier
}

func (d *DereferenceExpression) Base() IExpression {
	return d.base
}

func (d *DereferenceExpression) Field() IIdentifier {
	return d.field
}

func NewDereferenceExpression(base IExpression, field IIdentifier, location ...*NodeLocation) *DereferenceExpression {
	return &DereferenceExpression{
		Expression: NewExpression(location...),
		base:       base,
		field:      field,
	}
}

func (d *DereferenceExpression) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitDereferenceExpression(d)
}

func (d *DereferenceExpression) Children() []Node {
	return []Node{d.base}
}
