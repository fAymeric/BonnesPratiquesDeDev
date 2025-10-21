package internal

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path
	filePath := "." + filename

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found", 404)
		return
	}

	var fileType string
	switch filepath.Ext(filename) {
	case ".js":
		fileType = "application/javascript"
	}

	w.Header().Set("Content-Type", fileType)
	w.Write(file)
}
