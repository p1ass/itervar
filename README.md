# itervar

itervar is a static analysis tool that detects references to loop iterator variable.

![test_and_lint](https://github.com/p1ass/itervar/workflows/test_and_lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/p1ass/itervar)](https://goreportcard.com/report/github.com/p1ass/itervar)


## Features

- Detect code using reference to loop iterator variable, a common mistake in Go 

https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable

### Example

```go
func forLoop() {
	var out []*int
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		out = append(out, &i) // want "using reference to loop iterator variable"
	}
}
```

## Installation

### go get

```shell script
GO111MODULE=off go get github.com/p1ass/itervar/cmd/itervar
```

## Usage

```shell script
go vet -vettool=`which itervar` ./...
``` 