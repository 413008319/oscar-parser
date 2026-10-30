package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

// OscarParserBase implementation.
type OscarParserBase struct {
	*antlr.BaseParser
	_isVersion12 bool
	_isVersion10 bool
}

func (p *OscarParserBase) IsTableAlias() bool {
	return p.GetCurrentToken().GetTokenType() != OscarLexerJOIN
}

func (p *OscarParserBase) isVersion12() bool {
	return p._isVersion12
}

func (p *OscarParserBase) SetVersion12(value bool) {
	p._isVersion12 = value
}

func (p *OscarParserBase) isVersion10() bool {
	return p._isVersion10
}

func (p *OscarParserBase) SetVersion10(value bool) {
	p._isVersion10 = value
}
