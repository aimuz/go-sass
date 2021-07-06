// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//
package token

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
	BAD_STRING
	URL
	BADURL
	DELIM
	NUMBER
	PERCENTAGE
	DIMENSION
	WHITE_SPACE
	CDO
	CDC
	COLON
	SEMICOLON
	COMMA
	LEFT_SQUARE_BRACKET
	RIGHT_SQUARE_BRACKET
	LEFT_PARENTHESIS
	RIGHT_PARENTHESIS
	LEFT_CURLY_BRACKET
	RIGHT_CURLY_BRACKET
	COMMENT
)
