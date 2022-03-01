package set

import "testing"

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
