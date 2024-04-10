package sliding_window

//Example 4: Given an integer array nums and an integer k, find the sum of the subarray with the largest sum whose length is k.
// nums := []int{3, -1, 4, 12, -8, 5, 6}
// k := 4

func sumOfSubarraysWithLargestSum(nums []int, k int) int {
	curr := 0

	// first pass for initial window, move from 0 to index at k
	for i := range k {
		curr += nums[i]
	}

	// save off this answer
	ans := curr

	// now loop over the rest of the array, i == k
	for i := k; i < len(nums); i++ {
		// find the current total
		curr += nums[i] - nums[i-k]

		// if the curr > ans save it off
		if curr > ans {
			ans = curr
		}
	}

	return ans
}
