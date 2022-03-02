# set

`set` is a data structure package written by **Go**. It provides some basic set functionalities of the user. It uses 
`map` data structure under the hood. It does not have any dependency. 

## Installation

```shell
go get github.com/gozeloglu/set
```


## Example

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

	s.Append(1, 2, 3, 4, "abc")    // Add multiple values
	values := []interface{}{"github", 100, 640, 0.43, false}
	s.Append(values...) // Append the array of elements 

	s.Remove(4)
	size := s.Size()
	fmt.Println(size)   // Prints 5

	s.Pop()    // Returns random value from the set
	s.Clear()
	fmt.Println(s.Size())   // Prints 0
}
```

## Supported methods:

* `Add(val interface{})`
* `Append(val ...interface{})`
* `Remove(val interface{})`
* `Contains(val interface{})`
* `Size()`
* `Pop()`
* `Clear()`
* `Empty()`

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
