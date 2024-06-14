package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

const sessionName = "session"

type contextKey string

const ctxKeyUserID contextKey = "userID"

type server struct {
	mux           *middlewareMux
	store         store.Store
	sessionsStore sessions.Store

	corsOrigins []string
}

func newServer(store store.Store, sessionsStore sessions.Store, corsOrigins []string) *server {
	s := &server{
		mux:           newMiddlewareMux(),
		store:         store,
		sessionsStore: sessionsStore,
		corsOrigins:   corsOrigins,
	}

	s.routes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
	}
}

func (s *server) userFromRequest(r *http.Request) (*model.User, error) {
	id := r.Context().Value(ctxKeyUserID)
	if id == nil {
		return nil, store.ErrRecordNotFound
	}

	idInt, ok := id.(int)
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	u, err := s.store.User().Find(idInt)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Handlers

func (s *server) handleCardLearn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}

func (s *server) handleWordDefine() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}
