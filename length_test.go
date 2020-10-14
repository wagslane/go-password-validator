package passwordvalidator

import (
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
