package set

import (
	"testing"
)

func TestThreadSafeSet_Add(t *testing.T) {
	testCases := []struct {
		name    string
		values  []interface{}
		expSize uint
		expSet  map[interface{}]struct{}
	}{
		{
			name:    "Add int value",
			values:  []interface{}{100},
			expSize: 1,
			expSet: map[interface{}]struct{}{
				100: setVal,
			},
		},
		{
			name:    "Add float value",
			values:  []interface{}{12.332},
			expSize: 1,
			expSet: map[interface{}]struct{}{
				12.332: setVal,
			},
		},
		{
			name:    "Add bool value",
			values:  []interface{}{true},
			expSize: 1,
			expSet: map[interface{}]struct{}{
				true: setVal,
			},
		},
		{
			name:    "Add byte value",
			values:  []interface{}{byte('e')},
			expSize: 1,
			expSet: map[interface{}]struct{}{
				byte('e'): setVal,
			},
		},
		{
			name:    "Add string value",
			values:  []interface{}{"str-key"},
			expSize: 1,
			expSet: map[interface{}]struct{}{
				"str-key": setVal,
			},
		},
		{
			name: "Add multiple values",
			values: []interface{}{1, 1.3, "str", byte('c'), struct {
				key string
			}{
				key: "key",
			}},
			expSize: 5,
			expSet: map[interface{}]struct{}{
				1:         setVal,
				1.3:       setVal,
				"str":     setVal,
				byte('c'): setVal,
				struct {
					key string
				}{
					key: "key",
				}: setVal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			for _, v := range tc.values {
				s.Add(v)
			}
			if s.Size() != tc.expSize {
				t.Errorf("expected length %v, actual length %v", tc.expSize, s.Size())
			}
			for v := range s.set {
				if _, ok := tc.expSet[v]; !ok {
					t.Errorf("value %v is expected, but not exist", v)
				}
			}
		})
	}
}

func TestThreadSafeSet_Append(t *testing.T) {
	testCases := []struct {
		name    string
		values  []interface{}
		expSize uint
	}{
		{
			name:    "Append nothing",
			expSize: 0,
		},
		{
			name:    "Append single value",
			values:  []interface{}{"single"},
			expSize: 1,
		},
		{
			name:    "Append multiple values",
			values:  []interface{}{1, 2, 3, 4},
			expSize: 4,
		},
		{
			name:    "Append different types",
			values:  []interface{}{1, 2, 3.5, uint16(4), "set", "key", true, byte('a')},
			expSize: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)

			size := s.Size()
			if size != tc.expSize {
				t.Errorf("expected length %v, actual length %v", tc.expSize, size)
			}
			for _, v := range tc.values {
				if _, ok := s.set[v]; !ok {
					t.Errorf("expected %v, but not found", v)
				}
			}
		})
	}
}

func TestThreadSafeSet_Remove(t *testing.T) {
	testCases := []struct {
		name            string
		values          []interface{}
		expRemoveValues []interface{}
		remainingValues []interface{}
		expSize         uint
	}{
		{
			name:            "Remove from the empty set",
			expRemoveValues: []interface{}{1, 2, 3},
		},
		{
			name:            "Remove from non-empty set",
			values:          []interface{}{1, 2, 3, 4.98, "remove", "set", byte('a')},
			expRemoveValues: []interface{}{1, 2, 3, 4.98, "remove", "set", byte('a')},
		},
		{
			name:            "Remove not exist values",
			values:          []interface{}{1, 2, 3, 4.98, "remove", "set", byte('a')},
			expRemoveValues: []interface{}{100, 22, 1.98, "str", byte('a')},
			remainingValues: []interface{}{1, 2, 3, 4.98, "remove", "set"},
			expSize:         6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)

			for _, v := range tc.expRemoveValues {
				s.Remove(v)
			}

			size := s.Size()
			if size != tc.expSize {
				t.Errorf("expected size %v, actual size %v", tc.expSize, size)
			}
			for _, v := range tc.remainingValues {
				if _, ok := s.set[v]; !ok {
					t.Errorf("expected %v, but nof found in set", v)
				}
			}
		})
	}
}

