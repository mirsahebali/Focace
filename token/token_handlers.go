package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	hsh := hmac.New(sha256.New, []byte(secret))

	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	payloadStr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error json marhsal")
	}
	payload64 := base64.StdEncoding.EncodeToString(payloadStr)

	message := header64 + "." + payload64

	unsignedStr := header + string(payloadStr)

	hsh.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(hsh.Sum(nil))

	token := message + "." + signature
	return token, nil
}

func ValidateToken(token, secret string) (bool, error) {
	splitStr := strings.Split(token, ".")
	if len(splitStr) < 3 {
		return false, errors.New("token error")
	}

	header64, err := base64.StdEncoding.DecodeString(splitStr[0])
	if err != nil {
		return false, errors.New("error decoding header")
	}
	payload64, err := base64.StdEncoding.DecodeString(splitStr[1])
	if err != nil {
		return false, errors.New("error decoding payload")
	}
	unsignedMessage := string(header64) + string(payload64)

	hsh := hmac.New(sha256.New, []byte(secret))
	hsh.Write([]byte(unsignedMessage))
	signature := base64.StdEncoding.EncodeToString(hsh.Sum(nil))
	if signature != splitStr[2] {
		return false, errors.New("unauthorized")
	}
	return true, nil
}
