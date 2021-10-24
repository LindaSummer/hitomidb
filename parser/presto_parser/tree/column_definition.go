package tree

type IColumnDefinition interface {
	ITableElement
}

type ColumnDefinition struct {
	*TableElement
}
