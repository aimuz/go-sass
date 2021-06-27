package sass

import (
	"fmt"
	"go/ast"
	"go/parser"
	"testing"
)

// The test case is from dart:
// https://github.com/sass/dart-sass/blob/master/test/dart_api_test.dart
func TestAPI(t *testing.T) {
	t.Run("importers", func(t *testing.T) {
		t.Run("is used to resolve imports", func(t *testing.T) {
		})
		t.Run("are checked in order", func(t *testing.T) {
		})
	})

	t.Run("loadPaths", func(t *testing.T) {
		t.Run("is used to import file: URLs", func(t *testing.T) {
		})

		t.Run("can import partials", func(t *testing.T) {
		})

		t.Run("adds a .scss extension", func(t *testing.T) {

		})

		t.Run("adds a .sass extension", func(t *testing.T) {

		})

		t.Run("are checked in order", func(t *testing.T) {

		})
	})

	t.Run("packageResolver", func(t *testing.T) {
		t.Run("is used to import package: URLs", func(t *testing.T) {

		})

		t.Run("can resolve relative paths in a package", func(t *testing.T) {

		})

		t.Run("doesn't import a package URL from a missing package", func(t *testing.T) {

		})
	})

	t.Run("import precedence", func(t *testing.T) {
		t.Run("relative imports take precedence over importers", func(t *testing.T) {

		})

		t.Run("the original importer takes precedence over other importers", func(t *testing.T) {

		})

		t.Run("importers take precedence over load paths", func(t *testing.T) {

		})

		t.Run("importers take precedence over packageConfig", func(t *testing.T) {

		})
	})

	t.Run("charset", func(t *testing.T) {
		t.Run("= true", func(t *testing.T) {
			t.Run("doesn't emit @charset for a pure-ASCII stylesheet", func(t *testing.T) {

			})

			t.Run("emits @charset with expanded output", func(t *testing.T) {

			})

			t.Run("emits a BOM with compressed output", func(t *testing.T) {

			})
		})

		t.Run("= false", func(t *testing.T) {
			t.Run("doesn't emit @charset with expanded output", func(t *testing.T) {

			})

			t.Run("emits a BOM with compressed output", func(t *testing.T) {

			})
		})
	})
}
