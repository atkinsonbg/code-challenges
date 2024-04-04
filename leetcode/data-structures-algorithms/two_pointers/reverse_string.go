package two_pointers

//Write a function that reverses a string. The input string is given as an array of characters s.
//You must do this by modifying the input array in-place with O(1) extra memory.
//
//Example 1:
//Input: s = ["h","e","l","l","o"]
//Output: ["o","l","l","e","h"]
//Example 2:
//
//Input: s = ["H","a","n","n","a","h"]
//Output: ["h","a","n","n","a","H"]
//
//Constraints:
//
//1 <= s.length <= 105
//s[i] is a printable ascii character.

func reverseString(s []byte) {
	// left side index
	i := 0

	// get the length of the array to use to get right side index later
	l := len(s)

	// get the middle index
	m := l / 2

	// since we work both sides of the array, just stop once you make it to the middle
	for i < m {
		// j is the right side index calculated by subtracting 1 and the current i value
		j := l - i - 1

		// inplace replacement of the array values
		s[i], s[j] = s[j], s[i]

		// increment the index
		i++
	}
}
