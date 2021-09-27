package tree

import (
	"fmt"
	"hitomidb/parser/presto_parser/tree/select_item"
	"hitomidb/parser/presto_parser/util"
)

type SelectStmt struct {
	*BaseNode
	distinct    bool
	selectItems []*select_item.SelectItem
}

func NewSelectStmt(distinct bool, selectItems []*select_item.SelectItem, location ...*NodeLocation) *SelectStmt {
	return &SelectStmt{
		BaseNode: &BaseNode{
			location: util.GetOptionalNodeLocation(location...),
		},
		distinct:    distinct,
		selectItems: selectItems,
	}
}

func (s *SelectStmt) Distinct() bool {
	return s.distinct
}

func (s *SelectStmt) SelectItems() []*select_item.SelectItem {
	return s.selectItems
}

func (s *SelectStmt) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitSelect(s)
}

func (s *SelectStmt) GetChildren() []Node {
	return util.TransformToNodeArray(s.selectItems)
}

func (s *SelectStmt) String() string {
	return fmt.Sprintf("distinct: [%b]; selectItems: [%s]", s.distinct, s.selectItems)
}
