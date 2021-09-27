package simpleparser

import (
	"hitomidb/parser/antlr_parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestSimpleSQLParser(t *testing.T) {
	sql := `select t1.* from t_test;`
	// Setup the input
	is := antlr.NewInputStream(sql)

	// Create the Lexer
	lexer := antlr_parser.NewSimpleSQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := antlr_parser.NewSimpleSQLParser(stream)

	var visitor SelectListener
	antlr.ParseTreeWalkerDefault.Walk(&visitor, p.Simple_selectstmt())
}
