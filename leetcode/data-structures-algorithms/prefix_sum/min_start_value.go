package prefix_sum

func minStartValue(nums []int) int {
	startValue := 4
	valueCount := 0
	curr := 0

	// we'll keep a running count of how many times we looped and the curr is not 0
	for startValue > valueCount {

		// save off the first item in the array
		curr = startValue + nums[0]

		// loop over the other items starting at 1
		for i := 1; i < len(nums); i++ {

			// add the curr plus the item at current idx
			curr = curr + nums[i]

			// if we fall below 1, up the startValue and break as we need to start over
			if curr < 1 {
				startValue++
				break
			}

			// up the valueCount to compare against startValue
			valueCount++
		}
	}

	return startValue
}
