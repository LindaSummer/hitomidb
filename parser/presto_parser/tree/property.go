package tree

import (
	"context"
)

type IProperty interface {
	Node
	Name() IIdentifier
	Value() IExpression
}

type Property struct {
	*BaseNode
	name  IIdentifier
	value IExpression
}

func (p *Property) Name() IIdentifier {
	return p.name
}

func (p *Property) Value() IExpression {
	return p.value
}

func NewProperty(name IIdentifier, value IExpression, location ...*NodeLocation) *Property {
	return &Property{
		// TODO fill context
		BaseNode: NewBaseNode(context.TODO()),
		name:     name,
		value:    value,
	}
}

func (p *Property) Children() []Node {
	return []Node{p.name, p.value}
}

func (p *Property) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitProperty(p)
}
