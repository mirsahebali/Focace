package stats

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("pages/stats/index.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}
