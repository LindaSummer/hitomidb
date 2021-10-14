package tree

type NodeLocation struct {
	lineNumber, charPositionInLine int
}

func (n *NodeLocation) LineNumber() int {
	return n.lineNumber
}

func (n *NodeLocation) ColumnNumber() int {
	return n.charPositionInLine + 1
}

func NewNodeLocation(lineNumber int, charPositionInLine int) *NodeLocation {
	return &NodeLocation{lineNumber: lineNumber, charPositionInLine: charPositionInLine}
}
