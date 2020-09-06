package api

import (
	"html/template"
	"io/ioutil"
	"models"
	"net/http"
	"os"
	"strings"
)

var pathSep = string(os.PathSeparator)
var topPath = "D:/go_space/awesomeProject/src/keke_test"
var currentPath string

func PackageList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query := r.URL.Query()
	names, ok := query["dir"]
	var dirs []os.FileInfo
	if ok && len(names[0]) > 0 && strings.HasPrefix(names[0], topPath) {
		dirs = GetDirTree(names[0])
		currentPath = names[0]
	} else {
		dirs = GetDirTree(topPath)
		currentPath = topPath
	}
	fileList := make([]models.DirInfo, 0)
	if dirs != nil {
		for i := range dirs {
			f := dirs[i]
			fileList = append(fileList, models.DirInfo{
				Name:   f.Name(),
				Path:   currentPath + pathSep + f.Name(),
				IsFile: !f.IsDir()})
		}
	}
	//b, _ := json.Marshal(fileList)
	//fmt.Fprintf(w, string(b))

	t, _ := template.ParseFiles("src/web/templates/files.html")
	data := map[string][]models.DirInfo{
		"files": fileList,
	}
	t.Execute(w, data)
}

func GetDirTree(currentPath string) []os.FileInfo {
	dirs, _ := ioutil.ReadDir(currentPath)
	return dirs
}
