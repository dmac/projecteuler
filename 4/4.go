package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1, n2, p, largest int64
	largest = 0
	for n1 = 100; n1 <= 999; n1++ {
		for n2 = 100; n2 <= 999; n2++ {
			p = n1 * n2
			if isPalindrome(p) && p > largest {
				largest = p
			}
		}
	}
	fmt.Println(largest)
}

func isPalindrome(n int64) bool {
	s := strconv.FormatInt(n, 10)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
