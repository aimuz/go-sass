package sass

type OutputStyle string

const (
	Expanded   OutputStyle = "expanded"
	Compressed OutputStyle = "compressed"
)

type Options struct {
	// Indented ...
	Indented bool

	// Color ...
	Color bool

	// Style ...
	Style OutputStyle

	// Verbose ...
	Verbose bool

	// LoadPaths ...
	LoadPaths []string

	// Charset This flag tells Sass never to emit a @charset declaration or a UTF-8 byte-order mark.
	// By default, or if --charset is passed, Sass will insert either a @charset declaration
	// (in expanded output mode) or a byte-order mark (in compressed output mode)
	// if the stylesheet contains any non-ASCII characters.
	Charset bool
}

// Compile sass file to css file
func Compile(path string, opt *Options) (string, error) {
	// TODO: implement
	return "", nil
}

// CompileString sass to css
func CompileString(source string, opt *Options) (string, error) {
	// TODO: implement
	return "", nil
}
