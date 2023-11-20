package auth

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/auth/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(w, nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.FormValue("Email")))
	email := r.FormValue("Email")
	fmt.Println(email)

}
