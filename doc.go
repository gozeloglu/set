/*
Package set is a set package provides set data structure for Go. It is written
without any dependency.

	// Create new thread-unsafe set
	unsafeSet := set.New(set.ThreadUnsafe)
	// Create new thread-safe set
	safeSet := set.New(set.ThreadSafe)

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
