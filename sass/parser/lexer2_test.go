package parser

import (
	"fmt"
	"github.com/aimuz/go-sass/sass/token"
	"testing"
)

func TestLexer2_Scan(t *testing.T) {
	// src is the input that we want to tokenize.
	src := `.onHoverLight:hover{background-color:#e8e8e8}   ._1jRXMOCPp1s9ncnDh0phj7:hover{background-color:rgba(55,53,47,.2)} url("xxxx") url('1111')`

	var s Lexer2
	s.Init([]byte(src))
	for tok, lit := s.Scan(); tok != token.EOF; tok, lit = s.Scan() {
		fmt.Println(tok, lit)
	}
}
