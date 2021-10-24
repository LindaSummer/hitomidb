package tree

import (
	"fmt"
)

type ISelectStmt interface {
	Node
	Distinct() bool
	SelectItems() []ISelectItem
}

type SelectStmt struct {
	*BaseNode
	distinct    bool
	selectItems []ISelectItem
}

func NewSelectStmt(distinct bool, selectItems []ISelectItem, location ...*NodeLocation) *SelectStmt {
	return &SelectStmt{
		BaseNode: &BaseNode{
			location: GetOptionalNodeLocation(location...),
		},
		distinct:    distinct,
		selectItems: selectItems,
	}
}

func (s *SelectStmt) Distinct() bool {
	return s.distinct
}

func (s *SelectStmt) SelectItems() []ISelectItem {
	return s.selectItems
}

func (s *SelectStmt) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitSelect(s)
}

func (s *SelectStmt) GetChildren() []Node {
	return TransformToNodeArray(s.selectItems)
}

func (s *SelectStmt) String() string {
	return fmt.Sprintf("distinct: [%b]; selectItems: [%s]", s.distinct, s.selectItems)
}
