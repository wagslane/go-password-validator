package lengthpattern

import "strings"

type keyNodeMap struct {
	internalKeyNodeMap map[key]*keyNode
	keyMap             *keyMap
}

func newKeyNodeMap(matrix [][]*keyNode) *keyNodeMap{
	res := &keyNodeMap{
		internalKeyNodeMap: map[key]*keyNode{},
		keyMap:             newKeyMap(matrix),
	}

	for rowNum, row := range matrix {
		for colNum, node := range row {
			res.internalKeyNodeMap[node.Key] = node
			populateNeighbors(node, matrix, colNum, rowNum)
		}
	}
	return res
}

// Gets Node associated with r.
// Accepts shift symbol or normal symbol for r.
func (knm keyNodeMap) Node(r rune) *keyNode {
	key := knm.keyMap.Get(r)
	return knm.internalKeyNodeMap[key]
}

// Used in keyGraphMap
type keyNode struct {
	Key key

	neighborTop    *keyNode
	neighborBottom *keyNode
	neighborLeft   *keyNode
	neighborRight  *keyNode
}

func (kn *keyNode) neighborOf(n *keyNode) bool {
	if n.neighborRight == kn || n.neighborLeft == kn || n.neighborBottom == kn || n.neighborTop == kn {
		return true
	}
	return false
}

type key struct {
	Symbol rune

	// What is typed when shift is held down.
	ShiftSymbol rune
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
			neighborTop:    nil, // populated later using populateNeighbors
			neighborBottom: nil,
			neighborLeft:   nil,
			neighborRight:  nil,
		})
	}
	return res
}
