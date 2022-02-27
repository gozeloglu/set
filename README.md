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