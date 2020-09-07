package api

import (
	"html/template"
	"io/ioutil"
	"models"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var pathSep = string(os.PathSeparator)
var topPath, _ = filepath.Abs("E:/keke_release")
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
			path, _ := filepath.Abs(currentPath + pathSep + f.Name())
			fileList = append(fileList, models.DirInfo{
				Name:     f.Name(),
				Path:     path,
				IsFile:   !f.IsDir(),
				ModTime:  f.ModTime().Format("2006-01-02 15:04:05"),
				FileSize: f.Size(),
				MB:       int(f.Size() / 1024 / 1024),
			})
		}
	}
	//b, _ := json.Marshal(fileList)
	//fmt.Fprintf(w, string(b))

	t, _ := template.ParseFiles("src/web/templates/files.html")
	type TemplateData struct {
		Files    []models.DirInfo
		LastPath string
	}
	t.Execute(w, TemplateData{Files: fileList, LastPath: filepath.Dir(currentPath)})
}

func GetDirTree(currentPath string) []os.FileInfo {
	dirs, _ := ioutil.ReadDir(currentPath)
	return dirs
}
