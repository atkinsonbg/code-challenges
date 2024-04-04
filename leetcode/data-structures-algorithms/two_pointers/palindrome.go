package two_pointers

func checkIfPalindrome(s string) bool {
	// start index at 0
	startIdx := 0

	// endIdx is len - 1 of the string
	endIdx := len(s) - 1

	// while the startIdx is less than the endIdx, this will catch middle chars in an odd length
	for startIdx < endIdx {
		// if the startIdx char != endIdx char, this cannot be a palidrome, return false
		if s[startIdx] != s[endIdx] {
			return false
		}

		// update the indexes
		startIdx++
		endIdx--
	}

	// return true if we made it here
	return true
}
