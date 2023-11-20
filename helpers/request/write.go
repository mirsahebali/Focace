package request

import "net/http"

func WriteResponse(w http.ResponseWriter, statusCode int, writeString string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(writeString))
}
