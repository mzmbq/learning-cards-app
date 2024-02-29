package apiserver

import (
	"fmt"
	"net/http"
)

type server struct {
	mux *http.ServeMux
}

func newServer() *server {
	s := &server{
		mux: http.NewServeMux(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.mux.HandleFunc("GET /cards/", apiHandler)
	s.mux.HandleFunc("GET /", rootHandler)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "(ʘ ͜ʖ ʘ)")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello")
}
