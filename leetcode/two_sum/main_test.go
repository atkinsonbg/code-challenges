package two_sum

import (
	"reflect"
	"testing"
)

func Test1_two_sum(t *testing.T) {

	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test2_two_sum(t *testing.T) {

	got := twoSum([]int{3, 2, 4}, 6)
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test3_two_sum(t *testing.T) {

	got := twoSum([]int{3, 3}, 6)
	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test1_two_sum_fail(t *testing.T) {

	got := twoSum([]int{3, 3}, 70)
	want := []int{0, 1}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test4_two_sum(t *testing.T) {

	got := twoSumBruteForce([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test5_two_sum(t *testing.T) {

	got := twoSumBruteForce([]int{3, 2, 4}, 6)
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test6_two_sum(t *testing.T) {

	got := twoSumBruteForce([]int{3, 3}, 6)
	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test2_two_sum_fail(t *testing.T) {

	got := twoSumBruteForce([]int{3, 3}, 70)
	want := []int{0, 1}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{3, 3}, 6)
	}
	b.ReportAllocs()
}

func BenchmarkTwoSumBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumBruteForce([]int{3, 3}, 6)
	}
	b.ReportAllocs()
}
