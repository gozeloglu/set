package set

import (
	"testing"
)

func TestSet_Add(t *testing.T) {
	testCases := []struct {
		name   string
		set    *Set
		val    interface{}
		expLen uint
		expSet map[interface{}]struct{}
	}{
		{
			name:   "Add a value from scratch",
			set:    New(),
			val:    "first",
			expLen: 1,
			expSet: map[interface{}]struct{}{
				"first": setVal,
			},
		},
		{
			name:   "Add int value",
			set:    New(),
			val:    12,
			expLen: 1,
			expSet: map[interface{}]struct{}{
				12: setVal,
			},
		},
		{
			name:   "Add float value",
			set:    New(),
			val:    12.3,
			expLen: 1,
			expSet: map[interface{}]struct{}{
				12.3: setVal,
			},
		},
		{
			name:   "Add bool value",
			set:    New(),
			val:    true,
			expLen: 1,
			expSet: map[interface{}]struct{}{
				true: setVal,
			},
		},
		{
			name:   "Add byte value",
			set:    New(),
			val:    byte('b'),
			expLen: 1,
			expSet: map[interface{}]struct{}{
				byte('b'): setVal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Add(tc.val)
			if tc.expLen != uint(len(tc.set.set)) {
				t.Errorf("expected length: %v, actual length: %v", tc.expLen, uint(len(tc.set.set)))
			}
			for v := range tc.set.set {
				if _, ok := tc.expSet[v]; !ok {
					t.Errorf("expected value: %v, but not found in expected set", v)
				}
			}
			for v := range tc.expSet {
				if _, ok := tc.set.set[v]; !ok {
					t.Errorf("expected value: %v, but not found in test case set", v)
				}
			}
		})
	}
}

