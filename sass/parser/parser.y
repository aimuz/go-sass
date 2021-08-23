%{
package parser

import (
    "github.com/aimuz/go-sass/sass/ast"
)
%}

%union {
	offset int // offset
	item interface{}
	token string
}

%token <token>
	IDENT
	HASH
	STRING
	NUMBER
	WHITESPACE

	ADD
	SUB
	MUL
	QUO
	REM

	LPAREN
	LBRACK
	LBRACE
	COMMA
	PERIOD

	RPAREN
	RBRACK
	RBRACE
	SEMICOLON
	COLON
%%
Ident:
	IDENT {
		$$ = ast.Ident{
			Name: $1,
		}
	}
	;
Field: 
	STRING
	;
Function:
	IDENT LPAREN STRING RPAREN {

	}
	| IDENT LPAREN IDENT RPAREN {

	}
	;
SelectorClass: PERIOD IDENT {

	}
	;
SelectorID: HASH {

	}
	;
SelectorAll: MUL {

	}
	;
SelectorAttribute: LBRACK IDENT RBRACK
	| LBRACK IDENT '=' IDENT RBRACK
	{

	}
	| LBRACK IDENT '|' '=' IDENT RBRACK
	{

	}
	| LBRACK IDENT '~' '=' IDENT RBRACK
	{

	}
	| LBRACK IDENT '^' '=' IDENT RBRACK
	{

	}
	| LBRACK IDENT '$' '=' IDENT RBRACK
	{

	}
	| LBRACK IDENT '*' '=' IDENT RBRACK
	{
	
	}
	;
Selector: SelectorClass {

	}
	| SelectorID {

	}
	| SelectorAll {

	}
	| SelectorAttribute {

	}
	;
SelectorList:
	Selector {
		$$ = ast.SelectorList{

		}
	}
	| SelectorList ',' Selector {

	}
	;
selector:

%%
