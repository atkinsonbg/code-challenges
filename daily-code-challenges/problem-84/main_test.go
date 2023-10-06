package problem_84

import "testing"

//This problem was asked by Amazon.
//
//Given a matrix of 1s and 0s, return the number of "islands" in the matrix. A
//1 represents land and 0 represents water, so an island is a group of 1s
//that are neighboring whose perimeter is surrounded by water.
//
//For example, this matrix has 4 islands.
//
//1 0 0 0 0
//0 0 1 1 0
//0 1 1 0 0
//0 0 0 0 0
//1 1 0 0 1
//1 1 0 0 1

func Test1_84(t *testing.T) {

	m := [6][5]int{
		{1, 0, 0, 0, 0}, /*  initializers for row indexed by 0 */
		{0, 0, 1, 1, 0}, /*  initializers for row indexed by 1 */
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{1, 1, 0, 0, 1}, /*  initializers for row indexed by 2 */
	}

	got := FindIslands(m)
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func BenchmarkFindIslands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := [6][5]int{
			{1, 0, 0, 0, 0}, /*  initializers for row indexed by 0 */
			{0, 0, 1, 1, 0}, /*  initializers for row indexed by 1 */
			{0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{1, 1, 0, 0, 1},
			{1, 1, 0, 0, 1}, /*  initializers for row indexed by 2 */
		}
		FindIslands(m)
	}
}
