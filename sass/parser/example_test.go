package parser

import (
	"fmt"
	"github.com/aimuz/go-sass/sass/token"
	gotoken "go/token"
)

func ExampleScanner_Scan() {
	// src is the input that we want to tokenize.
	src := []byte(`#aaaaa "aaaa" #aaaa`)

	// Initialize the scanner.
	var s Lexer
	fset := gotoken.NewFileSet()                    // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

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
