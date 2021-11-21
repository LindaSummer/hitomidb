package tree

import (
	"context"

	funk "github.com/thoas/go-funk"
)

type IGroupBy interface {
	Node
	Distinct() bool
	Expressions() []IGroupingElement
}

type GroupBy struct {
	*BaseNode
	distinct    bool
	expressions []IGroupingElement
}

func (g *GroupBy) Distinct() bool {
	return g.distinct
}

func (g *GroupBy) Expressions() []IGroupingElement {
	return g.expressions
}

func NewGroupBy(distinct bool, expressions []IGroupingElement, location ...*NodeLocation) *GroupBy {
	return &GroupBy{
		// TODO add context
		BaseNode:    NewBaseNode(context.TODO(), location...),
		distinct:    distinct,
		expressions: expressions,
	}
}

func (g *GroupBy) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitGroupBy(g)
}

func (g *GroupBy) Children() []Node {
	return funk.Map(g.expressions, func(element IGroupingElement) Node {
		return element
	}).([]Node)
}
