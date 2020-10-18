package passwordvalidator

import (
	"fmt"
	"strings"
)

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

// Gets the length.
// Key patterns that can be connected in non-diagonal straight lines count as zero.
//func getLengthPattern(password string, graphMap *keyGraphMap) int {
//	if len(password) == 0 {
//		return 0
//	}
//	length := 0
//
//	for _, char := range password {
//		shiftedSymbol := graphMap.Map[].Key.ShiftSymbol
//	}
//	return 0
//}

type keyMap struct {
	keyMap map[rune]key // normalSymbols index to map
	shiftedKeyMap map[rune]key // shiftedSymbols as index to map
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

func NewKeyMap() *keyMap {
	numRow       := "`1234567890-= "
	numRowShift  := `~!@#$%^&*()_+ `
	topRow       := ` qwertyuiop[]\`
	topRowShift  := ` QWERTYUIOP{}|`
	homeRow      := ` asdfghjkl;'  `
	homeRowShift := ` ASDFGHJKL:"  `
	botRow       := ` zxcvbnm,./   `
	botRowShift  := ` ZXCVBNM<>?   `

	km := &keyMap{
		keyMap:        make(map[rune]key, len(numRow) + len(topRow) + len(homeRow) + len(botRow)),
		shiftedKeyMap: make(map[rune]key, len(numRow) + len(topRow) + len(homeRow) + len(botRow)),
	}
	km.PushString(numRow, numRowShift)
	km.PushString(topRow, topRowShift)
	km.PushString(homeRow, homeRowShift)
	km.PushString(botRow, botRowShift)
	return km
}

// Keys are linked like so: https://i.imgur.com/RjBMAEW.png
// Rows are separated by newlines.
// Spaces denote a nil neighbor.
func NewKeyGraphMap(symbols, shiftSymbols string) *keyGraphMap {
	// spaces used below to align keys and denote nil neighbors
	//numRow       := "`1234567890-= "
	//numRowShift  := `~!@#$%^&*()_+ `
	//topRow       := ` qwertyuiop[]\`
	//topRowShift  := ` QWERTYUIOP{}|`
	//homeRow      := ` asdfghjkl;'  `
	//homeRowShift := ` ASDFGHJKL:"  `
	//botRow       := ` zxcvbnm,./   `
	//botRowShift  := ` ZXCVBNM<>?   `

	//mat := [][]*keyNode{}
	//mat = append(mat, genKeyNodeArray(numRow, numRowShift))
	//mat = append(mat, genKeyNodeArray(topRow, topRowShift))
	//mat = append(mat, genKeyNodeArray(homeRow, homeRowShift))
	//mat = append(mat, genKeyNodeArray(botRow, botRowShift))

	mat := genKeyNodeMatrix(symbols, shiftSymbols)

	for _,v := range mat {
		for _,vv := range v {
			fmt.Println(string(vv.Key.Symbol), string(vv.Key.ShiftSymbol),vv.Key.Symbol,vv.Key.ShiftSymbol,vv.Key.Symbol - vv.Key.ShiftSymbol)
		}
	}

	return genKeyGraphMap(mat)
}

func genKeyGraphMap(matrix [][]*keyNode) *keyGraphMap {
	res := &keyGraphMap{
		Root: matrix[0][0],
		Map: map[key]*keyNode{},
	}

	for rowNum, row := range matrix {
		for colNum, node := range row {
			res.Map[node.Key] = node
			populateNeighbors(node, matrix, colNum, rowNum)
		}
	}
	return res
}

func populateNeighbors(kn *keyNode, matrix [][]*keyNode, col, row int) {
	nilNode := &keyNode{
		Key:           key{rune(' '),rune(' ')},
	}
	neighborLeft, neighborRight, neighborBottom, neighborTop := nilNode, nilNode, nilNode, nilNode

	if insideBounds(matrix, col-1, row) {
		neighborLeft = matrix[row][col-1]
	}
	if insideBounds(matrix, col+1, row) {
		neighborRight = matrix[row][col+1]
	}
	if insideBounds(matrix, col, row-1) {
		neighborTop = matrix[row-1][col]
	}
	if insideBounds(matrix, col, row+1) {
		neighborBottom = matrix[row+1][col]
	}

	kn.NeighborBottom = neighborBottom
	kn.NeighborTop = neighborTop
	kn.NeighborLeft = neighborLeft
	kn.NeighborRight = neighborRight
}

func insideBounds(matrix [][]*keyNode, col, row int) bool {
	if row > len(matrix)-1 || row < 0 {
		return false
	}

	if col > len(matrix[row])-1 || col < 0 {
		return false
	}
	return true
}

//https://upload.wikimedia.org/wikipedia/commons/thumb/9/99/Numpad.svg/1200px-Numpad.svg.png
//  /*-
// 789+
// 456+
// 123
// 00.
func buildKeypadKeyGraphMap() {

}

// symbols len and shiftSymbols len must be equivalent
func genKeyNodeArray(symbols, shiftSymbols string) []*keyNode {
	res := make([]*keyNode, 0, len(symbols))
	for i := 0; i < len(symbols); i++ {
		res = append(res, &keyNode{
			Key:            key{rune(symbols[i]), rune(shiftSymbols[i])},
			NeighborTop:    nil, // populated later using populateNeighbors
			NeighborBottom: nil,
			NeighborLeft:   nil,
			NeighborRight:  nil,
		})
	}
	return res
}

// each row must be separated by newline
func genKeyNodeMatrix(symbols, shiftSymbols string) [][]*keyNode {
	symArr := strings.Split(symbols,"\n")
	shSymArr := strings.Split(shiftSymbols,"\n")

	mat := make([][]*keyNode, 0, len(symArr))

	for i := 0; i < len(symArr)-1; i++ {
		mat = append(mat, genKeyNodeArray(symArr[i], shSymArr[i]))
	}
	return mat
}

type keyGraphMap struct {
	Root *keyNode

	// In case we hit a dead-end in the graph, we don't need to traverse the graph to find it
	Map map[key]*keyNode
}

// Used in keyGraphMap
type keyNode struct {
	Key key

	NeighborTop    *keyNode
	NeighborBottom *keyNode
	NeighborLeft   *keyNode
	NeighborRight  *keyNode
}

// Used in keyNode
type key struct {
	Symbol rune

	// What's typed when you hold down shift.
	ShiftSymbol rune
}
