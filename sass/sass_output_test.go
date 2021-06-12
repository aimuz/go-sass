package sass

import "testing"

// The test case is from dart:
// https://github.com/sass/dart-sass/blob/master/test/output_test.dart
func TestOutput(t *testing.T) {
	// Regression test for sass/dart-sass#623. This needs to be tested here
	// because sass-spec normalizes CR LF newlines.
	t.Run("normalizes newlines in a loud comment", func(t *testing.T) {
		t.Run("in SCSS", func(t *testing.T) {

		})

		t.Run("in Sass", func(t *testing.T) {
		})
	})

	// Regression test for sass/dart-sass#688. This needs to be tested here
	// because it varies between Dart and Node.
	t.Run("removes exponential notation", func(t *testing.T) {
		t.Run("for integers", func(t *testing.T) {
			t.Run(">= 1e21", func(t *testing.T) {
			})

			// At time of writing, numbers that are 20 digits or fewer are not printed
			// in exponential notation by either Dart or Node, and we rely on that to
			// determine when to get rid of the exponent. This test ensures that if that
			// ever changes, we know about it.
			t.Run("< 1e21", func(t *testing.T) {
			})
		})

		t.Run("for floating-point numbers", func(t *testing.T) {
			t.Run("Infinity", func(t *testing.T) {
			})

			t.Run(">= 1e21", func(t *testing.T) {
			})

			// At time of writing, numbers that are 20 digits or fewer are not printed
			// in exponential notation by either Dart or Node, and we rely on that to
			// determine when to get rid of the exponent. This test ensures that if that
			// ever changes, we know about it.
			t.Run("< 1e21", func(t *testing.T) {
			})
		})
	})
}
