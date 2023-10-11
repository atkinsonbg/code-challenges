package problem_90

import (
	"log"
	"testing"
)

func Test1_90(t *testing.T) {

	got := GenerateRandomNumber(10, []int{8, 3, 6})

	log.Println(got)

	if got == 8 || got == 3 || got == 6 {
		t.Errorf("got %q, wanted anything but that", got)
	}
}

func BenchmarkGenerateRandomNumber_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomNumber(10, []int{8, 3, 6})
	}
}

func BenchmarkGenerateRandomNumber_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomNumber(100, []int{8, 3, 6, 22, 56, 34, 90, 23, 33, 12, 11, 7, 99, 55, 43, 78})
	}
}
