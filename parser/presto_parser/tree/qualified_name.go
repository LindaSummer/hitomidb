package tree

import (
	"strings"

	funk "github.com/thoas/go-funk"
)

type QualifiedName struct {
	parts         []string
	originalParts []string
}

func NewQualifiedNameWithOriginAndParts(parts []string, originalParts []string) *QualifiedName {
	return &QualifiedName{parts: parts, originalParts: originalParts}
}

func NewQualifiedNameWithOrigin(first string, rest ...string) *QualifiedName {
	origin := make([]string, len(rest)+1)
	origin[0] = first
	copy(origin[1:], rest)
	return NewQualifiedNameWithOriginAndParts(origin, funk.Map(origin, strings.ToLower).([]string))
}

func (q *QualifiedName) Parts() []string {
	return q.parts
}

func (q *QualifiedName) OriginalParts() []string {
	return q.originalParts
}

// GetPrefix when a.b.c.d return a.b.c
// if only a return nil
func (q *QualifiedName) GetPrefix() *QualifiedName {
	if len(q.parts) == 1 {
		return nil
	}
	return NewQualifiedNameWithOriginAndParts(q.parts[:len(q.parts)-1], q.originalParts[:len(q.originalParts)-1])
}
