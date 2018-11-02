package main

import (
	"fmt"
	"math"
)

func main() {
	var sum, i int64
	sum = 0
	for i = 0; i < 2E6; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isPrime(n int64) bool {
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
