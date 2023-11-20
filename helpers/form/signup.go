package form

import "strings"

func FormValidator(email, password string) (bool, bool) {
	containsAT := strings.ContainsRune(email, '@')
	containsDotCom := strings.Contains(email, ".com")
	return containsAT && containsDotCom, true
}
