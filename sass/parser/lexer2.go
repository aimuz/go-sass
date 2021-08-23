package parser

import (
	"bytes"
	"github.com/aimuz/go-sass/sass/token"
	"text/scanner"
)

type Lexer2 struct {
	src []byte // source

	s *scanner.Scanner
}

func (l *Lexer2) Init(src []byte) {
	l.src = src

	l.s = &scanner.Scanner{}
	l.s.Whitespace = 0
	l.s.Filename = "example"
	l.s.Init(bytes.NewBuffer(l.src))
}

func (l *Lexer2) Scan() (tok token.Token, lit string) {
	tok1 := l.s.Scan()
	switch tok1 {
	case scanner.Ident:
		tok = token.IDENT
		lit = l.s.TokenText()
		return
	case scanner.Int, scanner.Float:
		tok = token.NUMBER
		lit = l.s.TokenText()
		return
	case scanner.Char, scanner.String, scanner.RawString:
		tok = token.STRING
		lit = l.s.TokenText()
		return
	case scanner.Comment:
		tok = token.COMMENT
		lit = l.s.TokenText()
		return
	case scanner.EOF:
		tok = token.EOF
		return
	case '#':
		ch := l.s.Peek()
		tok = token.HASH
		if isLetter(ch) {
			l.s.Scan()
			lit = l.s.TokenText()
		}
		return
	case '+':
		tok = token.ADD
		return
	case '-':
		tok = token.SUB
		return
	case '*':
		tok = token.MUL
		return
	case '/':
		tok = token.QUO
		return
	case '%':
		tok = token.REM
		return
	case '(':
		tok = token.LPAREN
		return
	case '[':
		tok = token.LBRACK
		return
	case '{':
		tok = token.LBRACE
		return
	case ',':
		tok = token.COMMA
		return
	case '.':
		tok = token.PERIOD
		return
	case ')':
		tok = token.RPAREN
		return
	case ']':
		tok = token.RBRACK
		return
	case '}':
		tok = token.RBRACE
		return
	case ';':
		tok = token.SEMICOLON
		return
	case ':':
		tok = token.COLON
		return
	}
	//fmt.Printf("%s: %s %s\n", l.s.Position, scanner.TokenString(tok1), l.s.TokenText())
	return
}
