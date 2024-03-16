package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(7 / 2)
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
