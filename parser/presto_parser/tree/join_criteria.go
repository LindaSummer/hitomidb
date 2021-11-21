package tree

import "context"

type IJoinCriteria interface {
	Node
}

type IJoinOn interface {
	IJoinCriteria
	Expression() IExpression
}

type JoinOn struct {
	*BaseNode
	expression IExpression
}

func NewJoinOn(
	expression IExpression,
	location ...*NodeLocation,
) *JoinOn {
	return &JoinOn{
		BaseNode:   NewBaseNode(context.TODO(), location...),
		expression: expression,
	}
}

func (j *JoinOn) Children() []Node {
	return []Node{j.expression}
}

func (j *JoinOn) Expression() IExpression {
	return j.expression
}

func (j *JoinOn) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitJoinOn(j)
}

type IJoinUsing interface {
	IJoinCriteria
	Columns() []IIdentifier
}

type JoinUsing struct {
	*BaseNode
	columns []IIdentifier
}

func NewJoinUsing(columns []IIdentifier, location ...*NodeLocation) *JoinUsing {
	return &JoinUsing{
		BaseNode: NewBaseNode(context.TODO(), location...),
		columns:  columns,
	}
}

func (j *JoinUsing) Columns() []IIdentifier {
	return j.columns
}

func (j *JoinUsing) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitJoinUsing(j)
}

func (j *JoinUsing) Children() []Node {
	return EmptyChildren
}

type INaturalJoin interface {
	Node
}

type NaturalJoin struct {
	*BaseNode
}

func NewNaturalJoin(location ...*NodeLocation) *NaturalJoin {
	return &NaturalJoin{BaseNode: NewBaseNode(context.TODO(), location...)}
}

func (n *NaturalJoin) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitNaturalJoin(n)
}

func (n *NaturalJoin) Children() []Node {
	return EmptyChildren
}
