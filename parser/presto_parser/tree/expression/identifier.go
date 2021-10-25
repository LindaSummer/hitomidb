package expression

import (
	"hitomidb/parser/presto_parser/tree"
	"regexp"

	"github.com/pkg/errors"
)

type Identifier struct {
	*Expression
	value     string
	delimited bool
}

var delimitedReg = regexp.MustCompile(`[a-zA-Z_]([a-zA-Z0-9_:@])*`)

func NewIdentifierWithDelimited(value string, delimited bool, location ...*tree.NodeLocation) (*Identifier, error) {
	if delimited || delimitedReg.Match([]byte(value)) {
		return nil, errors.WithStack(errors.Errorf("identifier has invalid characters: %s", value))
	}

	return &Identifier{
		Expression: NewExpression(location...),
		value:      value,
		delimited:  delimited,
	}, nil
}

func NewIdentifier(value string, location ...*tree.NodeLocation) *Identifier {
	return &Identifier{
		Expression: NewExpression(location...),
		value:      value,
		delimited:  !delimitedReg.Match([]byte(value)),
	}
}

func (i *Identifier) Value() string {
	return i.value
}

func (i *Identifier) Delimited() bool {
	return i.delimited
}

func (i *Identifier) Accept(visitor tree.AstVisitor) interface{} {
	return visitor.VisitIdentifier(i)
}

func (i *Identifier) Children() []tree.Node {
	return tree.EmptyChildren
}