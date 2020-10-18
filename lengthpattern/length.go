package lengthpattern

// Gets the length.
// Key patterns that can be connected in non-diagonal straight lines count as zero.
//func getLengthPattern(password string, graphMap *keyGraphMap) int {
//	if len(password) == 0 {
//		return 0
//	}
//	length := 0
//
//	for _, char := range password {
//		shiftedSymbol := graphMap.KeyNodeMap[].Key.ShiftSymbol
//	}
//	return 0
//}

type key struct {
	Symbol rune

	// What is typed when shift is held down.
	ShiftSymbol rune
}
