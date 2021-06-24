package parser

import (
	"github.com/aimuz/go-sass/sass/token"
	gotoken "go/token"
)

type Lexer struct {
}

func (s *Lexer) Scan() (pos gotoken.Pos, tok token.Token, lit string) {
	// TODO: implementation
	return 1, token.NUMBER, ""
}

func (s *Lexer) scanNumber() (token.Token, string) {
	// TODO: implementation
	return token.NUMBER, ""
}
