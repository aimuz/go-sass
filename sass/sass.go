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

	// Charset ...
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
