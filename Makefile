.PHONY: generate
generate:
	echo "generate"

.PHONY: release
release:
	echo "release"

.PHONY: wasm
wasm:
	tinygo build -o sass.wasm -target=wasm
