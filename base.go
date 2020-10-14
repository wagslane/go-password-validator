package passwordvalidator

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
		if containsRune(specialChars, c) {
			hasSpecial = true
			continue
		}
		if containsRune(lowerChars, c) {
			hasLower = true
			continue
		}
		if containsRune(upperChars, c) {
			hasUpper = true
			continue
		}
		if containsRune(digitsChars, c) {
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

func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
