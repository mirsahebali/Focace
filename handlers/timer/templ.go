package timer

import (
	"fmt"
	"html/template"
	"net/http"

	"mirsahebali/focace/helpers/parsers"
)

var JSLoaded = false

func RenderTemplate(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("pages/timer/index.html")
	if err != nil {
		fmt.Println(err)
	}
	parsedJS, err := parsers.JSParser("pages/timer/timer.js")
	if err != nil {
		fmt.Println(err)
	}
	parseFlushJS, err := parsers.JSParser("pages/timer/flush.js")
	if err != nil {
		fmt.Println(err)
	}
	if !JSLoaded {
		JSLoaded = true
		tmpl.Execute(w, parsedJS)
	} else {
		tmpl.Execute(w, parseFlushJS)
	}
}
