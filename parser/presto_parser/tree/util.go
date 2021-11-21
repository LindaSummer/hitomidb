package tree

import (
	funk "github.com/thoas/go-funk"
)

func TransformToNodeArray(list interface{}) []Node {
	return funk.Map(list, func(i interface{}) Node {
		return i.(Node)
	}).([]Node)
}

func AppendNodeIfNotNull(src []Node, item Node) []Node {
	if item != nil {
		return append(src, item)
	}
	return src
}
func GetOptionalNodeLocation(location ...*NodeLocation) *NodeLocation {
	var l *NodeLocation
	if len(location) > 0 {
		l = location[0]
	}
	return l
}
