package expression

import (
	"context"
	"hitomidb/parser/presto_parser/tree"
)

type Expression struct {
	*tree.BaseNode
}

func NewExpression(location ...*tree.NodeLocation) *Expression {
	// TODO fill context
	return &Expression{BaseNode: tree.NewBaseNode(context.TODO(), location...)}
}

func (e *Expression) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitExpression(e)
}
