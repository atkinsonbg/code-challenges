package problem_99

import "sort"

//Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
//For example, given [100, 4, 200, 1, 3, 2], the longest consecutive element sequence is [1, 2, 3, 4].
// Return its length: 4.
//Your algorithm should run in O(n) complexity.

func FindLength(numbers []int) int {
	longestStreak := 0
	currentStreak := 0

	sort.Ints(numbers)

	for i, num := range numbers {
		// we just started, so this is the current streak
		if i == 0 {
			currentStreak++
			continue
		}

		// check if the current number, equals the previous number
		if num-1 == numbers[i-1] {
			currentStreak++
		}

		// if the current streak is longer, save it off
		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
	}

	return currentStreak
}
