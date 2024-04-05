package two_pointers

//Example 2: Given a sorted array of unique integers and a target integer, return true if there exists a pair of numbers that sum to target, false otherwise. This problem is similar to Two Sum. (In Two Sum, the input is not sorted).
//
//For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13, return true because 4 + 9 = 13.

func isTargetInArray(nums []int, target int) bool {
	// set the start index
	startIdx := 0

	// set the end index
	endIdx := len(nums) - 1

	for startIdx < endIdx {

		// sum the numbers at the current indexes
		currentSum := nums[startIdx] + nums[endIdx]

		// if the current sum matches target, return true
		if currentSum == target {
			return true
		}

		// if not deteremine which index to modify
		if currentSum > target {
			// if the current sum is greater than the target, we can decrease the endIdx
			endIdx--
		} else {
			// if the current sum is less than the target, we can increase the startIdx
			startIdx++
		}
	}

	// we found no match, return false
	return false
}
