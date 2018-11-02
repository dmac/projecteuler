package main

import (
	"fmt"
)

func main() {
	var n int64 = 1
	for {
		if divisibleByAll(n, 20) {
			fmt.Println(n)
			break
		}
		n++
	}
}

func divisibleByAll(n int64, upto int64) bool {
	if n < upto {
		return false
	}
	var i int64
	for i = 1; i <= upto; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
