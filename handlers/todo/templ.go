package todo

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Columns struct {
	CreatedAt time.Time
	Title     string
	Id        string
	Color     string
	todos     []Todo
}
type Todo struct {
	CreatedAt time.Time
	Email     string
	Title     string
	Id        string
	Done      bool
}

var pwd, err = os.Getwd()

func RenderTemplate(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("pages/todo/index.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, todos)
}
