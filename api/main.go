package api

import (
	"mirsahebali/focace/server"
)

func Run() error {
	server := server.InitServer()

	return server.ListenAndServe()
}
