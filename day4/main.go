package main

import (
	"fmt"
	"strconv"
)

/*
--- Day 4: Secure Container ---
You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.

However, they do remember a few key facts about the password:

It is a six-digit number.
The value is within the range given in your puzzle input.
Two adjacent digits are the same (like 22 in 122345).
Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
Other than the range rule, the following are true:

111111 meets these criteria (double 11, never decreases).
223450 does not meet these criteria (decreasing pair of digits 50).
123789 does not meet these criteria (no double).
How many different passwords within the range given in your puzzle input meet these criteria?
*/

func isValidNum(number string) bool {
	var prevDigit uint8
	var hasConsecutiveDigit bool
	prevDigit = number[0]
	for i:=1;i<len(number);i++ {
		digit := number[i]
		if digit == prevDigit {
			hasConsecutiveDigit = true
		} else if digit < prevDigit {
			return false
		}
		prevDigit = digit
	}
	return hasConsecutiveDigit
}


func main() {
	lowerBound := 273025
	upperBound := 767253
	count := 0
	for i:=lowerBound; i<=upperBound; i++ {
		if isValidNum(strconv.Itoa(i)) {
			count++
		}
	}
	fmt.Printf("Part 1: Num Passwords: %d\n", count)
}
