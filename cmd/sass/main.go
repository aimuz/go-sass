package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	stdin    bool
	indented bool
	loadPath string
	style    string
	charset,
	noCharset bool
	errorCss      bool
	update        string
	noSourceMap   bool
	sourceMapUrls string

	embedSources,
	embedSourceMap bool

	watch,
	poll,
	stopOnError bool
	interactive bool
	showColor   bool
	noUnicode   bool

	showVersion, showHelp bool
)

// Options ...
type Options struct {
	// Charset This flag tells Sass never to emit a @charset declaration or a UTF-8 byte-order mark.
	// By default, or if --charset is passed, Sass will insert either a @charset declaration
	// (in expanded output mode) or a byte-order mark (in compressed output mode)
	// if the stylesheet contains any non-ASCII characters.
	Charset bool

	// Style expanded, compressed
	Style string
}

var version = "v0.0.1"

func main() {
	cmd := &cobra.Command{
		Use: "sass",
		Example: "sass <input.scss> [output.css]\n" +
			"sass [<input.css>:<output.css>] [<input/>:<output/>]",
		Run: func(cmd *cobra.Command, args []string) {
			if showHelp {
				_ = cmd.Usage()
				return
			}
			if showVersion {
				fmt.Println(version)
				return
			}
			panic("TODO sass-cli")
		},
	}

	cmd.Flags().SortFlags = false

	// Input and Output
	cmd.Flags().BoolVar(&stdin, "stdin", stdin, "Read the stylesheet from stdin.")
	cmd.Flags().BoolVar(&indented, "indented", indented, "Use the indented syntax for input from stdin.")
	cmd.Flags().StringVar(&loadPath, "load-path", "", "A path to use when resolving imports.\n"+
		"May be passed multiple times.")
	cmd.Flags().StringVar(&style, "style", "expanded", "Output style. [ expanded, compressed ]")
	cmd.Flags().BoolVar(&charset, "charset", true, "Emit a @charset or BOM for CSS with non-ASCII characters.")
	cmd.Flags().BoolVar(&noCharset, "no-charset", noCharset, "Emit a @charset or BOM for CSS with non-ASCII characters.")
	cmd.Flags().BoolVar(&errorCss, "error-css", errorCss, "When an error occurs, emit a stylesheet describing it.\n"+
		"Defaults to true when compiling to a file.")
	cmd.Flags().StringVar(&update, "update", "", "Only compile out-of-date stylesheets.")

	// Source Maps
	cmd.Flags().BoolVar(&noSourceMap, "no-source-map", noSourceMap, "Whether to generate source maps.")
	cmd.Flags().StringVar(&sourceMapUrls, "source-map-urls", sourceMapUrls, "How to link from source maps to source files.")
	cmd.Flags().BoolVar(&embedSources, "embed-sources", embedSources, "Embed source file contents in source maps.")
	cmd.Flags().BoolVar(&embedSourceMap, "embed-source-map", embedSourceMap, "Embed source map contents in CSS.")

	// Other
	cmd.Flags().BoolVarP(&watch, "watch", "w", watch, "Watch stylesheets and recompile when they change.")
	cmd.Flags().BoolVar(&poll, "poll", poll, "Manually check for changes rather than using a native\n"+
		"watcher.\n"+
		"Only valid with --watch.")
	cmd.Flags().BoolVar(&stopOnError, "stop-on-error", stopOnError, "")
	cmd.Flags().BoolVarP(&interactive, "interactive", "i", interactive, "Run an interactive SassScript shell.")
	cmd.Flags().BoolVarP(&showColor, "color", "c", showColor, "")
	cmd.Flags().BoolVar(&showHelp, "help", showHelp, "Print this usage information.")
	cmd.Flags().BoolVar(&showVersion, "version", showVersion, "Print the version of Golang Sass.")

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