func TestThreadSafeSet_Contains(t *testing.T) {
	testCases := []struct {
		name        string
		values      []interface{}
		checkValues []interface{}
		exist       []bool
	}{
		{
			name:        "Check in empty set",
			checkValues: []interface{}{1, 2, 3.4, "str", true},
			exist:       []bool{false, false, false, false, false},
		},
		{
			name:        "Check existing and non existing values",
			values:      []interface{}{1, 2, 3, true, 4.5, "str", byte('b')},
			checkValues: []interface{}{1, 2, 5, false, "str", true},
			exist:       []bool{true, true, false, false, true, true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)

			for i, v := range tc.checkValues {
				if s.Contains(v) != tc.exist[i] {
					t.Errorf("value: %v\nexpected %v, actual %v", v, tc.exist[i], s.Contains(v))
				}
			}
		})
	}
}

func TestThreadSafeSet_Size(t *testing.T) {
	testCases := []struct {
		name         string
		values       []interface{}
		removeValues []interface{}
		expSize      uint
	}{
		{
			name: "Empty set",
		},
		{
			name:    "Single value set",
			values:  []interface{}{1},
			expSize: 1,
		},
		{
			name:    "Multiple value set",
			values:  []interface{}{1, 2, 3.45, "str", true, 'c'},
			expSize: 6,
		},
		{
			name:         "Add and remove multiple values",
			values:       []interface{}{1, 2, 3, 4.564, "str", false},
			removeValues: []interface{}{1, "str", false},
			expSize:      3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)

			for _, v := range tc.removeValues {
				s.Remove(v)
			}

			size := s.Size()
			if size != tc.expSize {
				t.Errorf("expected size %v, actual size %v", tc.expSize, size)
			}
		})
	}
}

func TestThreadSafeSet_Pop(t *testing.T) {
	testCases := []struct {
		name    string
		values  []interface{}
		expSize uint
	}{
		{
			name: "Pop from empty set",
		},
		{
			name:    "Pop from a non-empty set",
			values:  []interface{}{1, 2, 3, 4.53, "str", "set", true, 'b'},
			expSize: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)

			val := s.Pop()
			if s.Size() == 0 && val != nil { // empty set
				t.Errorf("expected popped value is nil, actual %v", val)
			} else if s.Size() > 0 && !s.Contains(val) { // non-empty set
				t.Errorf("set should contain %v, but not contain", val)
			}
		})
	}
}

func TestThreadSafeSet_Clear(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
	}{
		{
			name: "Clear empty set",
		},
		{
			name:   "Clear non-empty set",
			values: []interface{}{1, 2, 3, 4.542, "test", true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)
			s.Clear()
			size := s.Size()
			if size != 0 {
				t.Errorf("expected size is 0, actual size %v", size)
			}
		})
	}
}

func TestThreadSafeSet_Empty(t *testing.T) {
	testCases := []struct {
		name         string
		values       []interface{}
		removeValues []interface{}
		empty        bool
	}{
		{
			name:  "Check empty set",
			empty: true,
		},
		{
			name:   "Check non-empty set",
			values: []interface{}{1, 2, 3, 4.12, true, "test"},
			empty:  false,
		},
		{
			name:         "Firstly fill, then clear",
			values:       []interface{}{1, 2, 3, 4.12, true, "test"},
			removeValues: []interface{}{1, 2, 3, 4.12, true, "test"},
			empty:        true,
		},
		{
			name:         "Firstly fill, then remove some elements",
			values:       []interface{}{1, 2, 3, 4.12, true, "test"},
			removeValues: []interface{}{1, 2, 3, true},
			empty:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)
			for _, val := range tc.removeValues {
				s.Remove(val)
			}
			empty := s.Empty()
			if empty != tc.empty {
				t.Errorf("expected %v, actual %v", tc.empty, empty)
			}
		})
	}
}

func TestThreadSafeSet_Slice(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
	}{
		{
			name: "Empty set",
		},
		{
			name:   "Non-empty set slice",
			values: []interface{}{1, 2, 3, 4.123, "test", true, 'b'},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newThreadSafeSet()
			s.Append(tc.values...)
			setSlice := s.Slice()
			if len(setSlice) != len(tc.values) {
				t.Errorf("expected size %v, actual size %v", len(tc.values), len(setSlice))
			}
			for _, v := range setSlice {
				if !s.Contains(v) {
					t.Errorf("%v should be in set, but not found", v)
				}
			}
		})
	}
}

