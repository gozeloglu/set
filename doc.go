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

You can check the set which is subset of the another set with IsSubset() method.

	isSubset := set1.IsSubset(set2)	// Returns true is set1 is subset of set2.

You can check the set which is superset of the another set with IsSuperset() method.

	isSuperset := set1.IsSuperset(set2)	// Returns true is set1 is superset of set2.

You can check whether the sets are equal with the Equal() method.

	equal := set1.Equal(set2)	// Returns true if set1 and set2 values are exactly the same

You can check whether the sets are disjoint with IsDisjoint() method.

	isDisjoint := set1.IsDisjoint(set2)	// Returns true is set1 and set2 are disjoint.

You can get the symmetric difference of two sets by SymmetricDifference() method.

	// Returns a set which is the symmetric difference of the two sets.
	symDiffSet := set1.SymmetricDifference(set2)
*/
package set