func TestSet_Append(t *testing.T) {
	testCases := []struct {
		name   string
		set    *Set
		values []interface{}
	}{
		{
			name:   "Append single value",
			set:    New(),
			values: []interface{}{"test_value"},
		},
		{
			name:   "Append multiple values",
			set:    New(),
			values: []interface{}{"test_value1", "test_value2"},
		},
		{
			name:   "Append nothing",
			set:    New(),
			values: []interface{}{},
		},
		{
			name:   "Append different data types",
			set:    New(),
			values: []interface{}{"str", 12, true, 32.4, uint16(45), byte('a')},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Append(tc.values...)
			for _, val := range tc.values {
				if _, ok := tc.set.set[val]; !ok {
					t.Errorf("expected %v, but not found in set", val)
				}
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	testCases := []struct {
		name            string
		set             *Set
		values          []interface{}
		expRemoveValues []interface{}
		remainingValues []interface{}
		expLen          int
	}{
		{
			name:            "Remove from empty set",
			set:             New(),
			values:          []interface{}{},
			expRemoveValues: []interface{}{"test", 12},
			expLen:          0,
		},
		{
			name:            "Remove from 1-length set",
			set:             New(),
			values:          []interface{}{"test_val"},
			expRemoveValues: []interface{}{"test_val"},
			expLen:          0,
		},
		{
			name:            "Remove from 3-length set",
			set:             New(),
			values:          []interface{}{"val", 12, true},
			expRemoveValues: []interface{}{12},
			remainingValues: []interface{}{"val", true},
			expLen:          2,
		},
		{
			name:            "Remove multiple values",
			set:             New(),
			values:          []interface{}{"test", 12, 43.2, true, byte('a')},
			expRemoveValues: []interface{}{"test", true, 43.2},
			remainingValues: []interface{}{12, byte('a')},
			expLen:          2,
		},
		{
			name:            "Remove not exist value",
			set:             New(),
			values:          []interface{}{"test_val"},
			expRemoveValues: []interface{}{"test_key"},
			remainingValues: []interface{}{"test_val"},
			expLen:          1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Append(tc.values...)
			for _, v := range tc.expRemoveValues {
				tc.set.Remove(v)
			}

			actualLen := len(tc.set.set)
			if actualLen != tc.expLen {
				t.Errorf("expected length %v, actual length %v", tc.expLen, actualLen)
			}
			for _, v := range tc.remainingValues {
				if _, ok := tc.set.set[v]; !ok {
					t.Errorf("expected value %v, but not found in set", v)
				}
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	testCases := []struct {
		name     string
		set      *Set
		values   []interface{}
		checkVal interface{}
		exist    bool
	}{
		{
			name:     "Check exist value",
			set:      New(),
			values:   []interface{}{"test"},
			checkVal: "test",
			exist:    true,
		},
		{
			name:     "Check non-exist value",
			set:      New(),
			values:   []interface{}{"test"},
			checkVal: "value",
			exist:    false,
		},
		{
			name:     "Check empty set",
			set:      New(),
			checkVal: "test",
			exist:    false,
		},
		{
			name:     "Check integer - exist",
			set:      New(),
			values:   []interface{}{120},
			checkVal: 120,
			exist:    true,
		},
		{
			name:     "Check integer - not exist",
			set:      New(),
			values:   []interface{}{120},
			checkVal: 200,
			exist:    false,
		},
		{
			name:     "Check float - exist",
			set:      New(),
			values:   []interface{}{12.98},
			checkVal: 12.98,
			exist:    true,
		},
		{
			name:     "Check boolean",
			set:      New(),
			values:   []interface{}{false},
			checkVal: false,
			exist:    true,
		},
		{
			name:     "Check byte",
			set:      New(),
			values:   []interface{}{byte('a')},
			checkVal: byte('a'),
			exist:    true,
		},
		{
			name:     "Check value from multiple set",
			set:      New(),
			values:   []interface{}{120, 32.123, "test", false, byte('a')},
			checkVal: 120,
			exist:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Append(tc.values...)
			exist := tc.set.Contains(tc.checkVal)
			if exist != tc.exist {
				t.Errorf("Value: %v \nexpected %v, actual %v", tc.checkVal, tc.exist, exist)
			}

		})
	}
}

func TestSet_Size(t *testing.T) {
	testCases := []struct {
		name         string
		set          *Set
		values       []interface{}
		removeValues []interface{}
		expSize      uint
	}{
		{
			name:    "Empty set",
			set:     New(),
			expSize: 0,
		},
		{
			name:    "Add value, check size",
			set:     New(),
			values:  []interface{}{"test"},
			expSize: 1,
		},
		{
			name:         "Add value, remove value, check size",
			set:          New(),
			values:       []interface{}{"test"},
			removeValues: []interface{}{"test"},
			expSize:      0,
		},
		{
			name:    "Add multiple values",
			set:     New(),
			values:  []interface{}{"test", 125, true, 64.23},
			expSize: 4,
		},
		{
			name:         "Add multiple values, remove multiple values",
			set:          New(),
			values:       []interface{}{"test", 125, true, 64.23},
			removeValues: []interface{}{125, true},
			expSize:      2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.values) > 0 {
				tc.set.Append(tc.values...)
			}
			if len(tc.removeValues) > 0 {
				for _, val := range tc.removeValues {
					tc.set.Remove(val)
				}
			}
			size := tc.set.Size()
			if size != tc.expSize {
				t.Errorf("expected size %v, actual size %v", tc.expSize, size)
			}
		})
	}
}

func TestSet_Pop(t *testing.T) {
	testCases := []struct {
		name    string
		set     *Set
		values  []interface{}
		isEmpty bool
	}{
		{
			name:    "Pop from empty set",
			set:     New(),
			isEmpty: true,
		},
		{
			name:   "Pop from single value set",
			set:    New(),
			values: []interface{}{"test"},
		},
		{
			name:   "Pop from multiple value set",
			set:    New(),
			values: []interface{}{"test", 123, true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.values != nil {
				tc.set.Append(tc.values...)
			}
			value := tc.set.Pop()
			for _, val := range tc.values {
				if val == value {
					return
				}
			}
			if tc.isEmpty && value != nil {
				t.Errorf("expected nil, actual %v", value)
			}
			if !tc.isEmpty {
				t.Errorf("%v not exist in set", value)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	testCases := []struct {
		name   string
		set    *Set
		values []interface{}
	}{
		{
			name:   "Clear empty set",
			set:    New(),
			values: []interface{}{},
		},
		{
			name:   "Clear single value set",
			set:    New(),
			values: []interface{}{"test"},
		},
		{
			name:   "Clear multiple value set",
			set:    New(),
			values: []interface{}{"test", 12, false, byte('w'), 43.10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Append(tc.values...)
			tc.set.Clear()
			size := tc.set.Size()
			if size != 0 {
				t.Errorf("expected size is 0, actual size is %v", size)
			}
			for _, val := range tc.values {
				if tc.set.Contains(val) {
					t.Errorf("%v should not exist in the set, but it exists", val)
				}
			}
		})
	}
}

func TestSet_Empty(t *testing.T) {
	testCases := []struct {
		name         string
		set          *Set
		values       []interface{}
		removeValues []interface{}
		empty        bool
	}{
		{
			name:   "Check empty set",
			set:    New(),
			values: []interface{}{},
			empty:  true,
		},
		{
			name:   "Check single value set",
			set:    New(),
			values: []interface{}{"test"},
			empty:  false,
		},
		{
			name:   "Check multiple value set",
			set:    New(),
			values: []interface{}{"test", 100, true, false, 98.4},
			empty:  false,
		},
		{
			name:         "Check firstly filled, then cleared set",
			set:          New(),
			values:       []interface{}{"test", 100, true, 76.34},
			removeValues: []interface{}{"test", 100, true, 76.34},
			empty:        true,
		},
		{
			name:         "Check firstly filled, then removed some elements set",
			set:          New(),
			values:       []interface{}{"test", 100, true, 76.34},
			removeValues: []interface{}{"test", 100},
			empty:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Append(tc.values...)
			if tc.removeValues != nil {
				for _, val := range tc.removeValues {
					tc.set.Remove(val)
				}
			}
			empty := tc.set.Empty()
			if empty != tc.empty {
				t.Errorf("expected %v, actual %v", tc.empty, empty)
			}

		})
	}
}

func TestSet_Slice(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		expLen int
	}{
		{
			name: "Empty set",
		},
		{
			name:   "Multiple value set",
			values: []interface{}{1, 2, 3, 4, 5},
			expLen: 5,
		},
		{
			name:   "Multiple value, different type set",
			values: []interface{}{1, 2, 3, "test", "array", true, false, 43.12},
			expLen: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := New()
			s.Append(tc.values...)
			setSlice := s.Slice()
			if len(setSlice) != tc.expLen {
				t.Errorf("expected length %v, actual length %v", tc.expLen, len(setSlice))
			}
			for i := range setSlice {
				if !s.Contains(setSlice[i]) {
					t.Errorf("%v should be in array, but not exist", setSlice[i])
				}
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet  *Set
	}{
		{
			name:   "Both empty set",
			expSet: New(),
		},
		{
			name:    "Set-1 is empty",
			values2: []interface{}{"test", 1, 120.32, true},
			expSet: &Set{set: map[interface{}]struct{}{
				"test": setVal,
				1:      setVal,
				120.32: setVal,
				true:   setVal,
			}},
		},
		{
			name:    "Set-2 is empty",
			values1: []interface{}{"test", 1, 120.32, true},
			expSet: &Set{set: map[interface{}]struct{}{
				"test": setVal,
				1:      setVal,
				120.32: setVal,
				true:   setVal,
			}},
		},
		{
			name:    "Both sets not empty",
			values1: []interface{}{"testStr", 32, false},
			values2: []interface{}{"anotherStr", 43, true},
			expSet: &Set{set: map[interface{}]struct{}{
				"testStr":    setVal,
				32:           setVal,
				false:        setVal,
				"anotherStr": setVal,
				43:           setVal,
				true:         setVal,
			}},
		},
		{
			name:    "Duplicate values",
			values1: []interface{}{"test", 100, true},
			values2: []interface{}{"test", 120, 100, true, false},
			expSet: &Set{set: map[interface{}]struct{}{
				"test": setVal,
				100:    setVal,
				true:   setVal,
				120:    setVal,
				false:  setVal,
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)
			set2 := New()
			set2.Append(tc.values2...)

			unionSet := set1.Union(set2)

			if unionSet.Size() != tc.expSet.Size() {
				t.Errorf("expected size %v, actual size %v", tc.expSet.Size(), unionSet.Size())
			}
			for val := range tc.expSet.set {
				if !unionSet.Contains(val) {
					t.Errorf("expected %v, but not exists in unionSet", val)
				}
			}

		})
	}
}

func TestSet_Intersection(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet  *Set
	}{
		{
			name:   "Both empty sets",
			expSet: New(),
		},
		{
			name:    "Set1 is empty",
			values2: []interface{}{1, 2, 3, 4},
			expSet:  New(),
		},
		{
			name:    "Set2 is empty",
			values1: []interface{}{1, 2, 3, 4},
			expSet:  New(),
		},
		{
			name:    "Both set is non-empty",
			values1: []interface{}{1, 2, 3, 4, "test", true, byte('w')},
			values2: []interface{}{1, 2, "test", false, true, byte('a')},
			expSet: &Set{set: map[interface{}]struct{}{
				1:      setVal,
				2:      setVal,
				"test": setVal,
				true:   setVal,
			}},
		},
		{
			name:    "No intersection",
			values2: []interface{}{1, 2, 3, 4, "value"},
			values1: []interface{}{0, "test", true},
			expSet:  New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)

			set2 := New()
			set2.Append(tc.values2...)

			intersectSet1 := set1.Intersection(set2)
			for k := range tc.expSet.set {
				if !intersectSet1.Contains(k) {
					t.Errorf("expected value %v, but not found in intersection set", k)
				}
			}
			intersectSet2 := set2.Intersection(set1)
			for k := range tc.expSet.set {
				if !intersectSet2.Contains(k) {
					t.Errorf("expected value %v, but not found in intersection set", k)
				}
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet1 *Set
		expSet2 *Set
	}{
		{
			name:    "Both empty sets",
			expSet1: New(),
			expSet2: New(),
		},
		{
			name:    "Set1 is empty",
			values2: []interface{}{1, 2, 3, 4, 5},
			expSet1: New(),
			expSet2: &Set{set: map[interface{}]struct{}{
				1: setVal,
				2: setVal,
				3: setVal,
				4: setVal,
				5: setVal,
			}},
		},
		{
			name:    "Set2 is empty",
			values1: []interface{}{1, 2, 3, 4, 5},
			expSet1: &Set{set: map[interface{}]struct{}{
				1: setVal,
				2: setVal,
				3: setVal,
				4: setVal,
				5: setVal,
			}},
			expSet2: New(),
		},
		{
			name:    "Both sets are non-empty",
			values1: []interface{}{1, 2, 3, 4, 5},
			values2: []interface{}{1, 2, 3, 10, 20, "test", true},
			expSet1: &Set{set: map[interface{}]struct{}{
				4: setVal,
				5: setVal,
			}},
			expSet2: &Set{map[interface{}]struct{}{
				10:     setVal,
				20:     setVal,
				"test": setVal,
				true:   setVal,
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)

			set2 := New()
			set2.Append(tc.values2...)

			diffSet1 := set1.Difference(set2)
			if diffSet1.Size() != tc.expSet1.Size() {
				t.Errorf("expected set size %v, actual set size %v", tc.expSet1.Size(), diffSet1.Size())
			}
			for val := range diffSet1.set {
				if !tc.expSet1.Contains(val) {
					t.Errorf("expected %v, but not found", val)
				}
			}

			diffSet2 := set2.Difference(set1)
			if diffSet2.Size() != tc.expSet2.Size() {
				t.Errorf("expected set size %v, actual set size %v", tc.expSet2.Size(), diffSet2.Size())
			}
			for val := range diffSet2.set {
				if !tc.expSet2.Contains(val) {
					t.Errorf("expected %v, but not found", val)
				}
			}
		})
	}
}

func TestSet_IsSubset(t *testing.T) {
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
			name:     "set1 is empty",
			values2:  []interface{}{1, 2, 3},
			isSubset: true,
		},
		{
			name:     "set2 is empty",
			values1:  []interface{}{1, 2, 3},
			isSubset: false,
		},
		{
			name:     "set1 is subset of set2",
			values1:  []interface{}{1, 2, 3, 4},
			values2:  []interface{}{1, 2, 3, 4, 5, 6},
			isSubset: true,
		},
		{
			name:     "set1 is not subset of set2",
			values1:  []interface{}{1, 2, 3},
			values2:  []interface{}{2, 3, 4},
			isSubset: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)

			set2 := New()
			set2.Append(tc.values2...)

			isSubset := set1.IsSubset(set2)
			if isSubset != tc.isSubset {
				t.Errorf("Expected %v, actual %v", tc.isSubset, isSubset)
			}
		})
	}
}

func TestSet_IsSuperset(t *testing.T) {
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
			name:       "set1 is empty",
			values2:    []interface{}{1, 2, 3, 4, 5},
			isSuperset: false,
		},
		{
			name:       "set2 is empty",
			values1:    []interface{}{1, 2, 3, 4, 5},
			isSuperset: true,
		},
		{
			name:       "set1 is superset of set2",
			values1:    []interface{}{1, 2, 3, 4, 5, 6, "test", true},
			values2:    []interface{}{1, 2, 3, true},
			isSuperset: true,
		},
		{
			name:       "set1 is not superset of set2",
			values1:    []interface{}{1, 2, 3, 4, 5, "test", true},
			values2:    []interface{}{1, 2, 3, false},
			isSuperset: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)

			set2 := New()
			set2.Append(tc.values2...)

			isSuperset := set1.IsSuperset(set2)
			if isSuperset != tc.isSuperset {
				t.Errorf("Expected %v, actual %v", tc.isSuperset, isSuperset)
			}
		})
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	testCases := []struct {
		name       string
		values1    []interface{}
		values2    []interface{}
		isDisjoint bool
	}{
		{
			name:       "Empty sets",
			isDisjoint: true,
		},
		{
			name:       "set1 is empty",
			values2:    []interface{}{1, 2, 3, 4},
			isDisjoint: true,
		},
		{
			name:       "set2 is empty",
			values1:    []interface{}{1, 2, 3, 4},
			isDisjoint: true,
		},
		{
			name:       "Disjoint sets",
			values1:    []interface{}{1, 2, 3, 4},
			values2:    []interface{}{5, 6, 7, 8},
			isDisjoint: true,
		},
		{
			name:       "Not disjoint sets",
			values1:    []interface{}{1, 2, 3, 4},
			values2:    []interface{}{1, 2, 5, 6, 7},
			isDisjoint: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)

			set2 := New()
			set2.Append(tc.values2...)

			isDisjoint := set1.IsDisjoint(set2)
			if isDisjoint != tc.isDisjoint {
				t.Errorf("Expected %v, actual %v", tc.isDisjoint, isDisjoint)
			}
		})
	}
}