func TestThreadSafeSet_Union(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet  Set
	}{
		{
			name:   "Both empty sets",
			expSet: New(ThreadSafe),
		},
		{
			name:    "First set is empty",
			values2: []interface{}{1, 2, 3, 4, "test", true},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				3:      setVal,
				4:      setVal,
				"test": setVal,
				true:   setVal,
			}},
		},
		{
			name:    "Second set is empty",
			values1: []interface{}{1, 2, 3, 4, "test", true},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				3:      setVal,
				4:      setVal,
				"test": setVal,
				true:   setVal,
			}},
		},
		{
			name:    "Both sets are not empty",
			values1: []interface{}{1, 2, 3, 4, 5.12, "test", false},
			values2: []interface{}{1.23, "union", true},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:       setVal,
				2:       setVal,
				3:       setVal,
				4:       setVal,
				5.12:    setVal,
				1.23:    setVal,
				"test":  setVal,
				"union": setVal,
				true:    setVal,
				false:   setVal,
			}},
		},
		{
			name:    "Duplicate values",
			values1: []interface{}{1, 2, 3, "test", false},
			values2: []interface{}{1, 2, "test", true},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				3:      setVal,
				"test": setVal,
				true:   setVal,
				false:  setVal,
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)
			unionSet := s1.Union(s2)

			ts := tc.expSet.(*ThreadSafeSet)
			if ts.Size() != unionSet.Size() {
				t.Errorf("expected size %v, actual size %v", ts.Size(), unionSet.Size())
			}
			for val := range ts.set {
				if !unionSet.Contains(val) {
					t.Errorf("expected %v, but not exists in union set", val)
				}
			}
		})
	}
}

func TestThreadSafeSet_Intersection(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet  Set
	}{
		{
			name:   "Both empty sets",
			expSet: New(ThreadSafe),
		},
		{
			name:    "First set is empty",
			values2: []interface{}{1, 2, 3, 4.12, "test", true},
			expSet:  New(ThreadSafe),
		},
		{
			name:    "Second set is empty",
			values1: []interface{}{1, 2, 3, 4.12, "test", true},
			expSet:  New(ThreadSafe),
		},
		{
			name:    "Both sets are not empty",
			values1: []interface{}{1, 2, 3.12, "test", false},
			values2: []interface{}{1, 2, "test", 'b'},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				"test": setVal,
			}},
		},
		{
			name:    "No intersection",
			values1: []interface{}{1, 2, 3, false, "test"},
			values2: []interface{}{3.21, true, "set"},
			expSet:  New(ThreadSafe),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)
			intersectionSet := s1.Intersection(s2)

			if intersectionSet.Size() != tc.expSet.Size() {
				t.Errorf("expected size %v, actual size %v", tc.expSet.Size(), intersectionSet.Size())
			}

			ts := tc.expSet.(*ThreadSafeSet)
			for val := range ts.set {
				if !intersectionSet.Contains(val) {
					t.Errorf("expected %v, but not exists in intersection set", val)
				}
			}
		})
	}
}

