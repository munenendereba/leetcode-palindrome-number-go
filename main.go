package main

import (
	"strconv"
)

func main() {
	isPalindrome2(20)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	y := strconv.Itoa(x)
	l := len(y) - 1 //7

	for i := 0; i <= l/2; i++ {
		n := y[i]
		if n != y[l-i] {
			return false
		}
	}

	return true
}

func isPalindrome2(x int) bool {
	result := 0
	input_num := x
	var remainder int
	for x > 0 {
		remainder = x % 10
		result = (result * 10) + remainder
		x /= 10
	}

	if input_num == result {
		return true
	} else {
		return false
	}
}
