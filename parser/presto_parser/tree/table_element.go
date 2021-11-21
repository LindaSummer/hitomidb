package tree

import (
	"context"
)

type ITableElement interface {
	Node
}

type TableElement struct {
	*BaseNode
}

func NewTableElement(location ...*NodeLocation) *TableElement {
	// TODO add context
	return &TableElement{BaseNode: NewBaseNode(context.TODO(), location...)}
}

func (t *TableElement) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitTableElement(t)
}
