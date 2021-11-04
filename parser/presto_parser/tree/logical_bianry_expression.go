package tree

type BinaryOperator int

const (
	AndOperator BinaryOperator = iota
	OrOperator
	InvalidOperator
)

type ILogicalBinaryExpression interface {
	IExpression
	Operator() BinaryOperator
	Left() IExpression
	Right() IExpression
}

type LogicalBinaryExpression struct {
	*Expression
	operator    BinaryOperator
	left, right IExpression
}

func NewLogicalBinaryExpression(
	operator BinaryOperator,
	left IExpression,
	right IExpression,
	location ...*NodeLocation,
) *LogicalBinaryExpression {
	return &LogicalBinaryExpression{
		Expression: NewExpression(location...),
		operator:   operator,
		left:       left,
		right:      right,
	}
}

func (l *LogicalBinaryExpression) Operator() BinaryOperator {
	return l.operator
}

func (l *LogicalBinaryExpression) Left() IExpression {
	return l.left
}

func (l *LogicalBinaryExpression) Right() IExpression {
	return l.right
}

func (l *LogicalBinaryExpression) Children() []Node {
	return []Node{l.left, l.right}
}

func (l *LogicalBinaryExpression) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitLogicalBinaryExpression(l)
}
