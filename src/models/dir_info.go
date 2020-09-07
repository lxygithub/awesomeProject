package models

type DirInfo struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	IsFile   bool   `json:"is_file"`
	ModTime  string `json:"mod_time"`
	FileSize int64  `json:"file_size"`
	MB       int    `json:"mb"`
}
