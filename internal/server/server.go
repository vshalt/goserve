package server

import (
	"log"
	"net"
	"net/http"

	"github.com/vshalt/goserve/internal/config"
)

type HttpServer interface {
	ListenAndServe() error
	ListenAndServeTLS(certfile string, keyfile string) error
}

func GetHttpServer(addr string, handler http.Handler) HttpServer {
	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
	}
}

func Server(log *log.Logger, opt config.Flags) error {
	fs := NewFileServer(Options{Directory: opt.Directory})

	addr := net.JoinHostPort(opt.Host, opt.Port)
	server := GetHttpServer(addr, fs)

	if opt.SSL {
		log.Printf("HTTPS server listening at %s serving %s", addr, opt.Directory)
		return server.ListenAndServeTLS(opt.CertFile, opt.KeyFile)
	}
	log.Printf("HTTP server listening at %s serving %s", addr, opt.Directory)
	return server.ListenAndServe()
}
