package lengthpattern

import "testing"

func TestGetLength(t *testing.T) {
	actual := GetLength("aaaa")
	expected := 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("12121234")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("1!1!1!1!!!!111122222@@@")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("qwerasdf")
	expected = 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("19dl")
	expected = 4
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("aa")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("aaaaAAAA") // a and A are on the same key; repeated keys
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("1111!!!!")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}


	actual = GetLength("bbbbbbbb")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("12121234")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("qwertyuiop[]vn")
	expected = 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("qti")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("ckr")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("c")
	expected = 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("ck")
	expected = 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("   ")
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength(" ti")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength(" ti ")
	expected = 4
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength(" bb ")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength(" bbbbbbbbbb ")
	expected = 3
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength(" bbbbaaabbbb bb")
	expected = 6
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("lbbbbaaabbbblbbbbb")
	expected = 6
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("lbabblbb")
	expected = 6
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("llbabblbb")
	expected = 5
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("cv") // it's a pattern from the very first char, so its 0
	expected = 0
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("rtyiop") // it's a pattern from the very first char, so its 0
	expected = 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = GetLength("password")
	expected = 5
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}