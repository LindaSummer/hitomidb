package tree

type IGroupingElement interface {
	Node
}

type GroupingElement struct {
	*BaseNode
}
