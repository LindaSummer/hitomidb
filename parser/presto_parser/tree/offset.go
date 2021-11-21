package tree

import "context"

type IOffset interface {
	Node
	RwoCount() string
}

type Offset struct {
	*BaseNode
	rwoCount string
}

func (o *Offset) RwoCount() string {
	return o.rwoCount
}

func NewOffset(rwoCount string, location ...*NodeLocation) *Offset {
	return &Offset{
		// TODO add context
		BaseNode: NewBaseNode(context.TODO(), location...),
		rwoCount: rwoCount,
	}
}

func (o *Offset) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitOffset(o)
}

func (o *Offset) Children() []Node {
	return EmptyChildren
}
