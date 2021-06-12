package sass

import "testing"

func compressed(t *testing.T, source string) string {
	t.Helper()
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
	t.Helper()
	if got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
