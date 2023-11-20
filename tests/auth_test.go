package tests

import (
	"fmt"
	"mirsahebali/focace/auth"
	"testing"
)

func TestGetUserObject(t *testing.T) {
	user, check := auth.GetUserObject("foo@bar.com")
	if !check {
		t.Error("user does not exist")
	}
	fmt.Println("User found:")

	fmt.Println(user)
}
