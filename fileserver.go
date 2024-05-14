package govite

import (
	"io/fs"
	"net/http"
	"path"
)

func (v *ViteAdapter) FileServer() http.Handler {
	// clean path ./dist to dist since fs.Sub cannot work with relative path
	outputPathCleaned := path.Clean(v.cfg.ViteOutputPath)
	outputDirFS, err := fs.Sub(v.fs, outputPathCleaned)
	if err != nil {
		panic(err)
	}
	return http.FileServerFS(outputDirFS)
}
