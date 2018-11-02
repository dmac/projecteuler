package main

import (
	"fmt"
)

func main() {
	var longest, longestStart, length int64
	for i := 1; i < 1E6; i++ {
		length = collatzSequenceLength(int64(i))
		if length > longest {
			longest = length
			longestStart = int64(i)
		}
	}
	fmt.Println(longestStart)
}

func collatzSequenceLength(n int64) (length int64) {
	var result int64 = n
	length = 1
	for result != 1 {
		result = nextCollatz(result)
		length++
	}
	return length
}

func nextCollatz(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
