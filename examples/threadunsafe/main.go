package main

import (
	"fmt"
	"github.com/gozeloglu/set"
)

func main() {
	s := set.New(set.ThreadUnsafe)

	s.Add(12)
	s.Append(1, 2, 3.4, true, "str")

	fmt.Printf("set size: %v\n", s.Size())
	fmt.Printf("does 12 contain in set?: Ans: %v\n", s.Contains(12))

	s.Remove(2)
	s.Remove(100)
}
