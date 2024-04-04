package two_pointers

//Example 1:
//
//Input: nums = [-4,-1,0,3,10]
//Output: [0,1,9,16,100]
//Explanation: After squaring, the array becomes [16,1,0,9,100].
//After sorting, it becomes [0,1,9,16,100].
//Example 2:
//
//Input: nums = [-7,-3,2,3,11]
//Output: [4,9,9,49,121]
//
//
//Constraints:
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums is sorted in non-decreasing order.

func sortedSquares(nums []int) []int {
	// first get the len of the array coming in, this is used later
	lenNums := len(nums)

	// create a new array to return, and set its capacity to save memory
	res := make([]int, lenNums)

	// pull the left and right indexes for inspecting values
	leftIdx := 0
	rightIdx := len(nums) - 1

	// loop it
	for i := 0; i < len(nums); i++ {

		// first square each left and right index which leaves us with positive only numbers
		leftNumSq := nums[leftIdx] * nums[leftIdx]
		rightNumSq := nums[rightIdx] * nums[rightIdx]

		// create a var to hold the value we want to insert to reduce the amount of code
		insertValue := 0

		// if the left side (negative side) is greater, we stuff that into the new array at the end
		if leftNumSq > rightNumSq {

			// we know the end since we pulled the len of the array at the start
			insertValue = leftNumSq

			// increase the left index since we don't need to inspect that one again
			leftIdx++
		} else {
			// if the right side (postive side) is greater, we stuff that into the new array at the end
			insertValue = rightNumSq

			// this time increase the right index as we don't need to look at it again
			rightIdx--
		}

		// insert the value we decided should go in
		res[lenNums-1] = insertValue

		// decrease the insert target index for the next pass
		lenNums--
	}

	return res
}
