package api

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/web/templates/index.html")
	//fmt.Fprintf(w,"hhhhhhhhhhhhhhhhhhhh")
	data := map[string]string{
		"name": "aaaaaaaaaaaa",
	}

	t.Execute(w, data)

}
