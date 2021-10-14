package simpleparser

import (
	"hitomidb/parser/antlr_parser"

	log "github.com/sirupsen/logrus"
)

type ColumnInfo struct {
	Table *string
}

type Node interface {
}

type Stmt struct {
	Node
}

type TableStmt struct {
	*Stmt
	TableName *string
}

type ColumnStmt struct {
	*Stmt
	columnName string
	columnInfo *ColumnInfo
}

type WhereStmt struct {
	*Stmt
}

type SelectStmt struct {
	*Stmt
	SelectFields []*ColumnStmt
	TableName    TableStmt
	Filters      []*WhereStmt
}

type SelectListener struct {
	*antlr_parser.BaseSimpleSQLParserListener
	selectStmt SelectStmt
}

//func (s *SelectListener) EnterSimple_selectstmt(c *antlr_parser.Simple_selectstmtContext) {
//	panic("implement me")
//}
//
//func (s *SelectListener) EnterResult_column(c *antlr_parser.Result_columnContext) {
//	panic("implement me")
//}
//
//func (s *SelectListener) EnterColumn_alias(c *antlr_parser.Column_aliasContext) {
//	panic("implement me")
//}
//
//func (s *SelectListener) EnterKeyword(c *antlr_parser.KeywordContext) {
//	panic("implement me")
//}
//
//func (s *SelectListener) EnterAny_name(c *antlr_parser.Any_nameContext) {
//	panic("implement me")
//}
//
//func (s *SelectListener) EnterTable_name(c *antlr_parser.Table_nameContext) {
//	panic("implement me")
//}

//// EnterSimple_selectstmt is called when production simple_selectstmt is entered.
//func (s *SelectListener) EnterSimple_selectstmt(ctx *antlr_parser.Simple_selectstmtContext) {
//}
//
//func (s *SelectListener) ExitSimple_selectstmt(c *antlr_parser.Simple_selectstmtContext) {
//}

//func (s *SelectListener) EnterAllColumns(c *antlr_parser.AllColumnsContext) {
//
//}

// ExitAllColumns is called when production Star is exited.
func (s *SelectListener) ExitAllColumns(c *antlr_parser.AllColumnsContext) {
	s.selectStmt.SelectFields = append(s.selectStmt.SelectFields, &ColumnStmt{
		columnName: "*",
		columnInfo: &ColumnInfo{
			s.selectStmt.TableName.TableName,
		},
	})
}

// EnterColFullName is called when production ColFullName is entered.
func (s *SelectListener) EnterColFullName(c *antlr_parser.ColFullNameContext) {
	s.selectStmt.SelectFields = append(s.selectStmt.SelectFields, &ColumnStmt{
		columnName: "",
		columnInfo: &ColumnInfo{},
	})
}

// EnterAllTableCols is called when entering the AllTableCols production.
func (s *SelectListener) EnterAllTableCols(c *antlr_parser.AllTableColsContext) {
	s.selectStmt.SelectFields = append(s.selectStmt.SelectFields, &ColumnStmt{
		columnName: "*",
		columnInfo: &ColumnInfo{},
	})
	tableName := c.GetChild(1)
	log.Info(tableName)
}

// EnterTable_name is called when entering the table_name production.
func (s *SelectListener) EnterTable_name(c *antlr_parser.Table_nameContext) {
	//c.GetParent().
}

// EnterTableCol is called when entering the TableCol production.
func (s *SelectListener) EnterTableCol(c *antlr_parser.TableColContext) {

}

// ExitColFullName is called when production ColFullName is exited.
func (s *SelectListener) ExitColFullName(c *antlr_parser.ColFullNameContext) {}

// EnterColAlias is called when production ColAlias is entered.
func (s *SelectListener) EnterColAlias(c *antlr_parser.ColAliasContext) {}

// ExitColAlias is called when production ColAlias is exited.
func (s *SelectListener) ExitColAlias(c *antlr_parser.ColAliasContext) {}

//func (s *SelectListener) ExitResult_column(c *antlr_parser.Result_columnContext) {
//	startToken := c.GetStart()
//	log.WithField("type", startToken.GetTokenType()).Info("ExitResult_column")
//	switch startToken.GetTokenType() {
//	case antlr_parser.SimpleSQLLexerSTAR:
//		log.Info("need to get info schema")
//	case antlr_parser.SimpleSQLParserRULE_column_full_name:
//		log.WithField("column_full_name", startToken.GetText())
//		s.selectStmt.SelectFields = append(s.selectStmt.SelectFields, ColumnStmt{
//			columnName: "",
//			columnInfo: ColumnInfo{},
//		})
//	case antlr_parser.SimpleSQLParserRULE_column_alias:
//		log.WithField("simple_cloumn_name", startToken.GetText())
//		s.selectStmt.SelectFields = append(s.selectStmt.SelectFields, ColumnStmt{
//			columnName: c.GetText(),
//			columnInfo: ColumnInfo{},
//		})
//	}
//}

func (s *SelectListener) ExitColumn_full_name(c *antlr_parser.Column_full_nameContext) {
	child := c.GetChild(0)
	log.WithField("child", child)
}

//func (s *SelectListener) ExitColumn_alias(c *antlr_parser.Column_aliasContext) {
//	panic("implement me")
//}

func (s *SelectListener) ExitTable_name(c *antlr_parser.Table_nameContext) {
	text := c.GetText()
	s.selectStmt.TableName.TableName = &text
}
