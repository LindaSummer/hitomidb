package tree

type ExpressionOperator string

const (
	Equal              ExpressionOperator = "="
	NotEqual           ExpressionOperator = "<>"
	LessThan           ExpressionOperator = "<"
	LessThanOrEqual    ExpressionOperator = "<="
	GreaterThan        ExpressionOperator = ">"
	GreaterThanOrEqual ExpressionOperator = ">="
	IsDistinctFrom     ExpressionOperator = "IS DISTINCT FROM"
)

type IComparisonExpression interface {
	IExpression
	Operator() ExpressionOperator
	Left() IExpression
	Right() IExpression
}

type ComparisonExpression struct {
	*Expression
	operator ExpressionOperator
	left     IExpression
	right    IExpression
}

func NewComparisonExpression(
	operator ExpressionOperator,
	left IExpression,
	right IExpression,
	location ...*NodeLocation,
) *ComparisonExpression {
	return &ComparisonExpression{
		Expression: NewExpression(location...),
		operator:   operator,
		left:       left,
		right:      right,
	}
}

func (c *ComparisonExpression) Operator() ExpressionOperator {
	return c.operator
}

func (c *ComparisonExpression) Left() IExpression {
	return c.left
}

func (c *ComparisonExpression) Right() IExpression {
	return c.right
}

func (c *ComparisonExpression) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitComparisonExpression(c)
}

func (c *ComparisonExpression) Children() []Node {
	return []Node{
		c.left,
		c.right,
	}
}
