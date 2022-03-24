package set

const (
	// ThreadSafe is used in New() for creating ThreadSafeSet.
	ThreadSafe = iota

	// ThreadUnsafe is used in New() for creating ThreadUnsafeSet.
	ThreadUnsafe
)

// Set is set interface.
type Set interface {
	Add(val interface{})
	Append(val ...interface{})
	Remove(val interface{})
	Contains(val interface{}) bool
	Size() uint
	Pop() interface{}
	Clear()
	Empty() bool
	Slice() []interface{}
	Union(set Set) Set
	Intersection(set Set) Set
	Difference(set Set) Set
	IsSubset(set Set) bool
	IsSuperset(set Set) bool
	IsDisjoint(set Set) bool
	Equal(set Set) bool
	SymmetricDifference(set Set) Set
}

// setType is type of the set which can be ThreadSafe or ThreadUnsafe.
type setType int

// setVal is the value of the map. It has 0 byte.
var setVal = struct{}{}

// New creates a set data structure regarding setType. You can call the New
// function with two different setType.
//
//	safeSet := New(set.ThreadSafe)	// Creates a thread-safe set.
//	unsafeSet := New(set.ThreadUnsafe)	// Creates a thread-unsafe set.
func New(t setType) Set {
	var set Set
	switch t {
	case ThreadSafe:
		set = newThreadSafeSet()
	case ThreadUnsafe:
		set = newThreadUnsafeSet()
	}
	return set
}
