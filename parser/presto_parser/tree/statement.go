package tree

import (
	"context"
)

type IStatement interface {
	Node
}

type Statement struct {
	*BaseNode
}

func NewStatement(location ...*NodeLocation) *Statement {
	// TODO add context
	return &Statement{BaseNode: NewBaseNode(context.TODO(), location...)}
}

func (s *Statement) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitStatement(s)
}
