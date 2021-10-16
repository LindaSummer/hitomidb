package parser

import (
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CaseInsensitiveStream struct {
	stream antlr.CharStream
}

func (c *CaseInsensitiveStream) Consume() {
	c.stream.Consume()
}

func (c *CaseInsensitiveStream) LA(i int) int {
	res := c.stream.LA(i)
	switch res {
	case 0:
		return res
	case antlr.TokenEOF:
		return res
	default:
		return int(unicode.ToUpper(rune(res)))
	}
}

func (c *CaseInsensitiveStream) Mark() int {
	return c.stream.Mark()
}

func (c *CaseInsensitiveStream) Release(marker int) {
	c.stream.Release(marker)
}

func (c *CaseInsensitiveStream) Index() int {
	return c.stream.Index()
}

func (c *CaseInsensitiveStream) Seek(index int) {
	c.stream.Seek(index)
}

func (c *CaseInsensitiveStream) Size() int {
	return c.stream.Size()
}

func (c *CaseInsensitiveStream) GetSourceName() string {
	return c.stream.GetSourceName()
}

func (c *CaseInsensitiveStream) GetText(i int, i2 int) string {
	return c.stream.GetText(i, i2)
}

func (c *CaseInsensitiveStream) GetTextFromTokens(start, end antlr.Token) string {
	return c.stream.GetTextFromTokens(start, end)
}

func (c *CaseInsensitiveStream) GetTextFromInterval(interval *antlr.Interval) string {
	return c.stream.GetTextFromInterval(interval)
}
