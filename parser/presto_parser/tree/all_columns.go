package tree

type IAllColumns interface {
	ISelectItem
	Prefix() *QualifiedName
}

type AllColumns struct {
	*SelectItem
	prefix *QualifiedName
}

func (a *AllColumns) Prefix() *QualifiedName {
	return a.prefix
}

func NewAllColumns(prefix *QualifiedName, location ...*NodeLocation) *AllColumns {
	return &AllColumns{
		SelectItem: NewSelectItem(location...),
		prefix:     prefix,
	}
}

func (a *AllColumns) Children() []Node {
	return EmptyChildren
}
