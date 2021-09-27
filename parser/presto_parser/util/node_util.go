package util

import (
	"hitomidb/parser/presto_parser/tree"

	funk "github.com/thoas/go-funk"
)

func GetOptionalNodeLocation(location ...*tree.NodeLocation) *tree.NodeLocation {
	var l *tree.NodeLocation
	if len(location) > 0 {
		l = location[0]
	}
	return l
}

func TransformToNodeArray(list interface{}) []tree.Node {
	return funk.Map(list, func(i interface{}) tree.Node {
		return i.(tree.Node)
	}).([]tree.Node)
}

func AppendNodeIfNotNull(src []tree.Node, item tree.Node) []tree.Node {
	if item != nil {
		return append(src, item)
	}
	return src
}
