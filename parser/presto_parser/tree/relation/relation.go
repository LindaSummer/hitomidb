package relation

import (
	"context"
	"hitomidb/parser/presto_parser/tree"
)

type Relation struct {
	*tree.BaseNode
}

func NewRelation(location ...*tree.NodeLocation) *Relation {
	// TODO add context
	return &Relation{BaseNode: tree.NewBaseNode(context.TODO(), location...)}
}

func (r *Relation) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitRelation(r)
}
