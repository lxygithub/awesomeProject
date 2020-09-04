package models

import "os"

type DirInfo struct {
	file     os.FileInfo
	DirInfos []DirInfo
}
