package select_item

import (
	"context"
	"hitomidb/parser/presto_parser/tree"
)

type SelectItem struct {
	*tree.BaseNode
}

func NewSelectItem(location ...*tree.NodeLocation) *SelectItem {
	return &SelectItem{
		// TODO add context
		BaseNode: tree.NewBaseNode(context.TODO(), location...),
	}
}
