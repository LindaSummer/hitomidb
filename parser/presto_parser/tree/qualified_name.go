package tree

type QualifiedName struct {
	parts         []string
	originalParts []string
}

func NewQualifiedName(parts []string, originalParts []string) *QualifiedName {
	return &QualifiedName{parts: parts, originalParts: originalParts}
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
	return NewQualifiedName(q.parts[:len(q.parts)-1], q.originalParts[:len(q.originalParts)-1])
}
