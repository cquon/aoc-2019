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

--- Part Two ---
An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

Given this additional criterion, but still ignoring the range rule, the following are now true:

112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
How many different passwords within the range given in your puzzle input meet all of the criteria?

Your puzzle input is still 273025-767253.

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

func isValidNumPt2(number string) bool {
	var prevDigit uint8
	var hasConsecutiveDigit bool
	var continuousCount = 1
	prevDigit = number[0]
	for i:=1;i<len(number);i++ {
		digit := number[i]
		if digit == prevDigit {
			continuousCount += 1
		} else if digit < prevDigit {
			return false
		} else {
			if continuousCount == 2 {
				hasConsecutiveDigit = true
			}
			continuousCount = 1
		}
		prevDigit = digit
	}
	return hasConsecutiveDigit || continuousCount == 2
}


func main() {
	lowerBound := 273025
	upperBound := 767253
	count := 0
	countPt2 := 0
	for i:=lowerBound; i<=upperBound; i++ {
		if isValidNum(strconv.Itoa(i)) {
			count++
		}
		if isValidNumPt2(strconv.Itoa(i)) {
			countPt2++
		}
	}
	fmt.Printf("Part 1: Num Passwords: %d\n", count)
	fmt.Printf("Part 2: Num Passwords: %d\n", countPt2)
}
