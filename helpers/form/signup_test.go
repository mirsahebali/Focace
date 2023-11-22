package form

import (
	"testing"
)

func TestFormValidation(t *testing.T) {
	checkEmail, checkPassword := FormValidator("Hellom@heloo.com", "helooworldfoo1")
	if !checkPassword {
		t.Error("password is invalid")
	}
	if !checkEmail {
		t.Error("email is invalid")
	}
}
