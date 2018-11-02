package main

import (
	"fmt"
	"math"
)

func main() {
	factors := primeFactors(600851475143)
	var max int64 = 0
	for _, v := range factors {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}

func primeFactors(n int64) []int64 {
	factorMap := make(map[int64]bool, 0)
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			intermediateFactors := append(primeFactors(i), primeFactors(n/i)...)
			for _, v := range intermediateFactors {
				factorMap[v] = true
			}
		}
	}
	var keys []int64
	for k := range factorMap {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return []int64{n}
	}
	return keys
}
