package middleware

import "net/http"

type ResponseWriter struct {
	writer http.ResponseWriter
	status int
}

func (rw *ResponseWriter) Header() http.Header {
	return rw.writer.Header()
}

func (rw *ResponseWriter) WriteHeader(status int) {
	rw.status = status
	rw.writer.WriteHeader(status)
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	return rw.writer.Write(b)
}
