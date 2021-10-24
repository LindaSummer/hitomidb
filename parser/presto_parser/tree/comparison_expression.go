package tree

type ExpressionOperator string

const (
	EQUAL                 ExpressionOperator = "="
	NOT_EQUAL                                = "<>"
	LESS_THAN                                = "<"
	LESS_THAN_OR_EQUAL                       = "<="
	GREATER_THAN                             = ">"
	GREATER_THAN_OR_EQUAL                    = ">="
	IS_DISTINCT_FROM                         = "IS DISTINCT FROM"
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
