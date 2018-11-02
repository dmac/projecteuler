package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, c float64
	a = 1
	for {
		b = a + 1
		for {
			c = math.Sqrt(float64(a*a + b*b))
			if a+b+c == 1000 {
				fmt.Printf("%d", int64(a*b*c))
				os.Exit(0)
			} else if a+b+c > 1000 {
				break
			}
			b++
		}
		a++
	}
}
