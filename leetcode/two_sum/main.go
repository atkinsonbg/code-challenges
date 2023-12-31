package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if _, seen := m[c]; seen {
			return []int{m[c], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func twoSumBruteForce(nums []int, target int) []int {
	var r []int
	for i1, n1 := range nums {
		for i2, n2 := range nums {
			if i1 != i2 {
				if n1+n2 == target {
					return append(r, i1, i2)
				}
			}
		}
	}
	return r
}
