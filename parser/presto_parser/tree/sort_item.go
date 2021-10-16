package tree

import "context"

type (
	OrderingType     int
	NullOrderingType int
)

const (
	_ OrderingType = iota
	ASCENDING
	DESCENDING

	_ NullOrderingType = iota
	FIRST
	LAST
	UNDEFINED
)

type SortItem struct {
	*BaseNode
	sortKey      *Expression
	ordering     OrderingType
	nullOrdering NullOrderingType
}

func (s *SortItem) SortKey() *Expression {
	return s.sortKey
}

func (s *SortItem) Ordering() OrderingType {
	return s.ordering
}

func (s *SortItem) NullOrdering() NullOrderingType {
	return s.nullOrdering
}

func NewSortItem(sortKey *Expression, ordering OrderingType, nullOrdering NullOrderingType, location ...*NodeLocation) *SortItem {
	return &SortItem{
		BaseNode:     NewBaseNode(context.TODO(), location...),
		sortKey:      sortKey,
		ordering:     ordering,
		nullOrdering: nullOrdering,
	}
}

func (s *SortItem) Children() []Node {
	return []Node{
		s.sortKey,
	}
}

func (s *SortItem) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitSortItem(s)
}
