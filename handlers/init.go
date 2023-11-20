package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func InitTemplates(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("pages/init/index.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}
