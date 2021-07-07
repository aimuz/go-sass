// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//
package token

import "strconv"

// Token is the set of lexical tokens of the Go programming language.
type Token int

// The list of tokens.
const (
	ILLEGAL Token = iota
	EOF
	IDENT
	FUNCTION
	AT_KEYWORD
	HASH
	STRING
	DELIM
	NUMBER
	PERCENTAGE
	DIMENSION
	WHITE_SPACE // SPACE

	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	COMMENT
)

var tokens = [...]string{
	ILLEGAL:     "ILLEGAL",
	EOF:         "EOF",
	IDENT:       "IDENT",
	AT_KEYWORD:  "@",
	HASH:        "#",
	STRING:      "STRING",
	NUMBER:      "NUMBER",
	WHITE_SPACE: "WHITE_SPACE",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
