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
			t.Log(tc.set.set)
			value := tc.set.Pop()
			t.Log(value)
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
