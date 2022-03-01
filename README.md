# set

`set` is a data structure package written by **Go**. It provides some basic set functionalities of the user. It uses 
`map` data structure under the hood. It sets 0-byte value which is `struct{}{}` as a value in map. 

## Installation

```shell
go get github.com/gozeloglu/set
```

## Methods:

* `Add(val interface{})`
* `Append(val ...interface{})`
* `Remove(val interface{})`
* `Contains(val interface{})`
* `Size()`
* `Pop()`
* `Clear()`
* `Empty()`

## Usage

```go
package main

import (
	"fmt"
	"github.com/gozeloglu/set"
)

func main() {
    s := set.New()
    s.Add(123)
	
    exist := s.Contains(123)
    if !exist {
		fmt.Println("123 not exist")
    }
	
	s.Append(1, 2, 3, 4, "abc")    # Add multiple values
	
	s.Remove(4)
	size := s.Size()
	fmt.Println(size)   # Prints 5
	s.Clear()
	fmt.Println(s.Size())   # Prints 0
}
```

## Tests

You can run the tests with the following command.

```shell
go test ./...
```

You can check the code coverage with the following commands:

```shell
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```
