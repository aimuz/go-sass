package sass

import "testing"

// The test case is from dart:
// https://github.com/sass/dart-sass/blob/master/test/doc_comments_test.dart
func TestDocComments(t *testing.T) {
	t.Run("documentation comments", func(t *testing.T) {
		t.Run("in SCSS", func(t *testing.T) {
			t.Run("attach to variable declarations", func(t *testing.T) {

			})

			t.Run("attach to function rules", func(t *testing.T) {

			})

			t.Run("attach to mixin rules", func(t *testing.T) {

			})

			t.Run("are null when there are no triple-slash comments", func(t *testing.T) {

			})

			t.Run("are not carried over across members", func(t *testing.T) {

			})

			t.Run("do not include double-slash comments", func(t *testing.T) {

			})
		})

		t.Run("in indented syntax", func(t *testing.T) {
			t.Run("attach to variable declarations", func(t *testing.T) {

			})

			t.Run("attach to function rules", func(t *testing.T) {

			})

			t.Run("attach to mixin rules", func(t *testing.T) {

			})

			t.Run("are null when there are no triple-slash comments", func(t *testing.T) {

			})

			t.Run("are not carried over across members", func(t *testing.T) {

			})

			t.Run("do not include double-slash comments", func(t *testing.T) {

			})

			t.Run("are compacted into one from adjacent comments", func(t *testing.T) {

			})
		})
	})
}
