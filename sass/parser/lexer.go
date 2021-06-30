package parser

import (
	"github.com/aimuz/go-sass/sass/token"
	gotoken "go/token"
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
	for isDecimal(s.ch) || s.ch == '.' {
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

func (s *Lexer) scanIdentifier2() string {
	offs := s.offset
	for isLetter(s.ch) || isDigit(s.ch) || s.ch == '-' {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

// Scan ...
// $roboto-font-path: "../fonts/roboto";
// return token.$ token.IDENT token.COLON token.Space token.
func (s *Lexer) Scan() (pos gotoken.Pos, tok token.Token, lit string) {
	// TODO: implementation

	ch := s.ch
	switch {
	case isLetter(ch):
		tok = token.IDENT
		if s.peek() == '@' {
			lit = s.scanIdentifier2()
			if len(lit) > 1 {
				tok = token.Lookup(lit)
			}
		} else {
			lit = s.scanIdentifier()
		}
	case isDecimal(ch) || ch == '.':
		// .10
		// 10.10
		tok = token.NUMBER
		lit = s.scanNumber()
	default:
		switch ch {
		case '@':
			tok = token.AT
		case '$':
			tok = token.DOLLAR
		case ' ':
			tok = token.SPACE
		}
	}
	return
}
