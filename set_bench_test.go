package set

import (
	"testing"
)

func BenchmarkSet_Add(b *testing.B) {
	s := newThreadUnsafeSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkSet_Append(b *testing.B) {
	s := newThreadUnsafeSet()
	for i := 0; i < b.N; i++ {
		s.Append(i, -1*i)
	}
}

func BenchmarkSet_Remove(b *testing.B) {
	b.StopTimer()
	s := newThreadUnsafeSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}
}

func BenchmarkSet_Contains(b *testing.B) {
	b.StopTimer()
	s := newThreadUnsafeSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(i)
	}
}
