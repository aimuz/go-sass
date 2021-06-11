# go-sass (No use, in progress)

[![License](https://img.shields.io/github/license/aimuz/go-sass)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/aimuz/go-sass)](https://goreportcard.com/report/github.com/aimuz/go-sass)
[![GoDoc](https://pkg.go.dev/badge/github.com/aimuz/go-sass?status.svg)](https://pkg.go.dev/github.com/aimuz/go-sass?tab=doc)


A Golang implementation of Sass.

## Why

This is a project for learning `AST` syntax analysis tree. It is not recommended to use it in production for the time
being. The project is not stable enough

## Using Golang Sass

### Form go get

```bash
go install github.com/aimuz/go-sass/cmd/sass@latest
```

If you are a golang user, you can install it globally through go get, which only provides cli usage

```bash
go get github.com/aimuz/go-sass/sass
```

If you want to import the library, just `go get` in your project

```go
package main

import (
	"fmt"
	"github.com/aimuz/go-sass/sass"
)

func main() {
	fmt.Println(sass.Compile("input.scss", nil))
	fmt.Println(sass.CompileString("h1 {font-size: 40px}", nil))
}
```

### From npm

We compile the core library into wasm, and then provide the NTP package through typescript wrapper

```bash
npm -i go-sass
```

Provides the same API as dart sass

### From Homebrew (OS X)

```bash
brew install go-sass
```

- [ ] TODO
