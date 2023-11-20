package server

import (
	"net/http"

	"mirsahebali/focace/handlers"
	"mirsahebali/focace/handlers/auth"
	"mirsahebali/focace/handlers/timer"
	"mirsahebali/focace/handlers/todo"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.InitTemplates)
	mux.HandleFunc("/todo", todo.RenderTemplate)
	mux.HandleFunc("/add-todo", todo.AddTodo)
	mux.HandleFunc("/checked", todo.Checked)
	mux.HandleFunc("/timer", timer.RenderTemplate)
	// INFO: Authorization routes
	mux.HandleFunc("/auth", auth.RenderPage)
	mux.HandleFunc("/signup", auth.SignUpHandler)
	mux.HandleFunc("/signin", auth.SignInHandler)

	return mux
}
