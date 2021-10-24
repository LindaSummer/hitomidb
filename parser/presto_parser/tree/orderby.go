package tree

type IOrderBy interface {
	Node
}

type OrderBy struct {
	*BaseNode
}
