package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func DownLoadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query := r.URL.Query()
	names, ok := query["file"]
	if ok && len(names[0]) > 0 && strings.HasPrefix(names[0], topPath) {
		file := names[0]
		if strings.HasPrefix(file, "E:/keke_release") {
			bytes, err := ioutil.ReadFile(file)
			if err == nil {
				fileName := file[strings.LastIndex(file, string(os.PathSeparator))+1:]
				w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
				w.Write(bytes)
			}
		} else {
			fmt.Fprintf(w, "文件不存在")
		}

	}
}
