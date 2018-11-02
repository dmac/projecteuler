package main

import (
	"fmt"
)

func permutations(ns []int) [][]int {
	if len(ns) <= 1 {
		return [][]int{ns}
	}
	var ps [][]int
	for i, n := range ns {
		var ns_ = make([]int, len(ns))
		copy(ns_, ns)
		rest := append(ns_[:i], ns_[i+1:]...)
		for _, p := range permutations(rest) {
			ps = append(ps, append([]int{n}, p...))
		}
	}
	return ps
}

func main() {
	ps := permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(ps[1e6-1])
}
