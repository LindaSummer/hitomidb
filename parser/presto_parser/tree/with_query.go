package tree

import "context"

type IWithQuery interface {
	Node
	Name() IIdentifier
	Query() IQuery
	ColumnNames() []IIdentifier
}

type WithQuery struct {
	*BaseNode
	name        IIdentifier
	query       IQuery
	columnNames []IIdentifier
}

func NewWithQuery(
	name IIdentifier,
	query IQuery,
	columnNames []IIdentifier,
	location ...*NodeLocation,
) *WithQuery {
	return &WithQuery{
		BaseNode:    NewBaseNode(context.TODO(), location...),
		name:        name,
		query:       query,
		columnNames: columnNames,
	}
}

func (w *WithQuery) Name() IIdentifier {
	return w.name
}

func (w *WithQuery) Query() IQuery {
	return w.query
}

func (w *WithQuery) ColumnNames() []IIdentifier {
	return w.columnNames
}

func (w *WithQuery) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitWithQuery(w)
}

func (w *WithQuery) Children() []Node {
	return []Node{w.query}
}
