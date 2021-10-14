package tree

import (
	"context"
	"hitomidb/parser/presto_parser/util"
)

type Node interface {
	Accept(visitor AstVisitor) interface{}
	Location() *NodeLocation
	Children() []Node
	Text() string
}

var EmptyChildren []Node

type BaseNode struct {
	Node
	ctx      context.Context
	location *NodeLocation
}

func NewBaseNode(ctx context.Context, location ...*NodeLocation) *BaseNode {
	return &BaseNode{ctx: ctx, location: util.GetOptionalNodeLocation(location...)}
}

func (b *BaseNode) Location() *NodeLocation {
	return b.location
}
