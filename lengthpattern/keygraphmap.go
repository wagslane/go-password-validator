package lengthpattern

type keyGraphMap struct {
	Head       *keyNode
	KeyNodeMap *keyNodeMap
}

// Keyboard rows are separated by newlines.
// Spaces denote a nil neighbor and are used to align keys.
func NewKeyGraphMap(symbols, shiftSymbols string) *keyGraphMap {
	mat := genKeyNodeMatrix(symbols, shiftSymbols)
	return genKeyGraphMap(mat)
}

func (kgm keyGraphMap) Get(r rune) *keyNode{
	return kgm.KeyNodeMap.Node(r)
}

func genKeyGraphMap(matrix [][]*keyNode) *keyGraphMap {
	return &keyGraphMap{
		Head:       matrix[0][0], // Head.Key.Symbol should equal ~ for qwerty
		KeyNodeMap: NewKeyNodeMap(matrix),
	}
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

	kn.neighborBottom = neighborBottom
	kn.neighborTop = neighborTop
	kn.neighborLeft = neighborLeft
	kn.neighborRight = neighborRight
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
