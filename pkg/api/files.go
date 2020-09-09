package api

import (
	"encoding/json"
	"fmt"
	"github.com/models"
	"html/template"
	"io/ioutil"
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
	name := r.PostForm.Get("dir")
	var dirs []os.FileInfo
	if name != "" && strings.HasPrefix(name, topPath) {
		dirs = GetDirTree(name)
		currentPath = name
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
	if strings.Contains(r.Header.Get("content-type"), "json") {
		b, _ := json.Marshal(models.BaseResponse{Code: 0, ErrMsg: "", Data: fileList})
		fmt.Fprintf(w, string(b))
	} else {
		t, _ := template.ParseFiles("web/templates/files.html")
		type TemplateData struct {
			Files    []models.DirInfo
			LastPath string
		}
		t.Execute(w, TemplateData{Files: fileList, LastPath: filepath.Dir(currentPath)})
	}
}

func GetDirTree(currentPath string) []os.FileInfo {
	dirs, _ := ioutil.ReadDir(currentPath)
	return dirs
}
