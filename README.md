# Go Notes

This repo contains notes, code challenges and example code for learning the Go language.

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

To run the code, use the `go run` command:
```bash
go run helloworld.go
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

#### Structuring A Go Project

The Go docs provide some guidance on [how to structure a Go project](https://go.dev/doc/modules/layout),
as well as tips on [how to distribute Go modules](https://go.dev/doc/modules/managing-source) for others
to use.

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

Functions are values too, which can be passed around just like other values by their function name.

Go functions may be closures. Closures are functions which reference values from outside its body
(yet not in the global scope).
The following examples illustrate a counter variable which pollutes the global scope, in contrast
to a closure counter function which maintains its own state:
```go
var count int   // the count var is global

func increment() int {
    count++
    return count
}
```

```go
func incrementer() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

inc := incrementer() // Each instance has its own state
fmt.Println(inc())   // Output: 1
```

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

Variable declarations can take initializers, one per variable.
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

### Loops

Go only has one type of looping construct, the `for` loop.

The `for` loop has three basic components, separated by semicolons and not surrounded by braces:
- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += 1
}
```

The init and post statements are optional:
```go
sum := 1
for ; sum < 1000; {
    sum += sum
}
```

The init and post statements can also be completely omitted, which is equivalent to a `while` loop
in other languages:
```go
sum := 1
for sum < 1000 {
    sum += sum
}
```

Omitting the condition expression is equivalent to an infinite loop:
```go
package main

import "fmt"

func main() {
    for {
        fmt.Print("HA")
    }
}
```

### Conditionals

Go's `if` statements are like its `for` loops:
```go
if x < 0 {
    x = 0
} else if x < 100 {
    x = 100
} else {
    x = 1000
}
```

Like `for`, `if` statements can also execute a short expression, such as an assignment before the
execution of the conditional statements:
```go
x := 10
n := 2
m := 999
if v := math.Pow(x, n); v < m {
    fmt.Println(v, "(small)")
} else {
    fmt.Println(v, "(big)")
}
```
Values declared in this expression are only valid within the scope of the `if`/`else` block.

### Defer statements

The `defer` statement defers the execution of a function until the end of the surrounding function:
```go
package main

import "fmt"

func main() {
    defer fmt.Println(" world!")

    fmt.Print("Hello")
}
```

The defer statement is evaluated immediately, but its execution is deferred until the end of the surrounding
function.
This is particularly useful for deallocating resources independently of the rest of the function's flow:
```go
func writeFile(v string) {
    f, err := os.Create("file.txt")
    if err != nil {
        panic("Could not create file!")
    }
    defer f.close()

    // ...
    fmt.Fprintf(f, v)
    // ...
}
```

### Pointers

Go has pointers :fire:.

A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value, with the initial value of `nil`:
```go
var p *int
```

The `&` operator generates a pointer to its operand:
```go
var i := 42
p = &i
```

The `*` operator denotes the pointers underlying value:
```go
fmt.Println(*p) // read i through pointer p
*p = 21 // set i through pointer p
```
which is also known as *dereferencing* or *indirecting*.

Unlike C, Go does not have pointer arithmetic.

### Structs

A `struct` is a collection of fields.
Struct fields are accessed using dot notation:
```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```

Struct fields can also be accessed through a struct pointer.
To access the field `X` of a struct when we have a struct pointer `p`
we could write `(*p).X`. This notation is cumbersome though, instead we
can just write `p.X` without the explicit dereference.
```go
v := Vertex{1, 2}
p := &v
p.X = 1e9
```

A struct literal denotes a newly allocated struct value by listing the values
of its fields.

You can list just a subset of fields by using the `Name:` syntax (and the order of
named fields is irrelevant).

The special prefix `&` returns a pointer to the struct value.
```go
package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}   // has type Vertex
    v2 = Vertex{X: 1}   // Y:0 is implicit
    v3 = Vertex{}       // X:0 Y:0
    p = &Vertex{1, 2}   // has type *Vertex
)

func main() {
    fmt.Println(v1, v2, v3, p)
}
```

### Arrays

The type `[n]T` is an array of `n` values of type `T`.

The expression
```go
var a [10]int
```
declares `a` as an array of ten integers.

An array's size is part of its type, so arrays cannot be resized.

```go
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1], a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
```

### Slices

Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized,
flexible view into the elements of an array. In practice, slices are much more
common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is created by specifying two indices, a low and high bound, separated by a colon:
```go
a[low:high]
```
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of `a`:
```go
a[1:4]
```

A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the contents of the underlying array.
Other slices that share the same underlying array will see those changes.

A slice literal is like an array literal without the length.

This is an array literal:
```go
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:
```go
[]bool{true, true, false}
```

When slicing, you may omit the high or low bounds to use their defaults instead.
The default is zero for the low bound and the length of the slice for the high bound.

For the array
```go
var a [10]int
```

these slice expressions are equivalent:
```go
a[0:10]
a[:10]
a[0:]
a[:]
```

A slice has both a *length* and a *capacity*.
The length and capacity of a slice `s` can be obtained using the expressions
`len(s)` and `cap(s)`.

The length of a slice is the number of elements it contains.

The capacity is the number of elements in the underlying array, counting from
the first element in the slice.

A slice's length can be extended by re-slicing it, provided it has sufficient
capacity.
```go
package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length.
    s = s[:4]
    printSlice(s)

    // Drop its first two values.
    s = s[2:]
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

The zero value of a slice is `nil`.
A nil slice has length and capacity zero and no underlying array.
```go
var s []int
```

Slices can be created with the builtin `make` function.
This is how you create dynamically-sized arrays.o

