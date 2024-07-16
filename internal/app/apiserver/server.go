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
	mux           *http.ServeMux
	store         store.Store
	sessionsStore sessions.Store

	middlewares []Middleware
	// allowed origins for CORS
	corsOrigins []string
}

func newServer(store store.Store, sessionsStore sessions.Store, corsOrigins []string) *server {
	s := &server{
		mux:           http.NewServeMux(),
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

func (s *server) Use(ms ...Middleware) {
	s.middlewares = append(s.middlewares, ms...)
}

func (s *server) Handle(pattern string, handler http.Handler) {
	for i := range s.middlewares {
		// aply last added middleware first
		handler = s.middlewares[len(s.middlewares)-1-i](handler)
	}
	s.mux.Handle(pattern, handler)
}

func (s *server) HandleNoMiddleware(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *server) handleHealthcheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
}
