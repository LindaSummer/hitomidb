package tree

type SingleColumn struct {
	*SelectItem
	alias      *Identifier
	expression *Expression
}

func (s *SingleColumn) Alias() *Identifier {
	return s.alias
}

func (s *SingleColumn) Expression() *Expression {
	return s.expression
}

func NewSingleColumn(alias *Identifier, expression *Expression, location ...*NodeLocation) *SingleColumn {
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
