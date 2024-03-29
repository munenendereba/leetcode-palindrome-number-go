package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	number   int
	expected bool
}{
	{121, true},
	{-121, false},
	{10, false},
	{11, true},
	{101, true},
	{1001, true},
	{110, false},
	{1011, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf("running test with input(%d)", test.number)

		t.Run(testname, func(t *testing.T) {
			out := isPalindrome(test.number)

			if out != test.expected {
				t.Errorf("isPalindrome(%d) expected %t, but got %t", test.number, test.expected, out)
			}
		})
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf("running test with input(%d)", test.number)

		t.Run(testname, func(t *testing.T) {
			out := isPalindrome2(test.number)

			if out != test.expected {
				t.Errorf("isPalindrome(%d) expected %t, but got %t", test.number, test.expected, out)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(101)
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2(101)
	}
}
