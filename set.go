package set

import (
	"math/rand"
	"time"
)

// S is set interface.
type S interface {
	Add(val interface{})
	Append(val ...interface{})
	Remove(val interface{})
	Contains(val interface{}) bool
	Size() uint
	Pop() interface{}
	Clear()
	Empty() bool
	Slice() []interface{}
	Union(set *Set) *Set
	Intersection(set *Set) *Set
	Difference(set *Set) *Set
	IsSubset(set *Set) bool
}

// Set is the data structure which provides some functionalities.
type Set struct {
	set map[interface{}]struct{}
}

// setVal is the value of the map. It has 0 byte.
var setVal = struct{}{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New creates a set data structure.
func New() *Set {
	return &Set{set: make(map[interface{}]struct{})}
}

// Add adds a new values to set if there is enough capacity.
func (s *Set) Add(val interface{}) {
	s.set[val] = setVal
}

// Append adds multiple values into set.
func (s *Set) Append(values ...interface{}) {
	for _, val := range values {
		s.Add(val)
	}
}

// Remove deletes the given value.
func (s *Set) Remove(val interface{}) {
	delete(s.set, val)
}

// Contains checks the value whether exists in the set.
func (s Set) Contains(val interface{}) bool {
	_, ok := s.set[val]
	return ok
}

// Size returns the length of the set which means that number of value of the set.
func (s Set) Size() uint {
	return uint(len(s.set))
}

// Pop returns a random value from the set. If there is no element in set, it
// returns nil.
func (s Set) Pop() interface{} {
	if len(s.set) == 0 {
		return nil
	}
	idx := rand.Int63n(int64(len(s.set)))
	var randVal interface{}
	for val := range s.set {
		if idx == 0 {
			randVal = val
			break
		}
		idx--
	}
	return randVal
}

// Clear removes everything from the set.
func (s *Set) Clear() {
	s.set = make(map[interface{}]struct{})
}

// Empty checks whether the set is empty.
func (s Set) Empty() bool {
	return len(s.set) == 0
}

// Slice returns the elements of the set as a slice. The slice type is
// interface{}. The elements can be in any order.
func (s Set) Slice() []interface{} {
	values := make([]interface{}, s.Size())

	i := 0
	for k := range s.set {
		values[i] = k
		i++
	}
	return values
}

// Union returns a new Set that contains all items from the receiver Set and
// all items from the given Set.
func (s Set) Union(set *Set) *Set {
	unionSet := New()
	for val := range s.set {
		unionSet.Add(val)
	}
	for val := range set.set {
		unionSet.Add(val)
	}
	return unionSet
}

// Intersection takes the common values from both sets and returns a new set
// that stores the common ones.
func (s *Set) Intersection(set *Set) *Set {
	intersectSet := New()
	for val := range s.set {
		if set.Contains(val) {
			intersectSet.Add(val)
		}
	}
	return intersectSet
}

// Difference takes the items that only is stored in s, receiver set. It returns
// a new set.
func (s *Set) Difference(set *Set) *Set {
	diffSet := New()
	for val := range s.set {
		if !set.Contains(val) {
			diffSet.Add(val)
		}
	}
	return diffSet
}

// IsSubset returns true if all items in the set exist in the given set.
// Otherwise, it returns false.
func (s *Set) IsSubset(set *Set) bool {
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
