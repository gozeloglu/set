package set

import "testing"

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
