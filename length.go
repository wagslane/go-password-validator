package passwordvalidator

func getLength(password string) int {
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
