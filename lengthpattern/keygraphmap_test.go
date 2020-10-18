package lengthpattern

import (
	"strings"
	"testing"
)

func TestQwertyKeyGraphMap(t *testing.T) {
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
	kgm := newKeyGraphMap(symbols, shiftSymbols)

	actual := recordGraphTraversal(kgm.Head.neighborRight.neighborBottom,"rrrrrrrrrrdlllllllllldrrruuu")
	expected := `qwertyuiop[';lkjhgfdsazxcvfr4`
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actualNode := kgm.Get('q')
	expectedNode := kgm.Head.neighborRight.neighborBottom
	if actualNode != expectedNode {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actualNode = kgm.Get('Q')
	expectedNode = kgm.Head.neighborRight.neighborBottom
	if actualNode != expectedNode {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

// Directions: l = left, r = right, u = up, d = down
func recordGraphTraversal(start *keyNode, directions string)  string {
	var res strings.Builder
	res.WriteRune(start.Key.Symbol)

	currNode := start
	for _, dir := range directions {
		if dir == 'r' {
			currNode = currNode.neighborRight
		}
		if dir == 'u' {
			currNode = currNode.neighborTop
		}
		if dir == 'd' {
			currNode = currNode.neighborBottom
		}
		if dir == 'l' {
			currNode = currNode.neighborLeft
		}
		res.WriteRune(currNode.Key.Symbol)
	}
	return res.String()
}