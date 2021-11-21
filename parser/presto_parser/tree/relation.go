package tree

import (
	"context"
)

type IRelation interface {
	Node
}

type Relation struct {
	*BaseNode
}

func NewRelation(location ...*NodeLocation) *Relation {
	// TODO add context
	return &Relation{BaseNode: NewBaseNode(context.TODO(), location...)}
}

func (r *Relation) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitRelation(r)
}
