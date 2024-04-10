package sliding_window

//You are given an integer array nums consisting of n elements, and an integer k.
//Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted.
//
//Example 1:
//Input: nums = [1,12,-5,-6,50,3], k = 4
//Output: 12.75000
//Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
//
//Example 2:
//Input: nums = [5], k = 1
//Output: 5.00000

func maxSubArray(nums []int, k int) float64 {
	curr := 0
	var maxAvg float64

	// do a first pass loop for the first window
	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	// get the current max average
	maxAvg = float64(curr) / float64(k)

	// now loop over the rest of the array, i == k
	for i := k; i < len(nums); i++ {
		// add the value of the current element nums[i] and subtract the value of the element k positions behind it nums[i-k].
		// This maintains the sum of the current window of size k. By adding the new element and subtracting the oldest element
		// in the window, the loop efficiently maintains the sum of the last k elements without having to sum them all in each iteration.
		curr += nums[i] - nums[i-k]

		// if the curr > ans save it off
		currAvg := float64(curr) / float64(k)

		// if this current avg is bigger save it
		if currAvg > maxAvg {
			maxAvg = currAvg
		}
	}

	return maxAvg
}
