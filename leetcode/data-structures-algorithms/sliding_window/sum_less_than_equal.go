package sliding_window

//Given an array of positive integers nums and an integer k, return the number of subarrays where the product of all the elements in the subarray is strictly less than k.
//
//For example, given the input nums = [10, 5, 2, 6], k = 100, the answer is 8. The subarrays with products less than k are:
//
//[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]

func SubarrayProductLessThanK(nums []int, k int) int {
	// init all our ans and left index to 0
	ans, left := 0, 0

	// init curr to 1 as on the first pass we multiply by 1 to get the first single digit subarray
	curr := 1

	// range over the nums with the right index
	for right := range len(nums) {

		// find the product of the first index which is nums[idx] * 1
		curr *= nums[right]

		// we get in here if the curr ever goes above 100
		for curr >= k {
			// first divide by the current left index, which undoes one multiplication operation
			curr = curr / nums[left]

			// increment the left index
			left++
		}

		// subtract right index from left to give us the subarray length
		ans += right - left + 1
	}

	return ans
}

//if k <= 1:
//return 0
//
//ans = left = 0
//curr = 1
//
//for right in range(len(nums)):
//curr *= nums[right]
//while curr >= k:
//curr //= nums[left]
//left += 1
//
//ans += right - left + 1
//
//return ans
