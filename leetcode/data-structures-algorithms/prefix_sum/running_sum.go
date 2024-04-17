package prefix_sum

//Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//Return the running sum of nums.
//Example 1:
//
//Input: nums = [1,2,3,4]
//Output: [1,3,6,10]
//Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

func runningSum(nums []int) []int {
	// create a slice the size of the nums coming in to save mem
	ans := make([]int, len(nums))

	// pull out the first value as it will always be equal to idx 0
	ans[0] = nums[0]

	// now loop over the remaining nums starting at 1, and add in the last value in ans
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i] + ans[i-1]
	}
	return ans
}
