package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func DownLoadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var fileName string
	if r.Method == "GET" {
		fileName = r.URL.Query().Get("file")
	} else if r.Method == "POST" {
		fileName = r.PostForm.Get("file")
	}
	if fileName != "" && strings.HasPrefix(fileName, topPath) {
		file, _ := filepath.Abs(fileName)
		preFix, _ := filepath.Abs("E:/keke_release")
		if strings.HasPrefix(file, preFix) {
			bytes, err := ioutil.ReadFile(file)
			if err == nil {
				fileName := filepath.Base(file)
				w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
				w.Header().Set("Content-Length", strconv.FormatInt(int64(len(bytes)), 10))
				w.Write(bytes)
			}
		} else {
			fmt.Fprintf(w, "文件不存在")
		}

	}
}
