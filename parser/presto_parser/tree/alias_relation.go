package tree

type IAliasRelation interface {
	IRelation
	AliasRelation() IRelation
	Alias() IIdentifier
	ColumnNames() []IIdentifier
}

type AliasRelation struct {
	*Relation
	aliasRelation IRelation
	alias         IIdentifier
	columnNames   []IIdentifier
}

func NewAliasRelation(
	aliasRelation IRelation,
	alias IIdentifier,
	columnNames []IIdentifier,
	location ...*NodeLocation,
) *AliasRelation {
	return &AliasRelation{
		Relation:      NewRelation(location...),
		aliasRelation: aliasRelation,
		alias:         alias,
		columnNames:   columnNames,
	}
}

func (a *AliasRelation) AliasRelation() IRelation {
	return a.aliasRelation
}

func (a *AliasRelation) Alias() IIdentifier {
	return a.alias
}

func (a *AliasRelation) ColumnNames() []IIdentifier {
	return a.columnNames
}

func (a *AliasRelation) Accept(visitor AstVisitor) interface{} {
	return visitor.VisitAliasRelation(a)
}

func (a *AliasRelation) Children() []Node {
	return []Node{a.aliasRelation}
}
