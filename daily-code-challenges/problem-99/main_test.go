package problem_99

import (
	"testing"
)

func Test1_99(t *testing.T) {

	got := FindLength([]int{100, 4, 200, 1, 3, 2})
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test2_99(t *testing.T) {

	got := FindLength([]int{4, 1, 3, 2})
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test3_99(t *testing.T) {

	got := FindLength([]int{4, 1, 5, 22, 89, 6})
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func BenchmarkFindLength_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindLength([]int{4, 1, 5, 22, 89, 6})
	}
}

func BenchmarkFindLength_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindLength([]int{55, 23, 12, 90, 99, 123, 124, 167, 1789, 15676, 4, 1, 5, 22, 89, 6})
	}
}
