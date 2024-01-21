# Getting started

Code in go is organized into modules wich in turn is organized into packages.
To create a module run:

```ps
go mod init org/name
```

Go creates a `go.mod` file that tracks the version of go and name of the module.

## Packages

A package groups functions together, it contains all the go-files in the same directory. `main` is a special package requiring a main function.

```go
//main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

To run the program execute:

```ps
go run .
```

## Add external packages

Find packages as [pkg.go.dev](https://pkg.go.dev/). To add the package run:

```ps
go get example.com/pkg
```
