package sass

import "testing"

// The test case is from dart:
// https://github.com/sass/dart-sass/blob/master/test/compressed_test.dart
func TestCompressed(t *testing.T) {
	t.Run(`in style rules`, func(t *testing.T) {
		t.Run(`removes unnecessary whitespace and semicolons`, func(t *testing.T) {
			assert(t, compressed(t, `a {x: y}`), `a{x: y}`)
		})

		t.Run(`for selectors`, func(t *testing.T) {
			t.Run(`preserves whitespace where necessary`, func(t *testing.T) {
				assert(t, compressed(t, "a b .c {x: y}"), "a b .c{x:y}")
			})
			t.Run("removes whitespace after commas", func(t *testing.T) {
				assert(t, compressed(t, "a, b, .c {x: y}"), "a,b,.c{x:y}")
			})
			t.Run("doesn't preserve newlines", func(t *testing.T) {
				assert(t, compressed(t, "a,\nb,\n.c {x: y}"), "a,b,.c{x:y}")
			})
			t.Run("removes whitespace around combinators", func(t *testing.T) {
				assert(t, compressed(t, "a > b {x: y}"), "a>b{x:y}")
				assert(t, compressed(t, "a + b {x: y}"), "a+b{x:y}")
				assert(t, compressed(t, "a ~ b {x: y}"), "a~b{x:y}")
			})
			t.Run("in prefixed pseudos", func(t *testing.T) {
				t.Run("preserves whitespace", func(t *testing.T) {
					assert(t, compressed(t, "a:nth-child(2n of b) {x: y}"), "a:nth-child(2n of b){x:y}")
				})
				t.Run("removes whitespace after commas", func(t *testing.T) {
					assert(t, compressed(t, "a:nth-child(2n of b, c) {x: y}"), "a:nth-child(2n of b,c){x:y}")
				})
			})
			t.Run("in attribute selectors with modifiers", func(t *testing.T) {
				t.Run("removes whitespace when quotes are required", func(t *testing.T) {
					assert(t, compressed(t, `a=" " b] {x: y}`), `[a=" "b]{x:y}`)
				})
				t.Run("doesn't remove whitespace when quotes aren't required", func(t *testing.T) {
					assert(t, compressed(t, `[a="b"c] {x: y}`), `[a=b c]{x:y}`)
				})
			})
		})
		t.Run("for declarations", func(t *testing.T) {
			t.Run("preserves semicolons when necessary", func(t *testing.T) {
				assert(t, compressed(t, `a {q: r; s: t}`), `a{q:r;s:t}`)
			})
			t.Run("of custom properties", func(t *testing.T) {
				t.Run("folds whitespace for multiline properties", func(t *testing.T) {
					assert(t, compressed(t, `						
						a {
						  --foo: {
							q: r;
							b {
							  s: t;
							}
						  }
						}`), `a{--foo: { q: r; b { s: t; } } }`)
				})
				t.Run("folds whitespace for single-line properties", func(t *testing.T) {
					assert(t, compressed(t, `
						a {
						  --foo: a   b\t\tc;
						}`), `a{--foo: a b\tc}`)
				})
				t.Run("preserves semicolons when necessary", func(t *testing.T) {
					assert(t, compressed(t, `						
						a {
						  --foo: {
							a: b;
						  };
						  --bar: x y;
						  --baz: q r;
						}`), `a{--foo: { a: b; };--bar: x y;--baz: q r}`)
				})
			})
		})

	})

	t.Run("values:", func(t *testing.T) {
		t.Run("numbers", func(t *testing.T) {
			t.Run("omit the leading 0", func(t *testing.T) {
				assert(t, compressed(t, `a {b: 0.123}`), `a{b:.123}`)
				assert(t, compressed(t, `a {b: 0.123px}`), `a{b:.123px}`)
			})
		})

		t.Run("lists", func(t *testing.T) {
			t.Run("don't include spaces after commas", func(t *testing.T) {
				assert(t, compressed(t, `a {b: x, y, z}`), `a{b:x,y,z}`)
			})
			t.Run("don't include spaces around slashes", func(t *testing.T) {
				assert(t, compressed(t, `
					@use "sass:list";
					a {b: list.slash(x, y, z)}`), `a{b:x/y/z}`)
			})
			t.Run("do include spaces when space-separated", func(t *testing.T) {
				assert(t, compressed(t, `a {b: x y z}`), `a{b:x y z}`)
			})
		})
		t.Run("colors", func(t *testing.T) {
			t.Run("use names when they're shortest", func(t *testing.T) {
				assert(t, compressed(t, `a {b: #f00}`), `a{b:red}`)
			})
			t.Run("use terse hex when it's shortest", func(t *testing.T) {
				assert(t, compressed(t, `a {b: white}`), `a{b:#fff}`)
			})
			t.Run("use verbose hex when it's shortest", func(t *testing.T) {
				assert(t, compressed(t, `a {b: darkgoldenrod}`), `a{b:#b8860b}`)
			})
			t.Run("use rgba() when necessary", func(t *testing.T) {
				assert(t, compressed(t, `a {b: rgba(255, 0, 0, 0.5)}`), `a{b:rgba(255,0,0,.5)}`)
			})
			t.Run("don't error when there's no name", func(t *testing.T) {
				assert(t, compressed(t, `a {b: #cc3232}`), `a{b:#cc3232}`)
			})
		})
	})

	t.Run("the top level", func(t *testing.T) {
		t.Run("removes whitespace and semicolons between at-rules", func(t *testing.T) {
			assert(t, compressed(t, `@foo; @bar; @baz;`), `@foo;@bar;@baz`)
		})
		t.Run("removes whitespace between style rules", func(t *testing.T) {
			assert(t, compressed(t, `a {b: c} x {y: z}`), `a{b:c}x{y:z}`)
		})
	})

	t.Run("@supports", func(t *testing.T) {
		t.Run("removes whitespace around the condition", func(t *testing.T) {
			assert(t, compressed(t, `@supports (display: flex) {a {b: c}}`), `@supports(display: flex){a{b:c}}`)
		})
		t.Run("preserves whitespace before the condition if necessary", func(t *testing.T) {
			assert(t, compressed(t, `@supports not (display: flex) {a {b: c}}`), `@supports not (display: flex){a{b:c}}`)
		})
	})

	t.Run("@media", func(t *testing.T) {
		t.Run("removes whitespace around the query", func(t *testing.T) {
			assert(t, compressed(t, `@media (min-width: 900px) {a {b: c}}`), `@media(min-width: 900px){a{b:c}}`)
		})
		t.Run("preserves whitespace before the query if necessary", func(t *testing.T) {
			assert(t, compressed(t, `@media screen {a {b: c}}`), `@media screen{a{b:c}}`)
		})

		// Removing whitespace after "and", "or", or "not" is forbidden because it
		// would cause it to parse as a function token.
		t.Run("removes whitespace before \"and\" when possible", func(t *testing.T) {
			assert(t, compressed(t, `
			   @media screen and (min-width: 900px) and (max-width: 100px) {
			     a {b: c}
			   }`), `@media screen and (min-width: 900px)and (max-width: 100px){a{b:c}}`)
		})
		t.Run("preserves whitespace around the modifier", func(t *testing.T) {
			assert(t, compressed(t, `@media only screen {a {b: c}}`), `@media only screen{a{b:c}}`)
		})
	})

	t.Run("@keyframes", func(t *testing.T) {
		t.Run("removes whitespace after the selector", func(t *testing.T) {
			assert(t, compressed(t, `@keyframes a {from {a: b}}`), `@keyframes a{from{a:b}}`)
		})
		t.Run("removes whitespace after commas", func(t *testing.T) {
			assert(t, compressed(t, `@keyframes a {from, to {a: b}}`), `@keyframes a{from,to{a:b}}`)
		})
	})

	t.Run("@import", func(t *testing.T) {
		t.Run("removes whitespace before the URL", func(t *testing.T) {
			assert(t, compressed(t, `@import "foo.css";`), `@import"foo.css"`)
		})
		t.Run("converts a url() to a string", func(t *testing.T) {
			assert(t, compressed(t, `@import url(foo.css);`), `@import"foo.css"`)
			assert(t, compressed(t, `@import url("foo.css");`), `@import"foo.css"`)
		})
		t.Run("removes whitespace before a media query", func(t *testing.T) {
			assert(t, compressed(t, `@import "foo.css" screen;`), `@import"foo.css"screen`)
		})
		t.Run("removes whitespace before a supports condition", func(t *testing.T) {
			assert(t, compressed(t, `@import "foo.css" supports(display: flex);`), `@import"foo.css"supports(display: flex)`)
		})
	})

	t.Run("comments", func(t *testing.T) {
		t.Run("are removed", func(t *testing.T) {
			assert(t, compressed(t, `/* foo bar */`), ``)
			assert(t, compressed(t, `
				a {
				 b: c;
				 /* foo bar */
				 d: e;
				}`), `a{b:c;d:e}`)
		})
		t.Run("remove their parents if they're the only contents", func(t *testing.T) {
			assert(t, compressed(t, `a {/* foo bar */}`), ``)
			assert(t, compressed(t, `
				a {
					/* foo bar */
					/* baz bang */
				}`), ``)
		})
		t.Run("are preserved with /*!", func(t *testing.T) {
			assert(t, compressed(t, `/*! foo bar */`), `/*! foo bar */`)
			assert(t, compressed(t, `/*! foo */\n/*! bar */`), `/*! foo *//*! bar */`)
			assert(t, compressed(t, `
				a {
					/*! foo bar */
				}`), `a{/*! foo bar */}`)
		})
	})
}

func compressed(t *testing.T, source string) string {
	// TODO compressed
	got, err := CompileString(source, &Options{
		Style: Compressed,
	})
	if err != nil {
		t.Errorf("CompileString() error = %v", err)
	}
	return got
}

func assert(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
