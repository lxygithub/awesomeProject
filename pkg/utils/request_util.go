package utils

import (
	"encoding/json"
	"net/http"
)

func GetParams(b *interface{}, r *http.Request) {
	_ = r.ParseForm()
	if r.Method == "GET" {
		json.Marshal(r.URL.Query())
	} else if r.Method == "POST" {
		json.Marshal(r.PostForm)
	}

}
