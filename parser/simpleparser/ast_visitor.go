package simpleparser

import (
	"hitomidb/parser/antlr_parser"
)

type StmtVisitor struct {
	*antlr_parser.BaseSimpleSQLParserVisitor
}

type SelectStmtVisitor struct {
	*StmtVisitor
}

type SimpleSelectStmtVisitor struct {
	*SelectStmtVisitor
}

// Visit a parse tree produced by SimpleSQLParser#SelectStmt.
func (v *StmtVisitor) VisitSelectStmt(ctx *antlr_parser.SelectStmtContext) interface{} {
	visitor := &SelectStmtVisitor{
		StmtVisitor: v,
	}
	return visitor.VisitChildren(ctx)
}

// Visit a parse tree produced by SimpleSQLParser#simple_selectstmt.
func (v *StmtVisitor) VisitSimple_selectstmt(ctx *antlr_parser.Simple_selectstmtContext) interface{} {
	//ctx.GetChildOfType()
	return nil
}
