/*
Package set is a set package provides set data structure for Go. It is written
without any dependency.

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
*/
package set
