package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	var n int64 = 2
	for {
		if isPrime(n) {
			count++
		}
		if count == 10001 {
			fmt.Println(n)
			break
		}
		n++
	}
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
