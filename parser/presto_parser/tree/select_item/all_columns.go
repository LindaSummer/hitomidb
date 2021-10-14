package select_item

import "hitomidb/parser/presto_parser/tree"

type AllColumns struct {
	*SelectItem
	prefix *tree.QualifiedName
}

func (a *AllColumns) Prefix() *tree.QualifiedName {
	return a.prefix
}

func NewAllColumns(prefix *tree.QualifiedName, location ...*tree.NodeLocation) *AllColumns {
	return &AllColumns{
		SelectItem: NewSelectItem(location...),
		prefix:     prefix,
	}
}

func (a *AllColumns) Children() []tree.Node {
	return tree.EmptyChildren
}
