package set

import "sync"

// ThreadSafeSet is a set type which provides the thread-safety.
type ThreadSafeSet struct {
	set map[interface{}]struct{}
	mu  sync.Mutex
	rw  sync.RWMutex
}

// newThreadSafeSet creates a new *ThreadSafeSet.
func newThreadSafeSet() *ThreadSafeSet {
	return &ThreadSafeSet{set: map[interface{}]struct{}{}}
}

// Add adds a new values to set if there is enough capacity.
func (s *ThreadSafeSet) Add(val interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.add(val)
}

// add is the implementation of the Add method which is not thread-safety.
// It is called by other methods for avoiding deadlock.
func (s *ThreadSafeSet) add(val interface{}) {
	s.set[val] = setVal
}

// Append adds multiple values into set.
func (s *ThreadSafeSet) Append(values ...interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, val := range values {
		s.add(val)
	}
}

// Remove deletes the given value.
func (s *ThreadSafeSet) Remove(val interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.set, val)
}

// Contains checks the value whether exists in the set.
func (s *ThreadSafeSet) Contains(val interface{}) bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.contains(val)
}

// contains is the implementation of the Contains method which is not
// thread-safety. It is called by other methods for avoiding deadlock.
func (s *ThreadSafeSet) contains(val interface{}) bool {
	_, ok := s.set[val]
	return ok
}

// Size returns the length of the set which means that number of value of the set.
func (s *ThreadSafeSet) Size() uint {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.size()
}

// size is the implementation of the Size method which is not thread-safety.
// It is called by other methods for avoiding deadlock.
func (s *ThreadSafeSet) size() uint {
	return uint(len(s.set))
}

// Pop returns a random value from the set. If there is no element in set, it
// returns nil.
func (s *ThreadSafeSet) Pop() interface{} {
	s.rw.RLock()
	defer s.rw.RUnlock()
	for val := range s.set {
		return val
	}
	return nil
}

// Clear removes everything from the set.
func (s *ThreadSafeSet) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set = make(map[interface{}]struct{})
}

// Empty checks whether the set is empty.
func (s *ThreadSafeSet) Empty() bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return len(s.set) == 0
}

// Slice returns the elements of the set as a slice. The slice type is
// interface{}. The elements can be in any order.
func (s *ThreadSafeSet) Slice() []interface{} {
	s.rw.RLock()
	defer s.rw.RUnlock()
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
func (s *ThreadSafeSet) Union(set Set) Set {
	o := set.(*ThreadSafeSet)

	s.rw.RLock()
	o.rw.RLock()
	defer s.rw.RUnlock()
	defer o.rw.RUnlock()

	unionSet := newThreadSafeSet()
	for val := range s.set {
		unionSet.add(val)
	}
	for val := range o.set {
		unionSet.add(val)
	}
	return unionSet
}

// Intersection takes the common values from both sets and returns a new set
// that stores the common ones.
func (s *ThreadSafeSet) Intersection(set Set) Set {
	o := set.(*ThreadSafeSet)

	s.rw.RLock()
	o.rw.RLock()
	defer s.rw.RUnlock()
	defer o.rw.RUnlock()

	intersectSet := newThreadSafeSet()

	for val := range s.set {
		if o.contains(val) {
			intersectSet.Add(val)
		}
	}
	return intersectSet
}

// Difference takes the items that only is stored in s, receiver set. It returns
// a new set.
func (s *ThreadSafeSet) Difference(set Set) Set {
	o := set.(*ThreadSafeSet)

	s.rw.RLock()
	o.rw.RLock()
	defer s.rw.RUnlock()
	defer o.rw.RUnlock()

	diffSet := newThreadSafeSet()
	for val := range s.set {
		if !o.contains(val) {
			diffSet.Add(val)
		}
	}
	return diffSet
}

// IsSubset returns true if all items in the set exist in the given set.
// Otherwise, it returns false. It is not a thread-safe method. It does not handle
// the concurrency.
func (s *ThreadSafeSet) IsSubset(set Set) bool {
	o := set.(*ThreadSafeSet)

	s.rw.RLock()
	o.rw.RLock()
	defer s.rw.RUnlock()
	defer o.rw.RUnlock()

	if s.size() > o.size() {
		return false
	}

	for val := range s.set {
		if !o.contains(val) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if all items in the given set exist in the set.
// Otherwise, it returns false.
func (s *ThreadSafeSet) IsSuperset(set Set) bool {
	o := set.(*ThreadSafeSet)

	s.rw.RLock()
	o.rw.RLock()
	defer s.rw.RUnlock()
	defer o.rw.RUnlock()

	if s.size() < o.size() {
		return false
	}

	for val := range o.set {
		if !s.contains(val) {
			return false
		}
	}
	return true
}

// IsDisjoint returns true if none of the items are present in the sets.
func (s *ThreadSafeSet) IsDisjoint(set Set) bool {
	o := set.(*ThreadSafeSet)

	s.rw.RLock()
	o.rw.RLock()
	defer s.rw.RUnlock()
	defer s.rw.RUnlock()

	if s.size() == 0 || o.size() == 0 {
		return true
	}
	for val := range s.set {
		if o.contains(val) {
			return false
		}
	}
	return true
}

// Equal checks whether both sets contain exactly the same values.
func (s *ThreadSafeSet) Equal(set Set) bool {
	o := set.(*ThreadSafeSet)

	s.rw.RLock()
	o.rw.RLock()
	defer s.rw.RUnlock()
	defer o.rw.RUnlock()

	if s.size() != o.size() {
		return false
	}

	for val := range s.set {
		if !o.contains(val) {
			return false
		}
	}
	return true
}

// SymmetricDifference returns a set that contains from two sets, but not the
// items are present in both sets.
func (s *ThreadSafeSet) SymmetricDifference(set Set) Set {
	o := set.(*ThreadSafeSet)
	return s.Difference(o).Union(o.Difference(s))
}
