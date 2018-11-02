package main

import (
	"fmt"
	"math"
)

func main() {
	amicables := make(map[int]bool, 0)
	for i := 1; i < 10000; i++ {
		r1 := d(i)
		r2 := d(r1)
		if r2 == i && r1 < 10000 && r1 != i {
			amicables[i] = true
			amicables[r1] = true
		}
	}
	sum := 0
	for n := range amicables {
		sum += n
	}
	fmt.Println(sum)
}

func d(n int) int {
	return sumSlice(factorsOf(n))
}

func factorsOf(n int) (factors []int) {
	factors = append(factors, 1)
	for i := 2; i < int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			factors = append(factors, n/i)
		}
	}
	return
}

func sumSlice(l []int) (sum int) {
	for _, e := range l {
		sum += e
	}
	return
}
