package token

import (
	"net/http"

	"mirsahebali/focace/helpers/parsers"
	"mirsahebali/focace/helpers/request"
)

var SECRET = parsers.GetEnvVars("FOCACE_AUTH_TOKEN")

func MiddlewareHandler(_ http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header["Token"]; !ok {
			request.WriteResponse(w, http.StatusNotAcceptable, "token missing")
			return
		}
		token := r.Header["Token"][0]
		valid, err := ValidateToken(token, SECRET)
		if err != nil {
			request.WriteResponse(w, http.StatusInternalServerError, "Validation failed")
			return
		}
		if !valid {
			request.WriteResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		request.WriteResponse(w, http.StatusAccepted, "Authorized token")
	})
}
