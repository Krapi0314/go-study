# Go Study
Go study with [Go offical documentation](https://go.dev/doc/)

# Tutorial Summary

### [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started.html)

- `go mod init` : Enable dependency tracking
    - **Dependency Tracking**: When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.
- `go run .` : Run go code
- `go mod tidy` : Locate dependencies and download imported package and check sums + synchronize and add module
    - [pkg.go.dev](http://pkg.go.dev) site to find published modules whose packages have functions you can use in your own code

### [Tutorial:](https://go.dev/doc/tutorial/getting-started.html) [Create a module](https://go.dev/doc/tutorial/create-module.html)

- **Modules**
    - You collect one or more related packages for a discrete and useful set of functions.
    - Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.
    - If you publish a module, this *must* be a path from which your module can be downloaded by Go tools. That would be your code's repository.
- **Function**
    
    ![function-syntax.png](https://go.dev/doc/tutorial/images/function-syntax.png)
    
    - In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
- `:= Operator`
    
    ```go
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    ```
    
    - shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).
    - Long verison:
        
        ```go
        var message string
        message = fmt.Sprintf("Hi, %v. Welcome!", name)
        ```

### [Tutorial:](https://go.dev/doc/tutorial/getting-started.html) [Call your code from another module](https://go.dev/doc/tutorial/call-module-code.html)

- `main` package: In Go, code executed as an application must be in a main package.
- `go mod edit -replace` : replace moudle location (ex: replace module location to internal file system)
    
    ```go
    $ go mod edit -replace example.com/greetings=../greetings
    ```
    
- `require` in go.mod: To reference a *published* module, a go.mod file would typically omit the `replace` directive and use a `require` directive with a tagged version number at the end.
    
    ```go
    require example.com/greetings v1.1.0
    ```

### [Tutorial:](https://go.dev/doc/tutorial/getting-started.html) **[Return and handle an error](https://go.dev/doc/tutorial/handle-errors.html)**

- Common error handling in Go: Return an error as a value so the caller can check for it.
- `[errors.New](https://pkg.go.dev/errors/#example-New)()` : Import the Go standard library `errors` package to use `errors.New()` function
- `nil` : meaning no error. add in the successful return. That way, the caller can see that the function succeeded.
- `log.SetPrefix()` : Import the `[log` package](https://pkg.go.dev/log/) to print prefix at the start of its log messages, without a time stamp or source file information.

### [Tutorial:](https://go.dev/doc/tutorial/getting-started.html) **[Return a random greeting](https://go.dev/doc/tutorial/random-greeting.html)**

- slice: A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.
- `randomFormat()` : Starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
- `[]string` : Declaring a slice, you omit its size in the brackets. This tells Go that the size of the array underlying the slice can be dynamically changed.
- `math/rand` : • Use the [math/rand package](https://pkg.go.dev/math/rand/) to generate a random number for selecting an item from the slice.
- `init()` : Go executes `init` functions automatically at program startup, after global variables have been initialized.

### [Tutorial:](https://go.dev/doc/tutorial/getting-started.html) **[Return greetings for multiple people](https://go.dev/doc/tutorial/greetings-multiple-people.html)**

- Backward compatibility: Changing the function's parameter from a single name to a set of names would change the function's signature. If you had already published the module and users had already written code calling that function, that change would break their programs. → In this situation, a better choice is to write a new function with a different name. The new function will take multiple parameters. That preserves the old function for backward compatibility.
- Calling existing function in overloaded function helps reduce duplication while also leaving both functions in place.
- `make(map[*key-type*]*value-type*)` : Initialize a map
- **`for _, name := range names { }`** : If you don't need the index, you use the Go blank identifier (an underscore) to ignore it.

### [Tutorial:](https://go.dev/doc/tutorial/getting-started.html) **[Add a test](https://go.dev/doc/tutorial/add-a-test.html)**

- Implement test functions in the same package as the code you're testing.
- `Test{*Name}` :* Test function names have the form `Test*Name*`, where *Name* says something about the specific test.
- `t *testing.T` : Test functions take a pointer to the `testing`
 package's [testing.T type](https://pkg.go.dev/testing/#T) as a parameter. You use this parameter's methods for reporting and logging from your test.
- `t.FatalF()` : Use the `t` parameter's [Fatalf method](https://pkg.go.dev/testing/#T.Fatalf) to print a message to the console and end execution.
- `go test` : The `go test` command executes test functions (whose names begin with `Test`) in test files (whose names end with _test.go).
    - The output will include results for only the tests that failed, which can be useful when you have a lot of tests.
    - You can add the `-v` flag to get verbose output that lists all of the tests and their results.

### [Tutorial:](https://go.dev/doc/tutorial/getting-started.html) **[Compile and install the application](https://go.dev/doc/tutorial/compile-install.html)**

- `go build` : The [go build command](https://go.dev/cmd/go/#hdr-Compile_packages_and_dependencies) compiles the packages into an executable, along with their dependencies, but it doesn't install the results.
- `go install` : The [go install command](https://go.dev/ref/mod#go-install) compiles and installs the packages.

---

# [Tour of Go](https://go.dev/tour/) Summary

## Packages, variables, and functions

### Packages

```go
**package main**
```

- Every Go program is made up of packages.
- Programs start running in package `main`.
- By convention, the package name is the same as the last element of the import path.
    - ex) `"math/rand"` package comprises files that begin with the statement `package rand`.

### Imports

```go
import (
	"fmt"
	"math"
)
```

- It is good style to use the factored import statement

### Exported names

```go
fmt.Println(math.Pi)
```

- A name is exported if it begins with a capital letter
    - ex) `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.

### Functions

```go
func add(x int, y int) int {
	return x + y
}
```

- A function can take zero or more arguments.
- Notice that the type comes *after* the variable name.
- When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
    - ex) `x int, y int` → `x, y int`

### Multiple results

```go
func swap(x, y string) (string, string) {
	return y, x
}

a, b := swap("hello", "world")
```

- A function can return any number of results.

### **Named return values**

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

- Go's return values may be named. If so, they are treated as variables defined at the top of the function.
- These names should be used to document the meaning of the return values.
- A `return` statement without arguments returns the named return values. This is known as a "*naked*" return.
    - Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

### Variables

```go
var c, python, java bool
```

- The `var` statement declares a list of variables; as in function argument lists, the type is last.
- A `var` statement can be at package or function level.

### **Variables with initializers**

```go
var i, j int = 1, 2
var c, python, java = true, false, "no!"
```

- A var declaration can include initializers, one per variable.
- If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

### **Short variable declarations**

```go
k := 3
```

- Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.
- Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

### **Basic types**

```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

- Variable declarations may be "factored" into blocks, as with import statements.
- The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.
- Go basic types
    
    ```go
    bool
    
    string
    
    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr
    
    byte // alias for uint8
    
    rune // alias for int32
         // represents a Unicode code point
    
    float32 float64
    
    complex64 complex128
    ```
    

### **Zero values**

```go
var i int // 0
var f float64 // 0
var b bool // false
var s string // ""
```

- Variables declared without an explicit initial value are given their *zero value*.
    - The zero value is:
        - `0` for numeric types,
        - `false` for the boolean type, and
        - `""` (the empty string) for strings.

### **Type conversions**

```go
var i int = 42 // i := 42
var f float64 = float64(i) // f := float64(i)
var u uint = uint(f) // u := uint(f)
```

- The expression `T(v)` converts the value `v` to the type `T`.
- Unlike in C, in Go assignment between items of different type requires an explicit conversion.

### **Type inference**

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

- When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.
- But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant

### **Constants**

```go
const Pi = 3.14
```

- Constants are declared like variables, but with the `const` keyword.
- Constants can be character, string, boolean, or numeric values.
- Constants cannot be declared using the `:=` syntax.

### **Numeric Constants**

```go
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
```

- Numeric constants are high-precision *values*.
- An untyped constant takes the type needed by its context.