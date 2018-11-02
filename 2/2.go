package main

import "fmt"

func main() {
	sum := 0
	f0 := 0
	f1 := 1
	f2 := f0 + f1
	for f2 < 4E6 {
		if f2%2 == 0 {
			sum += f2
		}
		f0 = f1
		f1 = f2
		f2 = f0 + f1
	}
	fmt.Println(sum)
}
