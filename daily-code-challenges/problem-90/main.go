package problem_90

import (
	"math/rand"
	"time"
)

//Given an integer n and a list of integers l, write a function that randomly
//generates a number from 0 to n-1 that isn't in l (uniform).

func GenerateRandomNumber(n int, excludeNums []int) int {
	// create a new rand.Source based off time.Now
	seed := rand.NewSource(time.Now().UnixNano())

	// create a new rand based on that seed
	ran := rand.New(seed)

	// convert the int slice to a map of [int]bool
	exclude := map[int]bool{}
	for _, x := range excludeNums {
		exclude[x] = true
	}

	// loop until we find a random number that is not in the exclude list
	for {
		// gen a random number
		ranNum := ran.Intn(n)

		// check against the exclude map
		if !exclude[ranNum] {
			return ranNum
		}
	}
}
