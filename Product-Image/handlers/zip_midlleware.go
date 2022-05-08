package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

func NewGzipHandler() *GzipHandler {
	return &GzipHandler{}
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
			wrw := NewWrapperedResponseWriter(writer)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, request)
			defer wrw.Flush()
			return
		}

		next.ServeHTTP(writer, request)
	})
}

type WrapperedResponseWriter struct {
	rw         http.ResponseWriter
	gzipwriter *gzip.Writer
}

func NewWrapperedResponseWriter(rw http.ResponseWriter) *WrapperedResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WrapperedResponseWriter{rw: rw, gzipwriter: gw}
}

func (wr *WrapperedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrapperedResponseWriter) Write(p []byte) (n int, err error) {
	return wr.gzipwriter.Write(p)
}

func (wr *WrapperedResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

func (wr *WrapperedResponseWriter) Flush() {
	wr.gzipwriter.Flush()
	wr.gzipwriter.Close()
}
