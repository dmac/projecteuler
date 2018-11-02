package main

import (
	"fmt"
)

const SIZE = 100000

var cache [SIZE + 1][SIZE + 1]int64

func main() {
	fmt.Println(numPaths(SIZE, SIZE))
}

func numPaths(width int64, height int64) int64 {
	if cache[width][height] != 0 {
		return cache[width][height]
	}
	if width == 0 && height == 0 {
		return 1
	}
	var paths int64 = 0
	if width > 0 {
		paths += numPaths(width-1, height)
	}
	if height > 0 {
		paths += numPaths(width, height-1)
	}

	cache[width][height] = paths
	return paths
}
