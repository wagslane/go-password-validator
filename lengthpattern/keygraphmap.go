package lengthpattern

import (
	"strings"
)

type keyGraphMap struct {
	Head       *keyNode
	KeyNodeMap map[key]*keyNode
	KeyMap     *keyMap
}

// Rows are separated by newlines.
// Spaces denote a nil neighbor and are used to align keys.
func NewKeyGraphMap(symbols, shiftSymbols string) *keyGraphMap {
	mat := genKeyNodeMatrix(symbols, shiftSymbols)

	//for _,v := range mat {
	//	for _,vv := range v {
	//		fmt.Println(string(vv.Key.Symbol), string(vv.Key.ShiftSymbol),vv.Key.Symbol,vv.Key.ShiftSymbol,vv.Key.Symbol - vv.Key.ShiftSymbol)
	//	}
	//}

	return genKeyGraphMap(mat)
}

func genKeyGraphMap(matrix [][]*keyNode) *keyGraphMap {
	res := &keyGraphMap{
		Head:       matrix[0][0], // Head.Key.Symbol should equal ~ for qwerty
		KeyNodeMap: map[key]*keyNode{},
		KeyMap:     NewKeyMap(matrix),
	}

	for rowNum, row := range matrix {
		for colNum, node := range row {
			res.KeyNodeMap[node.Key] = node
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

// Used in keyGraphMap
type keyNode struct {
	Key key

	NeighborTop    *keyNode
	NeighborBottom *keyNode
	NeighborLeft   *keyNode
	NeighborRight  *keyNode
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