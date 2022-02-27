package set

import "math/rand"

type Interface interface {
	Add(val interface{})
	Append(val ...interface{})
	Remove(val interface{})
	Contains(val interface{}) bool
	Capacity() uint
	Len() uint
	Pop() interface{}
	Clear()
}

// Set is the data structure which provides some functionalities.
type Set struct {
	set map[interface{}]struct{}
	cap uint
	len uint
}

// setVal is the value of the map. It has 0 byte.
var setVal = struct{}{}

// Add adds a new values to set if there is enough capacity.
func (s *Set) Add(val interface{}) {
	if s.len == s.cap {
		return
	}
	s.set[val] = setVal
	s.len++
}

// Append adds multiple values into set.
func (s *Set) Append(values ...interface{}) {
	for val := range values {
		s.Add(val)
	}
}

// Remove deletes the given value.
func (s *Set) Remove(val interface{}) {
	delete(s.set, val)
	s.len--
}

// Contains checks the value whether exists in the set.
func (s *Set) Contains(val interface{}) bool {
	_, ok := s.set[val]
	return ok
}

// Capacity returns the set capacity.
func (s *Set) Capacity() uint {
	return s.cap
}

// Len returns the length of the set which means that number of value of the set.
func (s *Set) Len() uint {
	return s.len
}

// Pop returns a random value from the set.
func (s *Set) Pop() interface{} {
	if s.len == 0 {
		return nil
	}
	idx := rand.Int63n(int64(s.len))
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
	s.set = make(map[interface{}]struct{}, s.cap)
}
