package tree

import (
	"context"

	funk "github.com/thoas/go-funk"
)

type IOrderBy interface {
	Node
	SortItems() []ISortItem
}

type OrderBy struct {
	*BaseNode
	sortItems []ISortItem
}

func NewOrderBy(
	sortItems []ISortItem,
	location ...*NodeLocation,
) *OrderBy {
	return &OrderBy{
		BaseNode:  NewBaseNode(context.TODO(), location...),
		sortItems: sortItems,
	}
}

func (o *OrderBy) SortItems() []ISortItem {
	return o.sortItems
}

func (o *OrderBy) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitOrderBy(o)
}

func (o *OrderBy) Children() []Node {
	return funk.Map(o.sortItems, func(s ISortItem) Node { return s }).([]Node)
}
