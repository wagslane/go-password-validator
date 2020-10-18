package lengthpattern

// Used to make keyNodeMap.Node(r rune) use either normal symbol or shift symbol for arg r.
type keyMap struct {
	keyMap map[rune]key // normalSymbols index to map
	shiftedKeyMap map[rune]key // shiftedSymbols as index to map
}

func newKeyMap(matrix [][]*keyNode) *keyMap {
	km := &keyMap{
		keyMap:        make(map[rune]key),
		shiftedKeyMap: make(map[rune]key),
	}
	for _, row := range matrix {
		for _, col := range row {
			km.Push(col.Key)
		}
	}
	return km
}

// Pushes a key to the structure.
func (km *keyMap) Push(k key) {
	km.keyMap[k.Symbol] = k
	km.shiftedKeyMap[k.ShiftSymbol] = k
}

// Pushes keys composed of symbols and shiftSymbols to the structure.
func (km *keyMap) PushString(symbols, shiftSymbols string) {
	for i := 0; i < len(symbols); i++ {
		km.Push(key{rune(symbols[i]), rune(shiftSymbols[i])})
	}
}

// Gets shift symbol if normal symbol passed, and vice versa.
func (km keyMap) AntiSymbol(r rune) rune {
	if anti, ok := km.keyMap[r]; ok {
		return anti.ShiftSymbol
	}
	if anti, ok := km.shiftedKeyMap[r]; ok {
		return anti.Symbol
	}
	return 0
}

// Gets key regardless if r is shift or non-shift
func (km keyMap) Get(r rune) key {
	if key, ok := km.keyMap[r]; ok {
		return key
	}
	if key, ok := km.shiftedKeyMap[r]; ok {
		return key
	}
	return key{}
}