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

You can call the Union() method for creating a new Set which contains the all
items from the set1 and set2. It concatenates two sets and creates a new one.

	union := set1.Union(set2)

In order to take intersection of the sets, you can call the Intersection() method.

	intersect := set1.Intersection(set2)

You can find the difference between sets by Difference() method.

	// Returns a set that contains the items which are only contained in the set1.
	diffSet := set1.Difference(set2)
*/
package set