func TestSet_Equal(t *testing.T) {
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
			values1: []interface{}{1, 2, 3, 4},
			values2: []interface{}{1, 2, 3, 4, 5},
			equal:   false,
		},
		{
			name:    "Not equal sets",
			values1: []interface{}{1, 2, 3, 4, 5},
			values2: []interface{}{1, 2, 3, 4, 6},
			equal:   false,
		},
		{
			name:    "Equal sets",
			values1: []interface{}{1, 2, 3, 4, 5},
			values2: []interface{}{1, 2, 3, 4, 5},
			equal:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)

			set2 := New()
			set2.Append(tc.values2...)

			equal := set1.Equal(set2)
			if equal != tc.equal {
				t.Errorf("Expected %v, actual %v", tc.equal, equal)
			}
		})
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	testCases := []struct {
		name    string
		values1 []interface{}
		values2 []interface{}
		expSet  *Set
	}{
		{
			name:   "Empty sets",
			expSet: New(),
		},
		{
			name:    "set1 is empty",
			values2: []interface{}{1, 2, 3, 4},
			expSet: &Set{set: map[interface{}]struct{}{
				1: setVal,
				2: setVal,
				3: setVal,
				4: setVal,
			}},
		},
		{
			name:    "Equal sets",
			values1: []interface{}{1, 2, 3, 4},
			values2: []interface{}{1, 2, 3, 4},
			expSet:  New(),
		},
		{
			name:    "All distinct item sets",
			values1: []interface{}{1, 2, 3, 4},
			values2: []interface{}{5, 6, 7, 8},
			expSet: &Set{set: map[interface{}]struct{}{
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
			values1: []interface{}{1, 2, 3, 4},
			values2: []interface{}{2, 3, 4},
			expSet: &Set{set: map[interface{}]struct{}{
				1: setVal,
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := New()
			set1.Append(tc.values1...)

			set2 := New()
			set2.Append(tc.values2...)

			symDiffSet := set1.SymmetricDifference(set2)
			if symDiffSet.Size() != tc.expSet.Size() {
				t.Errorf("Expected set size %v, actual set size %v", tc.expSet.Size(), symDiffSet.Size())
			}

			for val := range symDiffSet.set {
				if !tc.expSet.Contains(val) {
					t.Errorf("expected %v expSet, but not contains", val)
				}
			}

			for val := range tc.expSet.set {
				if !symDiffSet.Contains(val) {
					t.Errorf("expected %v in symDiffSet, but not contains", val)
				}
			}
		})
	}
}

var benchmarkData = []interface{}{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	true, false,
	"a", "b", "c", "d", "e", "f", "g", "h", "j", "k", "l", "m", "n",
	1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

func BenchmarkSet_Add(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkSet_Append(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Append(benchmarkData...)
	}
}

func BenchmarkSet_Remove(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Remove(i)
	}
}

func BenchmarkSet_Contains(b *testing.B) {
	s := New()
	idx := 0
	n := len(benchmarkData)
	s.Append(benchmarkData...)
	for i := 0; i < b.N; i++ {
		s.Contains(i)
		idx++
		if idx == n {
			idx = 0
		}
	}
}

func BenchmarkSet_Clear(b *testing.B) {
	s := New()
	s.Append(benchmarkData...)
	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}
