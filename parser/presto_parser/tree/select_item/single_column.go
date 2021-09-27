package select_item

import (
	"hitomidb/parser/presto_parser/tree"
	expr "hitomidb/parser/presto_parser/tree/expression"
)

type SingleColumn struct {
	*SelectItem
	alias      *expr.Identifier
	expression *expr.Expression
}

func (s *SingleColumn) Alias() *expr.Identifier {
	return s.alias
}

func (s *SingleColumn) Expression() *expr.Expression {
	return s.expression
}

func NewSingleColumn(alias *expr.Identifier, expression *expr.Expression, location ...*tree.NodeLocation) *SingleColumn {
	return &SingleColumn{
		SelectItem: NewSelectItem(location...),
		alias:      alias,
		expression: expression,
	}
}

func (s *SingleColumn) Children() []tree.Node {
	return []tree.Node{
		s.alias,
		s.expression,
	}
}

func (s *SingleColumn) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitSingleColumn(s)
}
