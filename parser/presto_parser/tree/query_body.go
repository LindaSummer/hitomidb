package tree

type IQueryBody interface {
	IRelation
}

type QueryBody struct {
	*Relation
}

func NewQueryBody(location ...*NodeLocation) *QueryBody {
	return &QueryBody{Relation: NewRelation(location...)}
}

func (q *QueryBody) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitQueryBody(q)
}
