package tree

import "context"

type Offset struct {
	*BaseNode
	rwoCount string
}

func (o *Offset) RwoCount() string {
	return o.rwoCount
}

func NewOffset(rwoCount string, location ...*NodeLocation) *Offset {
	return &Offset{
		// TODO add context
		BaseNode: NewBaseNode(context.TODO(), location...),
		rwoCount: rwoCount,
	}
}
