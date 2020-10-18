package passwordvalidator

import "github.com/lane-c-wagner/go-password-validator/lengthpattern"

// UsePatternPenalty, when true, counts substrings that can be connected in a continuous (non-diagonal) line count as zero.
var UsePatternPenalty = false

func getLength(password string) int {
	if UsePatternPenalty {
		return lengthpattern.GetLength(password)
	}
	const maxNumSameChar = 2
	chars := map[rune]int{}
	for _, c := range password {
		if _, ok := chars[c]; !ok {
			chars[c] = 0
		}
		if chars[c] >= maxNumSameChar {
			continue
		}
		chars[c]++
	}
	length := 0
	for _, count := range chars {
		length += count
	}
	return length
}
