package set

import "testing"

func TestThreadSafeSet_Add(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		expLen uint
		expSet map[interface{}]struct{}
	}{
		{
			name:   "Add int value",
			values: []interface{}{100},
			expLen: 1,
			expSet: map[interface{}]struct{}{
				100: setVal,
			},
		},
		{
			name:   "Add float value",
			values: []interface{}{12.332},
			expLen: 1,
			expSet: map[interface{}]struct{}{
				12.332: setVal,
			},
		},
		{
			name:   "Add bool value",
			values: []interface{}{true},
			expLen: 1,
			expSet: map[interface{}]struct{}{
				true: setVal,
			},
		},
		{
			name:   "Add byte value",
			values: []interface{}{byte('e')},
			expLen: 1,
			expSet: map[interface{}]struct{}{
				byte('e'): setVal,
			},
		},
		{
			name:   "Add string value",
			values: []interface{}{"str-key"},
			expLen: 1,
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
			expLen: 5,
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
			if s.Size() != tc.expLen {
				t.Errorf("expected length %v, actual length %v", tc.expLen, s.Size())
			}
			for v := range s.set {
				if _, ok := tc.expSet[v]; !ok {
					t.Errorf("value %v is expected, but not exist", v)
				}
			}
		})
	}
}
