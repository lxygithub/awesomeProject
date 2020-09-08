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
	query := r.URL.Query()
	names, ok := query["file"]
	if ok && len(names[0]) > 0 && strings.HasPrefix(names[0], topPath) {
		file, _ := filepath.Abs(names[0])
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
