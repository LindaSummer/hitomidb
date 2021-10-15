package tree

import (
	"context"
)

type Expression struct {
	*BaseNode
}

func NewExpression(location ...*NodeLocation) *Expression {
	// TODO fill context
	return &Expression{BaseNode: NewBaseNode(context.TODO(), location...)}
}

func (e *Expression) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitExpression(e)
}
