// Copyright 2019 Google Inc. Use of this source code is governed by an
// MIT-style license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

// Almost all CSS output tests should go in sass-spec rather than here. This
// just covers tests that explicitly validate out that's considered too
// implementation-specific to verify in sass-spec.

package sass

import "testing"

// The test case is from dart:
// https://github.com/sass/dart-sass/blob/master/test/output_test.dart
func TestOutput(t *testing.T) {
	// Regression test for sass/dart-sass#623. This needs to be tested here
	// because sass-spec normalizes CR LF newlines.
	t.Run("normalizes newlines in a loud comment", func(t *testing.T) {
		t.Run("in SCSS", func(t *testing.T) {
			css, err := CompileString("/* foo\r\n * bar */", nil)
			errorMust(t, err)
			assert(t, css, "/* foo\n * bar */")
		})

		t.Run("in Sass", func(t *testing.T) {
			// TODO Syntax.sass
			css, err := CompileString("/*\r\n  foo\r\n  bar", nil)
			errorMust(t, err)
			assert(t, css, "/* foo\n * bar */")
		})
	})

	// Regression test for sass/dart-sass#688. This needs to be tested here
	// because it varies between Dart and Node.
	t.Run("removes exponential notation", func(t *testing.T) {
		t.Run("for integers", func(t *testing.T) {
			t.Run(">= 1e21", func(t *testing.T) {
				css, err := CompileString("a {b: 1e21}", nil)
				errorMust(t, err)
				assert(t, css, "a { b: 1${'0' * 21}; }")
			})

			// At time of writing, numbers that are 20 digits or fewer are not printed
			// in exponential notation by either Dart or Node, and we rely on that to
			// determine when to get rid of the exponent. This test ensures that if that
			// ever changes, we know about it.
			t.Run("< 1e21", func(t *testing.T) {
				css, err := CompileString("a {b: 1e20}", nil)
				errorMust(t, err)
				assert(t, css, "a { b: 1${'0' * 20}; }")
			})
		})

		t.Run("for floating-point numbers", func(t *testing.T) {
			t.Run("Infinity", func(t *testing.T) {
				css, err := CompileString("a {b: 1e20}", nil)
				errorMust(t, err)
				assert(t, css, "a { b: Infinity; }")
			})

			t.Run(">= 1e21", func(t *testing.T) {
				css, err := CompileString("a {b: 1.01e21}", nil)
				errorMust(t, err)
				assert(t, css, "a { b: 101${'0' * 19}; }")
			})

			// At time of writing, numbers that are 20 digits or fewer are not printed
			// in exponential notation by either Dart or Node, and we rely on that to
			// determine when to get rid of the exponent. This test ensures that if that
			// ever changes, we know about it.
			t.Run("< 1e21", func(t *testing.T) {
				css, err := CompileString("a {b: 1.01e20}", nil)
				errorMust(t, err)
				assert(t, css, "a { b: 101${'0' * 18}; }")
			})
		})
	})
}
