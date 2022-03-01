package set

import (
	"math/rand"
	"time"
)

type S interface {
	Add(val interface{})
	Append(val ...interface{})
	Remove(val interface{})
	Contains(val interface{}) bool
	Size() uint
	Pop() interface{}
	Clear()
	Empty() bool
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
	s := &Set{}
	s.set = make(map[interface{}]struct{})
	return s
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

// Pop returns a random value from the set.
func (s Set) Pop() interface{} {
	idx := rand.Int63n(int64(len(s.set)))
	var randVal interface{}
	for i, val := range s.set {
		if i == idx {
			randVal = val
			break
		}
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
