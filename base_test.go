package passwordvalidator

import (
	"testing"
)

func TestGetBase(t *testing.T) {
	actual := getBase("abcd")
	expected := len(lowerChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getBase("abcdA")
	expected = len(lowerChars) + len(upperChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getBase("A")
	expected = len(upperChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getBase("!")
	expected = len(specialChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getBase("123")
	expected = len(digitsChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getBase("123ü")
	expected = len(digitsChars) + 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}
