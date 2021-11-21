package tree

type JoinType int

const (
	CROSS JoinType = iota
	INNER
	LEFT
	RIGHT
	FULL
	IMPLICIT
)

type IJoin interface {
	IRelation
	JoinType() JoinType
	Left() IRelation
	Right() IRelation
	Criteria() IJoinCriteria
}

type Join struct {
	*Relation
	joinType    JoinType
	left, right IRelation
	criteria    IJoinCriteria
}

func NewJoin(
	joinType JoinType,
	left IRelation,
	right IRelation,
	criteria IJoinCriteria,
	location ...*NodeLocation,
) *Join {
	return &Join{
		Relation: NewRelation(location...),
		joinType: joinType,
		left:     left,
		right:    right,
		criteria: criteria,
	}
}

func (j *Join) JoinType() JoinType {
	return j.joinType
}

func (j *Join) Left() IRelation {
	return j.left
}

func (j *Join) Right() IRelation {
	return j.right
}

func (j *Join) Criteria() IJoinCriteria {
	return j.criteria
}

func (j *Join) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitJoin(j)
}

func (j *Join) Children() []Node {
	res := []Node{
		j.left,
		j.right,
	}
	if j.criteria != nil {
		res = append(res, j.criteria.Children()...)
	}
	return res
}
