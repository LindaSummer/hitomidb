package tree

type ISingleColumn interface {
	ISelectItem
	Alias() IIdentifier
	Expression() IExpression
}

type SingleColumn struct {
	*SelectItem
	alias      IIdentifier
	expression IExpression
}

func (s *SingleColumn) Alias() IIdentifier {
	return s.alias
}

func (s *SingleColumn) Expression() IExpression {
	return s.expression
}

func NewSingleColumn(alias IIdentifier, expression IExpression, location ...*NodeLocation) *SingleColumn {
	return &SingleColumn{
		SelectItem: NewSelectItem(location...),
		alias:      alias,
		expression: expression,
	}
}

func (s *SingleColumn) Children() []Node {
	return []Node{
		s.alias,
		s.expression,
	}
}

func (s *SingleColumn) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitSingleColumn(s)
}
