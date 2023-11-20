package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"mirsahebali/focace/helpers/parsers"
	"mirsahebali/focace/helpers/request"
	"mirsahebali/focace/token"
)

var SECRET = parsers.GetEnvVars("FOCACE_AUTH_TOKEN")

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// NOTE change this from r.PostFormValue to r.Header for broader use

	if r.PostFormValue("Email") == "" {
		request.WriteResponse(w, http.StatusNotAcceptable, "email not found")
		return
	}

	if r.PostFormValue("Password") == "" {

		request.WriteResponse(w, http.StatusNotAcceptable, "password not found")
		return
	}
	if r.PostFormValue("Username") == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("username not found"))
		return
	}
	// NOTE change this from r.PostFormValue to r.Header for broader use
	check := AddUserObject(User{
		Username:     r.PostFormValue("Username"),
		Email:        r.PostFormValue("Email"),
		Passwordhash: r.PostFormValue("Password"),
		Fullname:     r.PostFormValue("Fullname"),
	})
	if !check {
		request.WriteResponse(w, http.StatusConflict, "user already exists")
		return
	}
	request.WriteResponse(w, http.StatusOK, "user created")
}

func getSignedToken() (string, error) {
	header := "HS256"
	// TODO: Keep this as env variable

	claims := map[string]string{
		"iss": "mirsahebali.dev",
		"aud": "focace.io",
		"exp": fmt.Sprint(time.Now().Unix()),
	}
	generated_token, err := token.GenerateToken(header, claims, SECRET)
	if err != nil {
		return generated_token, errors.New("error generating token")
	}
	return generated_token, nil
}

func validateUser(email, passwordHash string) (bool, error) {
	usr, exists := GetUserObject(email)
	if !exists {
		return false, errors.New("user does not exists")
	}
	valid := usr.ValidatePasswordHash(passwordHash)
	if !valid {
		return false, errors.New("incorrect password")
	}
	return true, nil
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		request.WriteResponse(w, http.StatusNoContent, "email is missing")
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		request.WriteResponse(w, http.StatusNoContent, "password not found")
		return
	}
	exists, err := validateUser(r.Header["Email"][0], r.Header["Passwordhash"][0])
	if err != nil {
		request.WriteResponse(w, http.StatusNotFound, "user not found")
		return
	}
	if !exists {
		request.WriteResponse(w, http.StatusUnauthorized, "Unauthorized request")
		return
	}
	token, err := getSignedToken()
	if err != nil {
		request.WriteResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	request.WriteResponse(w, http.StatusOK, token)
}
