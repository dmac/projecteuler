package main

import (
	"fmt"
)

func main() {
	fmt.Println(squareOfSums(100) - sumOfSquares(100))
}

func sumOfSquares(n int64) (sum int64) {
	var i int64
	for i = 1; i <= n; i++ {
		sum += i * i
	}
	return
}

func squareOfSums(n int64) int64 {
	var i int64
	var sum int64 = 0
	for i = 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
