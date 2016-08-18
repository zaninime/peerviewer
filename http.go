package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	rice "github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var (
	fileBox *rice.Box
)

func httpInit() http.Handler {
	m := mux.NewRouter()
	api := m.PathPrefix("/api").Subrouter()
	m.HandleFunc("/stream/{streamID}", httpLoggingMiddleware(httpHandleStream)).Methods("GET").Name("stream")
	m.HandleFunc("/{rest:.*}", httpLoggingMiddleware(httpHandleStatic)).Methods("GET").Name("static")
	api.HandleFunc("/streams", httpJSONHeaderMiddleware(httpHandleAPIStreams(m)))
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

func httpHandleAPIStreams(router *mux.Router) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		type stream struct {
			Description string
			Kind        string
			URL         string
		}
		streams := make([]stream, len(config.Streams))
		for i, v := range config.Streams {
			s := &streams[i]
			s.Description = v.Description
			s.Kind = strconv.Itoa(int(v.Kind.Value))
			var streamURL *url.URL
			streamURL, err = router.Get("stream").URL("streamID", strconv.Itoa(i))
			if err != nil {
				panic(err)
			}
			s.URL = streamURL.String()
		}
		j := json.NewEncoder(w)
		j.Encode(streams)
	}
}

func httpJSONHeaderMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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
