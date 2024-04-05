package two_pointers

//Example 3: Given two sorted integer arrays arr1 and arr2, return a new array that combines both of them and is also sorted.

func combineArrays(arrOne []int, arrTwo []int) []int {
	// get the length of each array
	arrOneLen := len(arrOne)
	arrTwoLen := len(arrTwo)

	// create a new array to return, and set its capacity to the length of both arrays
	res := make([]int, arrOneLen+arrTwoLen)

	// create indexes for each array pointer
	arrOneStartIdx := 0
	arrTwoStartIdx := 0

	// using a loop index to insert into the new array
	loopIdx := 0

	// first we loop over both arrays at the same time, checking to see if we've exhausted all the elements in either
	for (arrOneStartIdx < arrOneLen) && (arrTwoStartIdx < arrTwoLen) {

		// if the first element of arr1 is greater than arr2, we save off arr2's value and increase its index
		if arrOne[arrOneStartIdx] > arrTwo[arrTwoStartIdx] {
			res[loopIdx] = arrTwo[arrTwoStartIdx]
			arrTwoStartIdx++
		} else {
			// otherwise we save off arr1's value and increase its index
			res[loopIdx] = arrOne[arrOneStartIdx]
			arrOneStartIdx++
		}

		// we up the loopIdx on each pass
		loopIdx++
	}

	// if we made it here, one of the arrays has been fully traversed, now we just insert the remaining items from the other
	// we don't know which one, so we just code for both to be safe
	for arrOneStartIdx < arrOneLen {
		res[loopIdx] = arrOne[arrOneStartIdx]
		arrOneStartIdx++
		loopIdx++
	}

	for arrTwoStartIdx < arrTwoLen {
		res[loopIdx] = arrTwo[arrTwoStartIdx]
		arrTwoStartIdx++
		loopIdx++
	}

	return res
}
