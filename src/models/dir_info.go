package models

type DirInfo struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	IsFile bool   `json:"is_file"`
}
