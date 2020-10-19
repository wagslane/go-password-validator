package passwordvalidator

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	err := Validate("mypass", 50)
	expectedError := "Insecure password. Try including special characters, using uppercase letters, using numbers, using an uncommon password or using a longer password"
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	err = Validate("MYPASS", 50)
	expectedError = "Insecure password. Try including special characters, using lowercase letters, using numbers or using a longer password"
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	err = Validate("mypassword", 4)
	if err != nil {
		t.Errorf("Err should be nil")
	}

	err = Validate("aGoo0dMi#oFChaR2", 80)
	if err != nil {
		t.Errorf("Err should be nil")
	}

	for password, _ := range passwordsMap {
		err = Validate(password, 900)
		if !strings.Contains(err.Error(), "uncommon password") {
			t.Error("Err shouldn't be nil and should contain 'uncommon password' ", err, password)
		}
	}

	err = Validate("vSjasnel12", 90)
	expectedError = "Insecure password. Try including special characters, using an uncommon password or using a longer password"
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}
}
