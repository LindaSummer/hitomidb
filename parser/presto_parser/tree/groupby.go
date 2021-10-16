package tree

import "context"

type GroupBy struct {
	*BaseNode
	distinct    bool
	expressions []*GroupingElement
}

func (g *GroupBy) Distinct() bool {
	return g.distinct
}

func (g *GroupBy) Expressions() []*GroupingElement {
	return g.expressions
}

func NewGroupBy(distinct bool, expressions []*GroupingElement, location ...*NodeLocation) *GroupBy {
	return &GroupBy{
		// TODO add context
		BaseNode:    NewBaseNode(context.TODO(), location...),
		distinct:    distinct,
		expressions: expressions,
	}
}
