package parser

import (
	"fmt"
	"strings"
	"text/scanner"
)

func ExampleScanner_Scan() {
	// src is the input that we want to tokenize.
	src := `.onHoverLight:hover{background-color:#e8e8e8}   ._1jRXMOCPp1s9ncnDh0phj7:hover{background-color:rgba(55,53,47,.2)} url("xxxx") url('1111')`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Whitespace = 0
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s %s\n", s.Position, scanner.TokenString(tok), s.TokenText())
	}

	/*
		.
		onHoverLight
		:
		hover
		{
		background
		-
		color
		:
		#
		e8e8e8
		}
		.
		_1jRXMOCPp1s9ncnDh0phj7
		:
		hover
		{
		background
		-
		color
		:
		rgba
		(
		55
		,
		53
		.
		47
		,
		.2
		)
		}
	*/

	//// Initialize the scanner.
	//var s Lexer
	//fset := gotoken.NewFileSet()                    // positions are relative to fset
	//file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	//s.Init(file, src, nil)
	//
	//// Repeated calls to Scan yield the token sequence found in the input.
	//for {
	//	pos, tok, lit := s.Scan()
	//	if tok == token.EOF {
	//		break
	//	}
	//	fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	//}

	//output:
	//1:1	IDENT	"cos"
	//1:4	(	""
	//1:5	IDENT	"x"
	//1:6	)	""
	//1:8	+	""
	//1:10	IMAG	"1i"
	//1:12	*	""
	//1:13	IDENT	"sin"
	//1:16	(	""
	//1:17	IDENT	"x"
	//1:18	)	""
	//1:20	;	"\n"
	//1:20	COMMENT	"// Euler"
}
