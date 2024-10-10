package middleware

import (
	"log"
	"net/http"

	"github.com/vshalt/goserve/internal/config"
)

func Logger(logger *log.Logger, opt config.Flags) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(writer http.ResponseWriter, req *http.Request) {
			resp := ResponseWriter{writer: writer}
			defer func() {
				if opt.Verbose {
					logger.Println(req.RemoteAddr, req.Method, req.URL.Path, resp.status, req.UserAgent())
				}
			}()
			next.ServeHTTP(&resp, req)
		}
		return http.HandlerFunc(fn)
	}
}
