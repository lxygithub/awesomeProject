package api

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html")
	data := map[string]string{
		"name": "aaaaaaaaaaaa",
	}

	t.Execute(w, data)
}
