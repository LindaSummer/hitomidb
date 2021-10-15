package tree

import (
	"fmt"
)

type SelectStmt struct {
	*BaseNode
	distinct    bool
	selectItems []*SelectItem
}

func NewSelectStmt(distinct bool, selectItems []*SelectItem, location ...*NodeLocation) *SelectStmt {
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

func (s *SelectStmt) SelectItems() []*SelectItem {
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
