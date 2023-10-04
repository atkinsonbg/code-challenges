package problem_84

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

func FindIslands(matrix [6][5]int) int {
	islandCount := 0
	for _, row := range matrix {
		rowSize := len(row)
		landCount := 0
		for _, cell := range row {
			if cell == 1 {
				landCount++
			}
		}
		if landCount > 1 && landCount != rowSize {
			islandCount++
		}
	}
	return islandCount
}
