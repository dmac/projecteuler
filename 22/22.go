package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var namesFile = "names.txt"

func main() {
	namesBytes, _ := ioutil.ReadFile(namesFile)
	names := strings.Split(strings.Replace(string(namesBytes), "\"", "", -1), ",")
	sort.Strings(names)
	sum := 0
	for i, name := range names {
		sum += nameScore(name, i+1)
	}
	fmt.Println(sum)
}

func nameScore(name string, i int) int {
	sum := 0
	for _, val := range name {
		sum += int(val) - 64
	}
	return sum * i
}
