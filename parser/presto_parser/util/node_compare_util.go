package util

import (
	"hitomidb/parser/presto_parser/tree"
	"strings"

	funk "github.com/thoas/go-funk"

	"github.com/google/go-cmp/cmp"
)

func CompareNode(lhs, rhs tree.Node, comparePosition bool) bool {
	if lhs == nil && rhs == nil {
		return true
	}

	if (lhs == nil) != (rhs == nil) {
		return false
	}

	comparer := &nodeDataComparer{
		comparePosition: comparePosition,
		rhs:             rhs,
	}
	return comparer.visit(lhs)
}

func CompareStringPointerInsensitive(lhs, rhs *string) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if (lhs == nil) != (rhs == nil) {
		return false
	}

	return strings.EqualFold(*lhs, *rhs)
}

func CompareQualifiedName(lhs, rhs *tree.QualifiedName) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if (lhs == nil) != (rhs == nil) {
		return false
	}

	return cmp.Equal(lhs, rhs)
}

func CompareNodes(lhs, rhs []tree.Node, comparePosition bool) bool {
	if (lhs == nil) && (rhs == nil) {
		return true
	}

	if (lhs == nil) != (rhs == nil) {
		return false
	}

	if len(lhs) != len(rhs) {
		return false
	}

	for i := 0; i < len(lhs); i++ {
		if !CompareNode(lhs[i], rhs[i], comparePosition) {
			return false
		}
	}
	return true
}

type nodeDataComparer struct {
	comparePosition bool
	rhs             tree.Node
}

func (n *nodeDataComparer) visit(lhs tree.Node) bool {
	equal, ok := lhs.Accept(n).(bool)
	equal = ok && equal && CompareNodes(lhs.Children(), n.rhs.Children(), n.comparePosition)

	if n.comparePosition {
		equal = equal &&
			CompareLocation(lhs.Location(), n.rhs.Location())
	}

	return equal
}

func CompareLocation(lhs, rhs *tree.NodeLocation) bool {
	if (lhs == nil) || (rhs == nil) {
		return false
	}

	return lhs.ColumnNumber() == rhs.ColumnNumber() &&
		lhs.LineNumber() == rhs.LineNumber()
}

func (n *nodeDataComparer) VisitSelect(lhs tree.ISelectStmt) interface{} {
	rhs, ok := n.rhs.(tree.ISelectStmt)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return lhs.Distinct() == rhs.Distinct()
}

func (n *nodeDataComparer) VisitTableElement(lhs tree.ITableElement) interface{} {
	rhs, ok := n.rhs.(tree.ITableElement)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return true
}

func (n *nodeDataComparer) VisitSingleColumn(lhs tree.ISingleColumn) interface{} {
	rhs, ok := n.rhs.(tree.ISingleColumn)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return true
}

func (n *nodeDataComparer) VisitStatement(lhs tree.IStatement) interface{} {
	rhs, ok := n.rhs.(tree.IStatement)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return true
}

func (n *nodeDataComparer) VisitQuery(lhs tree.IQuery) interface{} {
	rhs, ok := n.rhs.(tree.IQuery)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return CompareStringPointerInsensitive(lhs.Limit(), rhs.Limit())
}

func (n *nodeDataComparer) VisitRelation(lhs tree.IRelation) interface{} {
	rhs, ok := n.rhs.(tree.IRelation)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return true
}

func (n *nodeDataComparer) VisitSampledRelation(lhs tree.ISampledRelation) interface{} {
	rhs, ok := n.rhs.(tree.ISampledRelation)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return lhs.SampledRelationType() == rhs.SampledRelationType()
}

func (n *nodeDataComparer) VisitQueryBody(lhs tree.IQueryBody) interface{} {
	rhs, ok := n.rhs.(tree.IQueryBody)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return true
}

func (n *nodeDataComparer) VisitSelectItem(lhs tree.ISelectItem) interface{} {
	rhs, ok := n.rhs.(tree.ISelectItem)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return true
}

func (n *nodeDataComparer) VisitQueryTable(lhs tree.IQueryTable) interface{} {
	rhs, ok := n.rhs.(tree.IQueryTable)
	if !ok || rhs == nil || lhs == nil {
		return false
	}
	ln, rn := lhs.Name(), rhs.Name()
	return CompareQualifiedName(&ln, &rn)
}

func (n *nodeDataComparer) VisitQuerySpecification(lhs tree.IQuerySpecification) interface{} {
	rhs, ok := n.rhs.(tree.IQuerySpecification)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return CompareStringPointerInsensitive(lhs.Limit(), rhs.Limit())
}

func (n *nodeDataComparer) VisitWith(lhs tree.IWith) interface{} {
	rhs, ok := n.rhs.(tree.IWith)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return lhs.Recursive() == rhs.Recursive()
}

func (n *nodeDataComparer) VisitWithQuery(lhs tree.IWithQuery) interface{} {
	rhs, ok := n.rhs.(tree.IWithQuery)
	if !ok || rhs == nil || lhs == nil {
		return false
	}

	return CompareNode(lhs.Name(), rhs.Name(), n.comparePosition) &&
		CompareNodes(funk.Map(lhs.ColumnNames(), func(identifier tree.IIdentifier) tree.Node { return identifier }).([]tree.Node),
			funk.Map(rhs.ColumnNames(), func(identifier tree.IIdentifier) tree.Node { return identifier }).([]tree.Node),
			n.comparePosition)
}

func (n *nodeDataComparer) VisitJoin(join tree.IJoin) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitJoinOn(on tree.IJoinOn) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitJoinUsing(using tree.IJoinUsing) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitNaturalJoin(join tree.INaturalJoin) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitOrderBy(orderBy tree.IOrderBy) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitGroupBy(groupBy tree.IGroupBy) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitOffset(offset tree.IOffset) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitTable(table tree.ITable) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitColumnDefinition(definition tree.IColumnDefinition) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitProperty(property tree.IProperty) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitAliasRelation(alias tree.IAliasRelation) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitExpression(expression tree.IExpression) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitIdentifier(identifier tree.IIdentifier) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitDereferenceExpression(expression tree.IDereferenceExpression) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitComparisonExpression(expression tree.IComparisonExpression) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitLogicalBinaryExpression(expression tree.ILogicalBinaryExpression) interface{} {
	panic("implement me")
}

func (n *nodeDataComparer) VisitSortItem(sortItem tree.ISortItem) interface{} {
	panic("implement me")
}
