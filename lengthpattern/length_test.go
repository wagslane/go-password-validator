package lengthpattern

import "testing"

func TestGetLength(t *testing.T) {
	actual := getLength("aaaa")
	expected := 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("aa")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("aaaaAAAA") // a and A are on the same key; repeated keys
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("1111!!!!")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}


	actual = getLength("bbbbbbbb")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("12121234")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("qwertyuiop[]vn")
	expected = 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("qti")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("ckr")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("c")
	expected = 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("ck")
	expected = 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("   ")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength(" ti")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength(" ti ")
	expected = 4
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength(" bb ")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength(" bbbbbbbbbb ")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength(" bbbbaaabbbb bb")
	expected = 6
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("lbbbbaaabbbblbbbbb")
	expected = 6
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("lbabblbb")
	expected = 6
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("llbabblbb")
	expected = 5
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("cv") // it's a pattern from the very first char, so its 0
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("rtyiop") // it's a pattern from the very first char, so its 0
	expected = 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getLength("password")
	expected = 5
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}