package main

import (
	"fmt"
	"github.com/gozeloglu/set"
)

func main() {
	s := set.New(set.ThreadSafe)

	s.Append(1, 2, 3, 4)

	fmt.Printf("set size: %v\n", s.Size())
	fmt.Printf("all values: ")
	for _, v := range s.Slice() {
		fmt.Printf("%+v ", v)
	}
	fmt.Println()
}
