package tree

import (
	"context"
)

type ISelectItem interface {
	Node
}

type SelectItem struct {
	*BaseNode
}

func NewSelectItem(location ...*NodeLocation) *SelectItem {
	return &SelectItem{
		// TODO add context
		BaseNode: NewBaseNode(context.TODO(), location...),
	}
}
