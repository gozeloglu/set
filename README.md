# set 
[![Go Reference](https://pkg.go.dev/badge/github.com/gozeloglu/set.svg)](https://pkg.go.dev/github.com/gozeloglu/set)
[![Go Report Card](https://goreportcard.com/badge/github.com/gozeloglu/set)](https://goreportcard.com/report/github.com/gozeloglu/set)
[![GoCover](http://gocover.io/_badge/github.com/gozeloglu/set)](http://gocover.io/github.com/gozeloglu/set)
----
`set` is a data structure package written by **Go**. It provides some basic set functionalities of the user. It uses 
`map` data structure under the hood. 

* Written in vanilla Go with no dependency.
* Supports thread-safety.
* Two different set options. Thread-safe and Thread-unsafe.
* Supports any data type(integer, float, string, byte, and so on).

## Installation

```shell
go get github.com/gozeloglu/set
```


## Thread Unsafe Set Example

```go
package main

import (
    "fmt"
    "github.com/gozeloglu/set"
)

func main() {
	unsafeSet := set.New(set.ThreadUnsafe)
	unsafeSet.Add(123)

	exist := unsafeSet.Contains(123)
	if !exist {
		fmt.Println("123 not exist")
	}

	unsafeSet.Append(1, 2, 3, 4, "abc")    // Add multiple values
	values := []interface{}{"github", 100, 640, 0.43, false}
	unsafeSet.Append(values...) // Append the array of elements 

	unsafeSet.Remove(4)
	size := unsafeSet.Size()
	fmt.Println(size)   // Prints 5

	unsafeSet.Pop()    // Returns random value from the set
	unsafeSet.Clear()
	fmt.Println(unsafeSet.Size())   // Prints 0
	
	if unsafeSet.Empty() {
            fmt.Println("set is empty")
        }   
}
```

## Thread Safe Set Example

```go
package main

import (
    "fmt"
    "github.com/gozeloglu/set"
)

func main() {
	safeSet := set.New(set.ThreadSafe)
	safeSet.Append(1, 2, 3, 4)  // Add multiple values

	exist := safeSet.Contains(2)
	if !exist {
		fmt.Println("2 not exist")
	}

	values := []interface{}{"github", 100, 640, 0.43, false}
	safeSet.Append(values...) // Append the array of elements 

	safeSet.Remove(4)
	size := safeSet.Size()
	fmt.Println(size)

	safeSet.Pop()
	safeSet.Clear()
	fmt.Println(safeSet.Size())
	
	if safeSet.Empty() {
            fmt.Println("set is empty")
        }   
}
```

## Supported methods

* `Add(val interface{})`
* `Append(val ...interface{})`
* `Remove(val interface{})`
* `Contains(val interface{})`
* `Size()`
* `Pop()`
* `Clear()`
* `Empty()`
* `Slice()`
* `Union()`
* `Intersection()`
* `Difference()`
* `IsSubset()`
* `IsSuperset()`
* `IsDisjoint()`
* `Equal()`
* `SymmetricDifference()`

## Tests

  You can run the tests with the following command.

```shell
make test   # Runs all tests
make test-v # Runs all tests with -v option
make cover  # Runs all tests with -cover option. Prints out coverage result
make race   # Runs all tests with -race option. 
make bench  # Runs benchmarks
```

## LICENSE

[MIT](https://github.com/gozeloglu/set/blob/main/LICENSE)