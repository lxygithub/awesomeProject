package api

import (
	"log"
	"net/http"
)

func Route() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/package_list", PackageList)
	http.HandleFunc("/index", Index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
