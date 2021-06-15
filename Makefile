.PHONY: generate
generate:
	echo "generate"

.PHONY: parser
parser:
	goyacc -o /dev/null sass/parser/parser.y
	goyacc -o sass/parser/parser.go sass/parser/parser.y 2>&1

.PHONY: release
release:
	echo "release"

.PHONY: wasm
wasm:
	tinygo build -o sass.wasm -target=wasm
