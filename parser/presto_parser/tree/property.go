package tree

import (
	"context"
	expr "hitomidb/parser/presto_parser/tree/expression"
)

type Property struct {
	*BaseNode
	name  *expr.Identifier
	value *expr.Expression
}

func (p *Property) Name() *expr.Identifier {
	return p.name
}

func (p *Property) Value() *expr.Expression {
	return p.value
}

func NewProperty(name *expr.Identifier, value *expr.Expression, location ...*NodeLocation) *Property {
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