func TestThreadSafeSet_Difference(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet1 Set
		expSet2 Set
	}{
		{
			name:    "Both empty sets",
			expSet1: New(ThreadSafe),
			expSet2: New(ThreadSafe),
		},
		{
			name:    "First set is empty",
			values2: []interface{}{1, 2, 3, 4.12, "test"},
			expSet1: New(ThreadSafe),
			expSet2: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				3:      setVal,
				4.12:   setVal,
				"test": setVal,
			}},
		},
		{
			name:    "Second set is empty",
			values1: []interface{}{1, 2, 3, 4.12, "test"},
			expSet1: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				3:      setVal,
				4.12:   setVal,
				"test": setVal,
			}},
			expSet2: New(ThreadSafe),
		},
		{
			name:    "Both sets are not empty",
			values1: []interface{}{1, 2, 3, 3.12, "test", false},
			values2: []interface{}{1, 2, 3.3, "test", "set", true},
			expSet1: &ThreadSafeSet{set: map[interface{}]struct{}{
				3:     setVal,
				3.12:  setVal,
				false: setVal,
			}},
			expSet2: &ThreadSafeSet{set: map[interface{}]struct{}{
				3.3:   setVal,
				"set": setVal,
				true:  setVal,
			}},
		},
		{
			name:    "No difference",
			values1: []interface{}{1, 2, 3.3, "test", "set", true},
			values2: []interface{}{1, 2, 3.3, "test", "set", true},
			expSet1: New(ThreadSafe),
			expSet2: New(ThreadSafe),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)

			diffSet1 := s1.Difference(s2)
			if diffSet1.Size() != tc.expSet1.Size() {
				t.Errorf("expected size %v, actual size %v", tc.expSet1.Size(), diffSet1.Size())
			}
			ts1 := tc.expSet1.(*ThreadSafeSet)
			for val := range ts1.set {
				if !diffSet1.Contains(val) {
					t.Errorf("expected %v, but not exists in difference set-1", val)
				}
			}

			diffSet2 := s2.Difference(s1)
			if diffSet2.Size() != tc.expSet2.Size() {
				t.Errorf("expected size %v, actual size %v", tc.expSet2.Size(), diffSet2.Size())
			}
			ts2 := tc.expSet2.(*ThreadSafeSet)
			for val := range ts2.set {
				if !diffSet2.Contains(val) {
					t.Errorf("expected %v, bot not exists in difference set-2", val)
				}
			}
		})
	}
}

func TestThreadSafeSet_IsSubset(t *testing.T) {
	testCases := []struct {
		name     string
		values1  []interface{}
		values2  []interface{}
		isSubset bool
	}{
		{
			name:     "Both empty sets",
			isSubset: true,
		},
		{
			name:     "First set is empty",
			values2:  []interface{}{1, 2, 3, 4, true, "test"},
			isSubset: true,
		},
		{
			name:     "Second set is empty",
			values1:  []interface{}{1, 2, 3, 4, true, "test"},
			isSubset: false,
		},
		{
			name:     "First set is subset of the second set",
			values1:  []interface{}{1, 2, "test"},
			values2:  []interface{}{1, 2, 5.12, "test", true},
			isSubset: true,
		},
		{
			name:     "First set is not subset of the second set",
			values1:  []interface{}{1, 2, 3, 4.23},
			values2:  []interface{}{1, 2, 3, 5.56, "test"},
			isSubset: false,
		},
		{
			name:     "Equal sets",
			values1:  []interface{}{1, 2, 3},
			values2:  []interface{}{1, 2, 3},
			isSubset: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)

			isSubset := s1.IsSubset(s2)
			if isSubset != tc.isSubset {
				t.Errorf("expected %v, actual %v", tc.isSubset, isSubset)
			}
		})
	}
}

func TestThreadSafeSet_IsSuperset(t *testing.T) {
	testCases := []struct {
		name       string
		values1    []interface{}
		values2    []interface{}
		isSuperset bool
	}{
		{
			name:       "Both empty sets",
			isSuperset: true,
		},
		{
			name:       "First set is empty",
			values2:    []interface{}{1, 2, 3, "test", 3.2, true},
			isSuperset: false,
		},
		{
			name:       "Second set is empty",
			values1:    []interface{}{1, 2, 3.21, "test", false},
			isSuperset: true,
		},
		{
			name:       "First set is superset of second set",
			values1:    []interface{}{1, 2, 3, 4.123, "test", false},
			values2:    []interface{}{1, 2, "test"},
			isSuperset: true,
		},
		{
			name:       "First set is not superset of second set",
			values1:    []interface{}{1, 2, 3, "test", false},
			values2:    []interface{}{1, 2, "set"},
			isSuperset: false,
		},
		{
			name:       "Equal sets",
			values1:    []interface{}{1, 2, 3, "test", false},
			values2:    []interface{}{1, 2, 3, "test", false},
			isSuperset: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)

			isSuperset := s1.IsSuperset(s2)
			if isSuperset != tc.isSuperset {
				t.Errorf("expected %v, actual %v", tc.isSuperset, isSuperset)
			}
		})
	}
}

