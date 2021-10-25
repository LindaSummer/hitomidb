package tree

import funk "github.com/thoas/go-funk"

type IWith interface {
	Node
	Recursive() bool
	Queries() []IWithQuery
}

type With struct {
	*BaseNode
	recursive bool
	queries   []IWithQuery
}

func (w *With) Recursive() bool {
	return w.recursive
}

func (w *With) Queries() []IWithQuery {
	return w.queries
}

func (w *With) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitWith(w)
}

func (w *With) Children() []Node {
	return funk.Map(w.queries, func(w IWithQuery) Node { return w }).([]Node)
}
