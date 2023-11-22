package form

import (
	"strings"
)

func FormValidator(email, password string) (bool, bool) {
	//  Checking if there is less than 1 split
	splits := strings.Split(email, "@")
	// Return false if not (this also checks if the email has @ in it)
	if len(splits) != 2 {
		return false, false
	}
	// checks if there isn't any numbers or invalid characters in the beginning of the email
	if splits[0][0] <= 64 || (splits[0][0] > 90 && splits[0][0] < 97) {
		return false, false
	}
	// Checks if there exists a domain name between @ and .com/.net/.org
	nextAt := splits[1][0]

	if string(nextAt) == "." {
		return false, false
	}
	// variable which stores for the value of boolean of contaiining domains
	verifyDomain := false
	for _, domains := range []string{".com", ".net", ".org"} {
		if strings.Contains(splits[1], domains) {
			verifyDomain = true
		}
	}
	// checks if length of the password is less than 8
	if len(password) < 8 {
		return verifyDomain, false
	}
	checkCapitalLetters := false
	checkSymbolsOrNumbers := false
	var i byte
	for i = 65; i < 91; i++ {
		r := []byte{i}
		if strings.ContainsAny(password, string(r)) {
			checkCapitalLetters = true
		}
	}

	for i = 33; i < 65; i++ {
		r := []byte{i}
		if strings.ContainsAny(password, string(r)) {
			checkSymbolsOrNumbers = true
		}
	}

	return verifyDomain, checkCapitalLetters && checkSymbolsOrNumbers
}
