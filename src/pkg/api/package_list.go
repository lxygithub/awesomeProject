package api

import (
	"net/http"
	"os"
)

var pathSep = string(os.PathSeparator)
var dirPath = "E:/keke_release/test"

func PackageList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//err := os.Chdir(dirPath)
	//var dirPaths []string
	//if err == nil {
	//	wd, _ := os.Getwd()
	//	dirs, _ := ioutil.ReadDir(wd)
	//
	//	for _, fileInfo := range dirs {
	//		if fileInfo.IsDir() {
	//			append(dirPaths, dirPath+pathSep+fileInfo.Name())
	//		}
	//	}
	//	fmt.Fprintf(w, wd)
	//}

}
