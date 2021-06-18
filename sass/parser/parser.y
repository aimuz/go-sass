%{
package parser

import (
    "fmt"
)
%}


%token <token>
	ILLEGAL "ILLEGAL"
	EOF "EOF"
	COMMENT "COMMENT"
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT  "IDENT"// main
	INT "INT" // 12345
	FLOAT  "FLOAT"// 123.45
	IMAG "IMAG" // 123.45i
	CHAR "CHAR"  // 'a'
	STRING "STRING" // "abc"

	// Operators and delimiters
	ADD "ADD" // +
	SUB "SUB" // -
	MUL "MUL" // *
	QUO "QUO" // /
	REM "REM" // %

	AND "AND" // &
	OR "OR"   // |
	XOR "XOR" // ^

	ADD_ASSIGN "ADD_ASSIGN" // +=
	SUB_ASSIGN "SUB_ASSIGN" // -=
	MUL_ASSIGN "MUL_ASSIGN" // *=
	QUO_ASSIGN "QUO_ASSIGN" // /=
	REM_ASSIGN "REM_ASSIGN" // %=

	AND_ASSIGN "AND_ASSIGN" // &=
	OR_ASSIGN "OR_ASSIGN" // |=
	XOR_ASSIGN "XOR_ASSIGN" // ^=
	SHL_ASSIGN "SHL_ASSIGN" // <<=
	SHR_ASSIGN "SHR_ASSIGN" // >>=
	AND_NOT_ASSIGN "AND_NOT_ASSIGN" // &^=

	LAND  "LAND" // &&
	LOR   "LOR"  // ||
	ARROW "ARROW"// <-
	INC "INC" // ++
	DEC "DEC" // --

	EQL "EQL" // ==
	LSS "LSS" // <
	GTR "GTR" // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

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
	operator_end

	keyword_beg
	// Keywords

	USE      // @use
	FORWARD  // @forward
	IMPORT   // @import
	FUNCTION // @function
	EACH     // @each

	RETURN  // @return
	FOR     // @for
	FROM    // from
	THROUGH // through
	IF      // @if

	ELSE    // @else
	MIXIN   // @mixin
	INCLUDE // @include
	EXTEND  // @extend
	DEBUG   // @debug

	ERROR // @error
