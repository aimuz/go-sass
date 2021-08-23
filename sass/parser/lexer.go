package parser

import (
	"fmt"
	"github.com/aimuz/go-sass/sass/token"
	gotoken "go/token"
	"unicode"
	"unicode/utf8"
)

type ErrorHandler func(pos gotoken.Position, msg string)

type Lexer struct {
	src        []byte // source
	file       *gotoken.File
	err        ErrorHandler // error reporting; or nil
	ErrorCount int
	// scanning state
	ch       rune
	offset   int
	rdOffset int // reading offset (position after current character)
}

const bom = 0xFEFF // byte order mark, only permitted as very first character

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
				// TODO error
			}
		}
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		s.ch = -1
	}
}

func (s *Lexer) error(offs int, msg string) {
	if s.err != nil {
		// s.err(s.file.Position(s.file.Pos(offs)), msg)
		// TODO:
	}
	s.ErrorCount++
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
	offs := s.offset - 1
	for isHexDigit(s.ch) {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

// scanEscape parses an escape sequence where rune is the accepted
// escaped quote. In case of a syntax error, it stops at the offending
// character (without consuming it) and returns false. Otherwise
// it returns true.
func (s *Lexer) scanEscape(quote rune) bool {
	offs := s.offset

	var n int
	var base, max uint32
	switch s.ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		s.next()
		return true
	case '0', '1', '2', '3', '4', '5', '6', '7':
		n, base, max = 3, 8, 255
	case 'x':
		s.next()
		n, base, max = 2, 16, 255
	case 'u':
		s.next()
		n, base, max = 4, 16, unicode.MaxRune
	case 'U':
		s.next()
		n, base, max = 8, 16, unicode.MaxRune
	default:
		msg := "unknown escape sequence"
		if s.ch < 0 {
			msg = "escape sequence not terminated"
		}
		s.error(offs, msg)
		return false
	}

	var x uint32
	for n > 0 {
		d := uint32(digitVal(s.ch))
		if d >= base {
			msg := fmt.Sprintf("illegal character %#U in escape sequence", s.ch)
			if s.ch < 0 {
				msg = "escape sequence not terminated"
			}
			s.error(s.offset, msg)
			return false
		}
		x = x*base + d
		s.next()
		n--
	}

	if x > max || 0xD800 <= x && x < 0xE000 {
		s.error(offs, "escape sequence is invalid Unicode code point")
		return false
	}

	return true
}

func (s *Lexer) scanString(quote rune) string {
	offs := s.offset - 1
	for {
		ch := s.ch
		if ch == '\n' || ch == -1 {
			s.error(offs, "rune literal not terminated")
			break
		}
		s.next()
		if s.ch == quote {
			s.next()
			break
		}
	}
	return string(s.src[offs:s.offset])
}

func (s *Lexer) Init(file *gotoken.File, src []byte, err ErrorHandler) {
	// Explicitly initialize all fields since a scanner may be reused.
	if file.Size() != len(src) {
		panic(fmt.Sprintf("file size (%d) does not match src len (%d)", file.Size(), len(src)))
	}
	s.file = file
	s.src = src
	s.err = err

	s.offset = 0
	s.rdOffset = 0

	s.next()
	if s.ch == bom {
		s.next() // ignore BOM at file beginning
	}
}

// Scan ...
// return token.$ token.IDENT token.COLON token.Space token.
func (s *Lexer) Scan() (pos gotoken.Pos, tok token.Token, lit string) {
	pos = s.file.Pos(s.offset)
	ch := s.ch
	switch {
	case isWhiteSpace(ch):
		tok = token.WHITE_SPACE
		lit = s.scanWhiteSpace()
	case isLetter(ch):
		tok = token.IDENT
		lit = s.scanIdentifier()
	case isDigit(ch):
		tok = token.NUMBER
		lit = s.scanNumber()
	default:
		s.next() // always make progress
		switch ch {
		case -1:
			tok = token.EOF
		case '"':
			tok = token.STRING
			lit = s.scanString('"')
		case '#':
			tok = token.HASH
			if isLetter(s.ch) || isDigit(s.ch) {
				lit = s.scanIdentifier()
			}
		case '\'':
			tok = token.STRING
			lit = s.scanString('\'')
		case '(':
			tok = token.LPAREN
		case ')':
			tok = token.RPAREN
		case '+':
			tok = token.ADD
		case ',':
			tok = token.COMMA
		case '-':
			tok = token.SUB
		case '.':
			tok = token.PERIOD
		case '/':
			//scanner.Float
		case ':':
			tok = token.COLON
		case ';':
			tok = token.SEMICOLON
		case '<':
		case '@':
			tok = token.AT_KEYWORD
		case '[':
			tok = token.LBRACK
		case '\\':
		case ']':
			tok = token.RBRACK
		case '{':
			tok = token.LBRACE
		case '}':
			tok = token.RBRACE
		default:

		}
	}
	return
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func lower(ch rune) rune { return ('a' - 'A') | ch }

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= lower(ch) && lower(ch) <= 'f':
		return int(lower(ch) - 'a' + 10)
	}
	return 16 // larger than any legal digit val
}

// letter
// An uppercase letter or a lowercase letter.
func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isNewline(ch rune) bool {
	return ch == '\n' || ch == '\r' || ch == '\f'
}

// whitespace
// A newline, U+0009 CHARACTER TABULATION, or U+0020 SPACE.
func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || isNewline(ch)
}

func isHexDigit(ch rune) bool {
	return isDigit(ch) || isLetter(ch)
}
