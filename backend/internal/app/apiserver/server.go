package apiserver

import (
	"fmt"
	"net/http"
)

type server struct {
}

func newServer() *server {
	return &server{}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