The `make` function allocates a zeroed array and returns a slice that refers to
that array:
```go
a := make([]int, 5) // len(a)=5
```

To specify a capacity, specify a third argument to `make`:
```go
b := make([]int, 0, 5) // len(b)=0 cap(b)=5
```

Slices can contain any type, including other slices.

It is common to append values to a slice, which the builtin `append` method is used for:
```
func append(s []T, vs ...T) []T
```
The first parameter of `append` is a slice of type `T`, and the rest are `T` values appended
to the slice.
The resulting value of `append` is a slice of type `T` containing all the elements of the
original slice, as well as the appended values.

If the underlying array of `s` is too small to fit all appended values, a bigger array will be
allocated, and the returned slice will point to this newly allocated array.

For more information, see the [documentation for append](https://go.dev/pkg/builtin/#append)
and the [article on slice usage and internals](https://go.dev/blog/go-slices-usage-and-internals).

The `range` form of the `for` loop iterates over a slice or map.
When ranging over a slice, two values are returned for each iteration: the current index, and a copy
of the element at that index.
```go
package main

import "fmt"

func main() {
    nums := []int{2, 4, 6, 8, 10}
    for i, v := range nums {
        fmt.Printf("Value at %d is %d\n", i, v)
    }
}
```

You can skip the index or value by assiging to `_`:
```go
for _, v := range nums {
    // ...
}
```

If you only want the index, you can omit `v`:
```go
for i := range nums {
    // ...
}
```

### Maps

A map maps keys to values.

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added to it.

The `make` function returns a map of the given type, initialized and ready for use.
```go
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m := make(map[string]Vertex)
    m["Bell Labs"] = Vertex{40.123, -74.231}
}
```

Map literals are like struct literals, but the keys are required:
```go
var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.123, -74.231,
    },
    "Google": Vertex{
        54.543, 43.123,
    }
}
```

If the top name is just the name of a type, you can omit it from the elements of the literal:
```go
var m := map[string]Vertex{
    "Bell Labs": {
        40.123, -74.231,
    },
    "Google": {
        54.543, 43.123,
    }
}
```

Insert or update an element in a map using square bracket notation:
```go
m[key] = elem
```

Retrieve an element:
```go
elem = m[key]
```

Delete an element:
```go
delete(m, key)
```

Test that a key is present with a two value assignment:
```go
elem, ok = m[key]
```
If `key` is in `m`, `elem` will be its value and `ok` will be `true`.
If not, `ok` is `false` and `elem` will be the zero value of that map's type.

If `elem` or `ok` are not yet declared, you can use the short declaration form:
```go
elem, ok := m[key]
```

### Methods

Go does not have classes. However, *methods* can be defined on types.

A method is a function with a special argument called a *receiver*.
The receiver appears in its own argument list between the `func` keyowrd and the method name:
```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct{
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
```
In this example, the `Abs` method has a receiver `v` of type `Vertex`.

You can declare methods on non-struct types as well:
```go
type MyThing uint

func (thing MyThing) DoSomething() {
    //...
}
```

However, methods must be declared in the same module in which the receiver's type is declared.

Methods can (and most oftenly do) have pointer receivers, in which case they modify the value to which
the receiver points.

### Interfaces

An *interface type* is a set of method signatures.

A value of an interface can hold any value that implements those methods.
```go
package main

import (
    "fmt"
)

type SelfIdentifier interface {
    IdentifySelf() string
}

type Human struct {
    Name string
}

func (h *Human) IdentifySelf() string {
    return fmt.Sprintf("Hello, I am %s.", h.Name)
}

type Robot struct {
    SerialNumber uint
}

func (r *Robot) IdentifySelf() string {
    identifier := fmt.Sprintf("0X010-%d", r.SerialNumber)
    return fmt.Sprintf("Beep Boop. I am %s.", identifier)
}

func main() {
    h := Human{"John Smith"}
    r := Robot{4269}

    entities := []SelfIdentifier{&h, &r}
    for _, e := range entities {
        fmt.Println(e.IdentifySelf())
    }
}
```

### Goroutines

A *goroutine* is a lightweight thread managed by the Go runtime:
```go
go f(x, y, z)
```
starts a new goroutine running
```go
f(x, y, z)
```

The evaluation of `f`, `x`, `y` and `z` happens in the current goroutine, and the execution
of `f` happens in the new goroutine.

Goroutines run in the same address space, so care must be taken to synchronize access to shared memory.
The `sync` package provides useful primitives for this purpose.

### Channels

Channels are a typed conduit through which you can send and receive values with the channel operator `<-`.
```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
The data flows in the direction of the arrow.

Like maps and slices, channels must be created before use:
```go
ch := make(chan int)
```

By default, sends and receives block until the other side is ready, allowing goroutines to sync without
explicitly settings locks or condition variables.

```go
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}
```

Channels can be *buffered*. Provide the buffer length as the second argument to `make` to initialize a
buffered channel:
```go
ch := make(chan int, 100)
```
Sends to a buffered channel block only when the buffer is full.
Receives block when the buffer is empty.

A sender can `close` a channel to indicate that no more values will be sent.
Receivers can test whether a channel is closed by assigning a second parameter to the receive expression:
```go
v, ok := <-ch
```
`ok` is `false` when there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

> Only the sender should close a channel, never the receiver.

```go
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}
```

### Select

The `select` statement lets a goroutine wait on multiple communication operations.
It blocks until one of its cases can run, then executes that case.
If multiple cases can run, it chooses one at random.

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
```

The `default` case is run if no other case is ready.
