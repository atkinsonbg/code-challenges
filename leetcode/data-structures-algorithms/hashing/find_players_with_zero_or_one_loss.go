package hashing

func findWinners(matches [][]int) [][]int {
	m := make(map[int]int)

	for i := range matches {
		winner := matches[i][0]
		loser := matches[i][1]

		m[winner] = m[winner]
		m[loser] = m[loser] + 1

	}

	noLosses := []int{}
	oneLoss := []int{}

	for k, v := range m {
		if v == 0 {
			noLosses = append(noLosses, k)
		}

		if v == 1 {
			oneLoss = append(oneLoss, k)
		}
	}

	return [][]int{noLosses, oneLoss}
}
