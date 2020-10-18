package lengthpattern

var builtQwertyKGM = false
var kgm = &keyGraphMap{
	Head:       nil,
	KeyNodeMap: nil,
}

// GetLength gets the length of a password.
// Substrings that can be connected in non-diagonal straight lines count as zero.
// Single char repeats (aaaaaaaa) count as zero as well.
func GetLength(password string) int {
	if len(password) == 0  {
		return 0
	}
	if len(password) == 1 {
		return 1
	}

	if !builtQwertyKGM {
		kgm = genQwertyKGM()
		builtQwertyKGM = true
	}

	length := 0

	prevNode := kgm.Get(rune(password[0]))
	for i := 1; i < len(password); i++ {
		currNode := kgm.Get(rune(password[i]))

		notNeighbor := false
		if !currNode.neighborOf(prevNode) {
			notNeighbor = true
			length++
		}

		currNodeIsPrevNode := false
		if currNode == prevNode {
			currNodeIsPrevNode = true
			length--
		}

		if notNeighbor && !currNodeIsPrevNode && i == 1 {
			length++ // first char in password is not part of a pattern itself, so lets count it
		}
		prevNode = currNode
	}
	return length
}

func isSelf(charNode *keyNode, currNode *keyNode) bool {
	return charNode != currNode
}

func genQwertyKGM() *keyGraphMap {
	// Keyboard rows are separated by newlines.
	// Spaces denote a nil neighbor and are used to align keys.
	numRow       := "`1234567890-= " + "\n"
	numRowShift  := `~!@#$%^&*()_+ ` + "\n"
	topRow       := ` qwertyuiop[]\` + "\n"
	topRowShift  := ` QWERTYUIOP{}|` + "\n"
	homeRow      := ` asdfghjkl;'  ` + "\n"
	homeRowShift := ` ASDFGHJKL:"  ` + "\n"
	botRow       := ` zxcvbnm,./   ` + "\n"
	botRowShift  := ` ZXCVBNM<>?   `

	symbols := numRow+topRow+homeRow+botRow
	shiftSymbols := numRowShift+topRowShift+homeRowShift+botRowShift
	return newKeyGraphMap(symbols, shiftSymbols)
}


