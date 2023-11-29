package encrypt

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password, secret string) {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := h.Sum(nil)
	fmt.Println("hashed", []byte(hashedPassword))
}
