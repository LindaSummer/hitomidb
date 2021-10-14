package statement

import (
	"context"
	"hitomidb/parser/presto_parser/tree"
)

type Statement struct {
	*tree.BaseNode
}

func NewStatement(location ...*tree.NodeLocation) *Statement {
	// TODO add context
	return &Statement{BaseNode: tree.NewBaseNode(context.TODO(), location...)}
}

func (s *Statement) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitStatement(s)
}
