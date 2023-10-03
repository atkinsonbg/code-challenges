package main

//This problem was asked by Facebook.
//
//Given three 32-bit integers x, y, and b, return x if b is 1 and y if b is 0, using
//only mathematical or bit operations. You can assume b can only be 1 or 0.

func FindB(x int, y int, b int) int {
	boolMap := map[bool]int{true: x, false: y}
	return boolMap[b == 1]
}
