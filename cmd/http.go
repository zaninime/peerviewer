package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func initHTTP() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(rice.MustFindBox("../public/dist").HTTPBox()))
	r.HandleFunc("/api", httpHelloHandler)
	return r
}

func httpHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
