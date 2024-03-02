package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type server struct {
	mux   *http.ServeMux
	store store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		mux:   http.NewServeMux(),
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.mux.HandleFunc("GET /", rootHandler)

	s.mux.HandleFunc("POST /api/user/create/", s.handleUserCreate())
	s.mux.HandleFunc("GET /api/user/{id}", s.handleUserFind())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func (s *server) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		var u model.User
		if err := dec.Decode(&u); err != nil {
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			return
		}

		if err := s.store.User().Create(&u); err != nil {
			http.Error(w, "create user failed", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, u.ID)
	}
}

func (s *server) handleUserFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := strings.TrimPrefix(r.URL.Path, "/api/user/")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		u, err := s.store.User().Find(id)
		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(u); err != nil {
			log.Fatal(err)
		}
	}
}