func TestThreadSafeSet_IsDisjoint(t *testing.T) {
	testCases := []struct {
		name       string
		values1    []interface{}
		values2    []interface{}
		isDisjoint bool
	}{
		{
			name:       "Both empty sets",
			isDisjoint: true,
		},
		{
			name:       "First set is empty",
			values1:    []interface{}{1, 2, 3, 4, true},
			isDisjoint: true,
		},
		{
			name:       "Second set is empty",
			values2:    []interface{}{1, 2, 3, "test", false},
			isDisjoint: true,
		},
		{
			name:       "Disjoint sets",
			values1:    []interface{}{1, 2, "test", true},
			values2:    []interface{}{4, 2.34, "disjoint", false},
			isDisjoint: true,
		},
		{
			name:       "Not disjoint sets",
			values1:    []interface{}{1, 2, "test", false},
			values2:    []interface{}{1, true},
			isDisjoint: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)

			isDisjoint := s1.IsDisjoint(s2)
			if isDisjoint != tc.isDisjoint {
				t.Errorf("expected %v, actual %v", tc.isDisjoint, isDisjoint)
			}
		})
	}
}

func TestThreadSafeSet_Equal(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		equal   bool
	}{
		{
			name:  "Empty sets",
			equal: true,
		},
		{
			name:    "Sizes are different",
			values1: []interface{}{1, 2, "test"},
			values2: []interface{}{1, 2, 3, "test", true},
			equal:   false,
		},
		{
			name:    "Not equal sets",
			values1: []interface{}{1, 2, 3, "test", false},
			values2: []interface{}{1, 2, "test", true, false},
			equal:   false,
		},
		{
			name:    "Equal sets",
			values1: []interface{}{1, 2, 3, "test", true},
			values2: []interface{}{1, 2, 3, "test", true},
			equal:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)

			equal := s1.Equal(s2)
			if equal != tc.equal {
				t.Errorf("expected %v, actual %v", tc.equal, equal)
			}
		})
	}
}

func TestThreadSafeSet_SymmetricDifference(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet  Set
	}{
		{
			name:   "Empty sets",
			expSet: New(ThreadSafe),
		},
		{
			name:    "First set is empty",
			values2: []interface{}{1, 2, "test"},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				"test": setVal,
			}},
		},
		{
			name:    "Equal sets",
			values1: []interface{}{1, 2, "test", 32.12, false},
			values2: []interface{}{1, 2, "test", 32.12, false},
			expSet:  New(ThreadSafe),
		},
		{
			name:    "All distinct item sets",
			values1: []interface{}{1, 2, 3, 4},
			values2: []interface{}{5, 6, 7, 8},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1: setVal,
				2: setVal,
				3: setVal,
				4: setVal,
				5: setVal,
				6: setVal,
				7: setVal,
				8: setVal,
			}},
		},
		{
			name:    "Symmetric difference sets",
			values1: []interface{}{1, 2, "test", false},
			values2: []interface{}{2, "set", true, false},
			expSet: &ThreadSafeSet{set: map[interface{}]struct{}{
				1:      setVal,
				true:   setVal,
				"test": setVal,
				"set":  setVal,
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s1 := newThreadSafeSet()
			s2 := newThreadSafeSet()
			s1.Append(tc.values1...)
			s2.Append(tc.values2...)

			symmetricDiffSet := s1.SymmetricDifference(s2)
			if symmetricDiffSet.Size() != tc.expSet.Size() {
				t.Errorf("expected size %v, actual set size %v", tc.expSet.Size(), symmetricDiffSet.Size())
			}

			sds := symmetricDiffSet.(*ThreadSafeSet)
			for val := range sds.set {
				if !tc.expSet.Contains(val) {
					t.Errorf("expected %v in expected set, but not exists", val)
				}
			}

			ts := tc.expSet.(*ThreadSafeSet)
			for val := range ts.set {
				if !sds.Contains(val) {
					t.Errorf("expected %v in symmetric difference set, but not exists", val)
				}
			}
		})
	}
}
