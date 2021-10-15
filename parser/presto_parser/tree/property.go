package tree

import (
	"context"
)

type Property struct {
	*BaseNode
	name  *Identifier
	value *Expression
}

func (p *Property) Name() *Identifier {
	return p.name
}

func (p *Property) Value() *Expression {
	return p.value
}

func NewProperty(name *Identifier, value *Expression, location ...*NodeLocation) *Property {
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
