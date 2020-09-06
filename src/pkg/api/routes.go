package api

import (
	"log"
	"net/http"
)

func Route() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/files", PackageList)
	http.HandleFunc("/index", Index)
	err := http.ListenAndServe(":9090", nil)
	http.ListenAndServe(":9090", http.FileServer(http.Dir("E:/keke_test")))
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
