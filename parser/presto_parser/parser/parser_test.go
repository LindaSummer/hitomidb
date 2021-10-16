package parser

import (
	"hitomidb/parser/presto_parser/tree"
	"reflect"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestSqlParser_invokeParser(t *testing.T) {
	type args struct {
		name          string
		sql           string
		parseFunction func(parser *SqlBaseParser) antlr.ParserRuleContext
		options       ParserOptions
	}
	tests := []struct {
		name string
		args args
		want tree.Node
	}{
		{
			name: "simple-test",
			args: args{
				name: "statement",
				sql:  "select t1.*, t2.a, c from t3 where d > c",
				parseFunction: func(parser *SqlBaseParser) antlr.ParserRuleContext {
					return parser.SingleStatement()
				},
				options: ParserOptions{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SqlParser{}
			if got := s.invokeParser(tt.args.name, tt.args.sql, tt.args.parseFunction, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invokeParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
