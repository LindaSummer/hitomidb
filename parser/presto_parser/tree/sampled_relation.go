package tree

type SampledRelationType int

const (
	_ SampledRelationType = iota
	BERNOULLI
	SYSTEM
)

type ISampledRelation interface {
	IRelation
	SampleRelation() IRelation
	SampledRelationType() SampledRelationType
	SamplePercentage() IExpression
}

type SampledRelation struct {
	*Relation
	sampleRelation      IRelation
	sampledRelationType SampledRelationType
	samplePercentage    IExpression
}

func NewSampledRelation(
	sampleRelation IRelation,
	sampledRelationType SampledRelationType,
	samplePercentage IExpression,
	location ...*NodeLocation,
) *SampledRelation {
	return &SampledRelation{
		Relation:            NewRelation(location...),
		sampleRelation:      sampleRelation,
		sampledRelationType: sampledRelationType,
		samplePercentage:    samplePercentage,
	}
}

func (s *SampledRelation) SampleRelation() IRelation {
	return s.sampleRelation
}

func (s *SampledRelation) SampledRelationType() SampledRelationType {
	return s.sampledRelationType
}

func (s *SampledRelation) SamplePercentage() IExpression {
	return s.samplePercentage
}

func (s *SampledRelation) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitSampledRelation(s)
}

func (s *SampledRelation) Children() []Node {
	return []Node{
		s.sampleRelation,
		s.samplePercentage,
	}
}
