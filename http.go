package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var (
	fileBox *rice.Box
)

func httpInit() http.Handler {
	m := mux.NewRouter()
	m.HandleFunc("/stream/{streamId}", httpLoggingMiddleware(httpHandleStream)).Methods("GET").Name("stream")
	m.HandleFunc("/{rest:.*}", httpLoggingMiddleware(httpHandleStatic)).Methods("GET").Name("static")
	fileBox = rice.MustFindBox("public/dist")
	return m
}

func httpHandleStream(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Stream"))
}

func httpHandleStatic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["rest"]
	b, err := fileBox.Bytes(path)
	if err != nil {
		w.Write(fileBox.MustBytes("index.html"))
	}
	w.Write(b)
}

func httpLoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		route := mux.CurrentRoute(r)
		fields := log.Fields{
			"host":    r.RemoteAddr,
			"method":  r.Method,
			"handler": route.GetName(),
			"path":    r.URL.String(),
			//"status":  r.Response.StatusCode,
		}
		logger.WithFields(fields).Debug("Serving HTTP")
	})
}
