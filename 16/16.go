package main

import (
	"fmt"
	"strconv"
	"strings"
)

var placeNames = map[int]string{
	2: "hundred",
	3: "thousand",
}

var tenPlaceNames = map[int]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var teenNames = map[int]string{
	0: "ten",
	1: "eleven",
	2: "twelve",
	3: "thirteen",
	4: "fourteen",
	5: "fifteen",
	6: "sixteen",
	7: "seventeen",
	8: "eighteen",
	9: "nineteen",
}

var digitNames = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func main() {
	names := make([]string, 0)
	for i := 1; i <= 1000; i++ {
		names = append(names, numberName(i))
	}
	bigName := strings.Join(names, "")
	fmt.Println(len(bigName))
}

func numberName(n int) string {
	digits := splitDigits(n)
	s := ""
	for i := 0; i < len(digits); i++ {
		place := len(digits) - i - 1
		nextDigit := 0
		if place > 0 {
			nextDigit = digits[i+1]
		}
		if place == 1 && n > 100 && n%100 > 0 {
			s = fmt.Sprintf("%sand", s)
		}
		s = fmt.Sprintf("%s%s", s, digitName(digits[i], nextDigit, place))
		// break after the teens
		if place == 1 && digits[i] == 1 {
			break
		}
	}
	return s
}

func digitName(digit int, nextDigit int, place int) string {
	if place == 0 {
		return digitNames[digit]
	}

	if place == 1 {
		if digit == 1 {
			return teenNames[nextDigit]
		} else {
			return tenPlaceNames[digit]
		}
	}

	if place >= 2 && digit > 0 {
		return fmt.Sprintf("%s%s", digitNames[digit], placeNames[place])
	}

	return ""
}

func splitDigits(n int) []int {
	nString := fmt.Sprintf("%d", n)
	digitStrings := strings.Split(nString, "")
	digits := make([]int, 0)
	for i := 0; i < len(digitStrings); i++ {
		r, _ := strconv.Atoi(digitStrings[i])
		digits = append(digits, r)
	}
	return digits
}
