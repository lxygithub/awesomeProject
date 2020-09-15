package api

import (
	"log"
	"net/http"
)

func Route() {
	//http.HandleFunc("/", Hello)
	http.HandleFunc("/files", PackageList)
	http.HandleFunc("/files/download", DownLoadFile)
	//http.HandleFunc("/index", Index)
	//http.HandleFunc("/qq_send", Send)

	//fs := http.FileServer(http.Dir("E:/keke_release/"))
	//http.Handle("/E:/keke_release/", http.StripPrefix("/E:/keke_release/", fs))
	err := http.ListenAndServe(":9090", nil)
	//http.ListenAndServe(":9090", http.FileServer(http.Dir("E:/keke_test")))

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
