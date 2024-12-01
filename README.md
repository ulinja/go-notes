# Go Notes

This repo contains notes and example code for the Go language.

Go is a statically typed, compiled general purpose programming language, designed at Google
by Ken Thompson among others.
While its syntax is similar to that of C, Go also features garbage collection, memory safety
and concurrency.

The notes and examples in this repo loosely follow the [official Golang tour](https://go.dev/tour/).

## The Go Environment

### Tooling

Go comes bundled with CLI tools for managing source files.
The [go command](https://pkg.go.dev/cmd/go) is the centerpiece command for this.

The following is a hello world application, saved as a file called `helloworld.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

#### Formatting

The [gofmt](https://pkg.go.dev/cmd/gofmt) command can be used to format the source file:
```bash
gofmt helloworld.go
```
It can also be invoked on a directory to recursively format all Go files under it.

#### Compiling

The `go build` command is used to compile and build Go packages into executables:
```bash
go build helloworld.go
```
which will create an executable called `helloworld` in the current directory.

One amazing feature of Go is how easy it is to cross-compile for different operating systems
and architectures.
To do this, simply set the `$GOOS` and `$GOARCH` environment variables to the target OS and
architecture respectively, and invoke `go build` as usual.

The following commonly used values can be set for `$GOOS`:
- `linux`
- `windows`
- `darwin`

The following values can be set for `$GOARCH`:
- `amd64`
- `386`
- `arm`
- `arm64`

### Conventions

#### Code Style

The [Go Styleguide](https://google.github.io/styleguide/go/) describes code style guidelines in detail.

However, different to other languages, the code formatting rules in Go are not really established on paper.
Rather, the correct source code formatting is applied by invoking the `gofmt` tool.

As for variable and package naming conventions, here are some basic rules as a quick summary:
- use `MixedCaps` or `camelCase` for multi-word names
    - constants are `MixedCaps` when exported and `camelCase` when not exported
- package names should contain only lowercase characters: `helloworld` and not `helloWorld` or `hello-world`

## The Go Language

### Packages

In Go, every program is made up of packages.
Programs start running in the package `main`.

The following code shows how to import packages:
```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favourite number is %d.\n", rand.Intn(10))
}
```

By convention, the package name is the same as the last element of the import path.
For example, the `math/rand` package comprises files which begin with the statement
`package rand`.

The Go build tool views all `.go` files under the same directory as being in the same package.
This also means that you cannot declare `package main` in two different `.go` files within the
same package.

In a Go package, a name is exported if it begins with a capital letter: `Pizza` would be
exported, while `pizza` would not.
When importing a package, you can only refer to its exported names.

### Functions

Functions can take zero or more arguments. In the following example, `add` takes
two parameters of type `int`. Note how the type comes *after* the variable name:
```go
func add(x int, y int) int {
    return x + y
}
```

When two or more consecutive function parameters share a type, you can omit the type from
all but the last:
```go
func subtract(x, y int) int {
    return x - y
}
```

A function can return any number of results.
The `swap` function returns two strings:
```go
func swap(a, b string) {
    return b, a
}
```

Return values in Go can be named, in which case they are treated as being declared at the
top of the function body:
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```
The return value names should be used to document their meaning.
A `return` statement without arguments (also known as a *naked return*) returns the named
return values.

### Variables

The `var` statement declares a list of variables:
```go
package main

import "fmt"

var c, python, java bool

func main() {
    var i int
    fmt.Println(i, c, python, java)
}
```

Variable declarations can take initializers, one per variiable.
If an initializer is present, the type declaration can be omitted:
```go
var s string
var b = true
var y, z = 1, 2
```

Variable declarations can also be factored, like import statements:
```go
var (
    foo str
    bar str = "BAR"
)
```
to avoid repetition.

Inside a function, the `:=` assignment can be used in place of a var declaration with an
implicit type:
```go
func add(x, y int) int {
    result := x + y
    return result
}
```

### Basic Types

Go's basic types are:
- `bool`
- `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- `byte` (alias for uint8)
- `rune` (alias for int32, represents a Unicode code point)
- `float32`, `float64`
- `complex64`, `complex128`

The `int`, `uint` and `uintptr` types are usually 32 bits wide on 32-bit systems and
64 bits wide on 64-bit systems. When you need an integer, you should use `int` by default
unless you have a specific reason to use one of its sized or unsigned variants.

Variables declared without an initializer are given their default *zero value*.
The zero values are:
- `bool`: `false`
- `string`: `""` (empty string)
- `0` for numeric types

The expression `T(v)` converts the value `v` to the type `T`:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
or, put more simply:
```go
i := 42
f := float64(i)
u := uint(f)
```

Constants are declared like variables, but using the `const` keyword instead:
```go
const pi = 3.14
```
They can be strings, booleans, character or numeric values.
