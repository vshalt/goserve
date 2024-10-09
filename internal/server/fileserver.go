package server

import (
	"net/http"
)

type Options struct {
	Directory string
	Prefix    string
}

type FileServer struct {
	opt     Options
	handler http.Handler
}

func NewFileServer(option Options) *FileServer {
	fs := &FileServer{
		opt: option,
	}
	fs.handler = http.StripPrefix(option.Prefix, http.FileServer(http.Dir(option.Directory)))
	return fs
}

func (fs *FileServer) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, handler := range middlewares {
		fs.handler = handler(fs.handler)
	}
}

func (fs *FileServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fs.handler.ServeHTTP(w, req)
}
