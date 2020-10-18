package passwordvalidator

import (
	"strings"
	"testing"
)

func TestGetLength(t *testing.T) {
	actual := getLength("aaaa")
	expected := 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("12121234")
	expected = 6
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

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
	kgm := NewKeyGraphMap(symbols, shiftSymbols)

	// Tests that neighbors are aligned correctly
	// starts at q
	actual := recordGraphTraversal(kgm.Root.NeighborRight.NeighborBottom,"rrrrrrrrrrdlllllllllldrrruuu")
	expected := `qwertyuiop[';lkjhgfdsazxcvfr4`
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actualNode := kgm.Map[key{
		Symbol:      'q',
		ShiftSymbol: 'Q',
	}]
	expectedNode := kgm.Root.NeighborRight.NeighborBottom
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
			currNode = currNode.NeighborRight
		}
		if dir == 'u' {
			currNode = currNode.NeighborTop
		}
		if dir == 'd' {
			currNode = currNode.NeighborBottom
		}
		if dir == 'l' {
			currNode = currNode.NeighborLeft
		}
		res.WriteRune(currNode.Key.Symbol)
	}
	return res.String()
}
