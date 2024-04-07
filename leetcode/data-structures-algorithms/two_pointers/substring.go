package two_pointers

//In this problem, we need to check if the characters of s appear in the same order in t,
//with gaps allowed. For example, "ace" is a subsequence of "abcde" because "abcde"
//contains the letters "ace" in that same order - the fact that they aren't consecutive
//doesn't matter.

func findSubstring(word string, subWord string) bool {
	// count of how many chars we found in the string
	foundCount := 0

	// length of the subword, this is check if the number of found chars matches length of foundCount
	sLen := len(subWord)

	// length of the word we're searching in, used for the loop
	wLen := len(word)

	// starting indexes for each word
	wIdx, sIdx := 0, 0

	// loop until the word len is hit
	for wIdx < wLen {

		// check the current indexes on each word to see if they match
		if word[wIdx] == subWord[sIdx] {
			// if they match increase the subword index and the foundcount
			sIdx++
			foundCount++
		}

		// no matter what, we incease the word index to keep searching
		wIdx++
	}

	// return if the foundCount matches the subword length, which would mean we found all the chars we were looking for
	return foundCount == sLen
}
