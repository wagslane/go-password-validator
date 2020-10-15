package passwordvalidator

import "strings"

const (
	specialChars = ` !"#$%&'()*+,-./:;<=>?@[\]^_{|}~`
	lowerChars   = `abcdefghijklmnopqrstuvwxyz`
	upperChars   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digitsChars  = `0123456789`
)

func getBase(password string) int {
	chars := map[rune]struct{}{}
	for _, c := range password {
		chars[c] = struct{}{}
	}

	hasSpecial := false
	hasLower := false
	hasUpper := false
	hasDigits := false
	base := 0

	for c := range chars {
		if strings.ContainsRune(specialChars, c) {
			hasSpecial = true
			continue
		}
		if strings.ContainsRune(lowerChars, c) {
			hasLower = true
			continue
		}
		if strings.ContainsRune(upperChars, c) {
			hasUpper = true
			continue
		}
		if strings.ContainsRune(digitsChars, c) {
			hasDigits = true
			continue
		}
		base++
	}

	if hasSpecial {
		base += len(specialChars)
	}
	if hasLower {
		base += len(lowerChars)
	}
	if hasUpper {
		base += len(upperChars)
	}
	if hasDigits {
		base += len(digitsChars)
	}
	return base
}
