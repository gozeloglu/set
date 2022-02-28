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
			set:    NewSet(),
			val:    "first",
			expLen: 1,
			expSet: map[interface{}]struct{}{
				"first": setVal,
			},
		},
		{
			name:   "Add int value",
			set:    NewSet(),
			val:    12,
			expLen: 1,
			expSet: map[interface{}]struct{}{
				12: setVal,
			},
		},
		{
			name:   "Add float value",
			set:    NewSet(),
			val:    12.3,
			expLen: 1,
			expSet: map[interface{}]struct{}{
				12.3: setVal,
			},
		},
		{
			name:   "Add bool value",
			set:    NewSet(),
			val:    true,
			expLen: 1,
			expSet: map[interface{}]struct{}{
				true: setVal,
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
