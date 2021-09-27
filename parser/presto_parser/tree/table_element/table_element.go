package table_element

import (
	"context"
	"hitomidb/parser/presto_parser/tree"
)

type TableElement struct {
	*tree.BaseNode
}

func NewTableElement(location ...*tree.NodeLocation) *TableElement {
	// TODO add context
	return &TableElement{BaseNode: tree.NewBaseNode(context.TODO(), location...)}
}

func (t *TableElement) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitTableElement(t)
}
