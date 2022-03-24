package set

// ThreadUnsafeSet is a set type which does not provide the thread-safety.
type ThreadUnsafeSet struct {
	set map[interface{}]struct{}
}

// newThreadUnsafeSet creates a new *ThreadUnsafeSet.
func newThreadUnsafeSet() *ThreadUnsafeSet {
	return &ThreadUnsafeSet{set: make(map[interface{}]struct{})}
}

// Add adds a new values to set if there is enough capacity. It is not a
// thread-safe method. It does not handle the concurrency.
//
// Example:
//	s.Add("str")
//	s.Add(12)
func (s *ThreadUnsafeSet) Add(val interface{}) {
	s.set[val] = setVal
}

// Append adds multiple values into set. It is not a thread-safe method. It
// does not handle the concurrency.
//
// Example:
//	s.Append(1,2,3,4, true, false, "str")
func (s *ThreadUnsafeSet) Append(values ...interface{}) {
	for _, val := range values {
		s.Add(val)
	}
}

// Remove deletes the given value. It is not a thread-safe method. It does not
// handle the concurrency.
//
// Example:
//	s.Remove(2)
func (s *ThreadUnsafeSet) Remove(val interface{}) {
	delete(s.set, val)
}

// Contains checks the value whether exists in the set. It is not a thread-safe
// method. It does not handle the concurrency.
//
// Example:
//	exist := s.Contains(1)
func (s ThreadUnsafeSet) Contains(val interface{}) bool {
	_, ok := s.set[val]
	return ok
}

// Size returns the length of the set which means that number of value of the set.
// It is not a thread-safe method. It does not handle the concurrency.
//
// Example:
//	size := s.Size()
func (s ThreadUnsafeSet) Size() uint {
	return uint(len(s.set))
}

// Pop returns a random value from the set. If there is no element in set, it
// returns nil. It is not a thread-safe method. It does not handle the concurrency.
//
// Example:
//	val := s.Pop()
func (s ThreadUnsafeSet) Pop() interface{} {
	for val := range s.set {
		return val
	}
	return nil
}

// Clear removes everything from the set. It is not a thread-safe method. It does
// not handle the concurrency.
//
// Example:
//	s.Clear()
func (s *ThreadUnsafeSet) Clear() {
	s.set = make(map[interface{}]struct{})
}

// Empty checks whether the set is empty. It is not a thread-safe method. It does
// not handle the concurrency.
//
// Example:
//	empty := s.Empty()
func (s ThreadUnsafeSet) Empty() bool {
	return len(s.set) == 0
}

// Slice returns the elements of the set as a slice. The slice type is
// interface{}. The elements can be in any order. It is not a thread-safe method.
// It does not handle the concurrency.
//
// Example:
//	setSlice := s.Slice()
func (s ThreadUnsafeSet) Slice() []interface{} {
	values := make([]interface{}, s.Size())

	i := 0
	for k := range s.set {
		values[i] = k
		i++
	}
	return values
}

// Union returns a new Set that contains all items from the receiver Set and
// all items from the given Set. It is not a thread-safe method. It does  not
// handle the concurrency.
//
// Example:
//	unionSet := s1.Union(s2)
func (s ThreadUnsafeSet) Union(set Set) Set {
	o := set.(*ThreadUnsafeSet)
	unionSet := New(ThreadUnsafe)
	for val := range s.set {
		unionSet.Add(val)
	}
	for val := range o.set {
		unionSet.Add(val)
	}
	return unionSet
}

// Intersection takes the common values from both sets and returns a new set
// that stores the common ones. It is not a thread-safe method. It does  not
// handle the concurrency.
//
// Example:
//	intersectionSet := s1.Intersection(s2)
func (s *ThreadUnsafeSet) Intersection(set Set) Set {
	intersectSet := newThreadUnsafeSet()
	for val := range s.set {
		if set.Contains(val) {
			intersectSet.Add(val)
		}
	}
	return intersectSet
}

// Difference takes the items that only is stored in s, receiver set. It returns
// a new set. It is not a thread-safe method. It does not handle the concurrency.
//
// Example:
//	diffSet := s1.Difference(s2)
func (s *ThreadUnsafeSet) Difference(set Set) Set {
	o := set.(*ThreadUnsafeSet)
	diffSet := newThreadUnsafeSet()
	for val := range s.set {
		if !o.Contains(val) {
			diffSet.Add(val)
		}
	}
	return diffSet
}

// IsSubset returns true if all items in the set exist in the given set.
// Otherwise, it returns false. It is not a thread-safe method. It does not handle
// the concurrency.
//
// Example:
//	isSubset := s1.IsSubset(s2)
func (s *ThreadUnsafeSet) IsSubset(set Set) bool {
	if s.Size() > set.Size() {
		return false
	}

	for val := range s.set {
		if !set.Contains(val) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if all items in the given set exist in the set.
// Otherwise, it returns false. It is not a thread-safe method. It does not
// handle the concurrency.
//
// Example:
//	isSuperset := s1.IsSuperset(s2)
func (s *ThreadUnsafeSet) IsSuperset(set Set) bool {
	o := set.(*ThreadUnsafeSet)
	if s.Size() < set.Size() {
		return false
	}

	for val := range o.set {
		if !s.Contains(val) {
			return false
		}
	}
	return true
}

// IsDisjoint returns true if none of the items are present in the sets. It is
// not a thread-safe method. It does not handle the concurrency.
//
// Example:
//	isDisjoint := s1.IsDisjoint(s2)
func (s *ThreadUnsafeSet) IsDisjoint(set Set) bool {
	if s.Size() == 0 || set.Size() == 0 {
		return true
	}
	for val := range s.set {
		if set.Contains(val) {
			return false
		}
	}
	return true
}

// Equal checks whether both sets contain exactly the same values. It is not a
// thread-safe method. It does not handle the concurrency.
//
// Example:
//	equal := s1.Equal(s2)
func (s *ThreadUnsafeSet) Equal(set Set) bool {
	if s.Size() != set.Size() {
		return false
	}

	for val := range s.set {
		if !set.Contains(val) {
			return false
		}
	}
	return true
}

// SymmetricDifference returns a set that contains from two sets, but not the
// items are present in both sets. It is not a thread-safe method. It does not
// handle the concurrency.
//
// Example:
//	symmetricDiffSet := s1.SymmetricDifference(s2)
func (s *ThreadUnsafeSet) SymmetricDifference(set Set) Set {
	return s.Difference(set).Union(set.Difference(s))
}
