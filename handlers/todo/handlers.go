package todo

import (
	"fmt"
	"html/template"
	"net/http"
)

var todos []Todo = ReadTodos("hello@mail.com")

func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("title") == "" {
		return
	}
	var returningTodo Todo

	err := InsertTodo(&returningTodo, r.FormValue("title"), "hello@mail.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	itemTmpl, err := template.ParseFiles("pages/todo/item.html")
	if err != nil {
		fmt.Println(err)
	}
	itemTmpl.Execute(w, returningTodo)
	todos = append(todos, returningTodo)
}

func Checked(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var currTodo Todo
	tmpl, err := template.ParseFiles("pages/todo/item.html")
	if err != nil {
		fmt.Println(err)
	}
	for _, todo := range todos {
		if todo.Id == id {
			currTodo = todo
		}
	}
	currTodo.Done = true
	tmpl.Execute(w, currTodo)
}
