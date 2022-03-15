package set

import "testing"

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
