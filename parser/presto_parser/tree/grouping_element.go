package tree

type IGroupingElement interface {
	Node
	Expressions() []IExpression
}

type GroupingElement struct {
	*BaseNode
	expressions []IExpression
}

func (g *GroupingElement) Expressions() []IExpression {
	return g.expressions
}
