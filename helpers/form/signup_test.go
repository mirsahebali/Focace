package form

import (
	"testing"
)

func TestFormValidation(t *testing.T) {

	checkEmail, checkPassword := FormValidator("Hellom@.com", "alksjdflkasdjflkasdjf")
	if !checkPassword {
		t.Error("password is wrong")
	}
	if !checkEmail {
		t.Error("email is invalid")
	}
}
