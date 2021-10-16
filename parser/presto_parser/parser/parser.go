package parser

import (
	"hitomidb/parser/presto_parser/tree"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SqlParser struct {
}

func (s *SqlParser) invokeParser(
	name, sql string,
	parseFunction func(parser *SqlBaseParser) antlr.ParserRuleContext,
	options ParserOptions,
) tree.Node {
	lexer := NewSqlBaseLexer(&CaseInsensitiveStream{stream: antlr.NewInputStream(sql)})
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewSqlBaseParser(tokenStream)
	parser.BuildParseTrees = true

	parsedTree := parseFunction(parser)
	visitor := NewAstBuilder(options)
	return visitor.Visit(parsedTree).(tree.Node)
}
