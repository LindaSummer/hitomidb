package tree

import (
	"context"
)

type Node interface {
	Accept(visitor AstVisitor) interface{}
	Location() *NodeLocation
	SetLocation(*NodeLocation)
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
	return &BaseNode{ctx: ctx, location: GetOptionalNodeLocation(location...)}
}

func (b *BaseNode) Location() *NodeLocation {
	return b.location
}

func (b *BaseNode) SetLocation(location *NodeLocation) {
	b.location = location
}
