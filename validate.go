package passwordvalidator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// Validate returns nil if the password has greater than or
// equal to the minimum entropy. If not, an error is returned
// that explains how the password can be strengthened. This error
// is safe to show the client
func Validate(password string, minEntropy float64) error {
	entropy := getEntropy(password)
	if entropy >= minEntropy {
		return nil
	}

	hasSpecial := false
	hasLower := false
	hasUpper := false
	hasDigits := false
	for _, c := range password {
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
	}

	allMessages := []string{}

	if !hasSpecial {
		allMessages = append(allMessages, "including special characters")
	}
	if !hasLower {
		allMessages = append(allMessages, "using lowercase letters")
	}
	if !hasUpper {
		allMessages = append(allMessages, "using uppercase letters")
	}
	if !hasDigits {
		allMessages = append(allMessages, "using numbers")
	}
	if _, exists := passwordsMap[password]; exists {
		allMessages = append(allMessages, "using an uncommon password")
	}

	if len(allMessages) > 0 {
		return fmt.Errorf(
			"Insecure password. Try %v or using a longer password",
			strings.Join(allMessages, ", "),
		)
	}

	return errors.New("Insecure password. Try using a longer password")
}

var passwordsMap = genPasswordsMap()

func genPasswordsMap() map[string]struct{} {
	b, err := ioutil.ReadFile("most-common-passwords.txt")
	if err != nil {
		log.Fatal(err)
	}

	passwords := strings.Split(string(b),"\n")
	sort.Sort(byLength(passwords))

	passwordsMap := make(map[string]struct{}, 10000)
	for _, password := range passwords {
		passwordsMap[password] = struct{}{}
	}

	return passwordsMap
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}