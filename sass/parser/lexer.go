package parser

import (
	"github.com/aimuz/go-sass/sass/token"
	gotoken "go/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	src []byte // source

	// scanning state
	ch       rune
	offset   int
	rdOffset int // reading offset (position after current character)
}

// peek returns the byte following the most recently read character without
// advancing the scanner. If the scanner is at EOF, peek returns 0.
func (s *Lexer) peek() byte {
	if s.rdOffset < len(s.src) {
		return s.src[s.rdOffset]
	}
	return 0
}

// Read the next Unicode char into s.ch.
// s.ch < 0 means end-of-file.
func (s *Lexer) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset

		r, w := rune(s.src[s.rdOffset]), 1
		switch {
		//
		case r == 0:
		case r > utf8.RuneSelf:
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				// error
			}
		}
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		s.ch = -1
	}
}

func (s *Lexer) scanNumber() string {
	offs := s.offset
	for isDigit(s.ch) || s.ch == '.' {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func (s *Lexer) scanWhiteSpace() string {
	offs := s.offset
	for isWhiteSpace(s.ch) {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func (s *Lexer) scanIdentifier() string {
	offs := s.offset
	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

// Scan ...
// return token.$ token.IDENT token.COLON token.Space token.
func (s *Lexer) Scan() (pos gotoken.Pos, tok token.Token, lit string) {
	// TODO: implementation

	// current token start
	pos = gotoken.NoPos

	ch := s.ch
	switch {
	case isWhiteSpace(ch):
		tok = token.WHITE_SPACE
		lit = s.scanWhiteSpace()
	case isLetter(ch):

	case isDigit(ch):

	default:
		s.next() // always make progress
		switch ch {
		case '"':
			tok = token.STRING
			// TODO: string
		case '#':
			if isLetter(s.ch) || isDigit(s.ch) {
				tok = token.HASH
				lit = s.scanIdentifier()
			}

		case '\'':
		case '(':
		case ')':
		case '+':
		case ',':
		case '-':
		case '.':
		case '/':
		case ':':
		case ';':
		case '<':
		case '@':
			tok = token.AT_KEYWORD
		case '[':
		case '\\':
		case ']':
		case '{':
		case '}':
			tok = token.ILLEGAL
			lit = string(ch)
		}
	}
	return
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func lower(ch rune) rune { return ('a' - 'A') | ch }

// letter
// An uppercase letter or a lowercase letter.
func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

// whitespace
// A newline, U+0009 CHARACTER TABULATION, or U+0020 SPACE.
func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\t'
}
