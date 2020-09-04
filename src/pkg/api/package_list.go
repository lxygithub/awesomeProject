package api

import (
	"io/ioutil"
	"net/http"
	"os"
)

var pathSep = string(os.PathSeparator)
var dirPath = "E:/keke_release/test"

func PackageList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := os.Chdir(dirPath)
	//var dirInfos = []models.DirInfo{}
	if err == nil {
		dirs, _ := ioutil.ReadDir(dirPath)
		for _, fileInfo := range dirs {
			if fileInfo.IsDir() {
				//append(dirPaths, dirPath+pathSep+fileInfo.Name())
			}
		}
	}

}

func getDirTree() {

}
