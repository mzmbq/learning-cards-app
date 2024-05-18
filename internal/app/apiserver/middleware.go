package apiserver

import (
	"context"
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type middlewareMux struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

func newMiddlewareMux() *middlewareMux {
	return &middlewareMux{
		mux:         http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

func (m *middlewareMux) AddMiddleware(mw Middleware) {
	m.middlewares = append(m.middlewares, mw)
}

func (m *middlewareMux) Handle(pattern string, handler http.Handler) {
	for i := range m.middlewares {
		// aply last added middleware first
		handler = m.middlewares[len(m.middlewares)-1-i](handler)
	}
	m.mux.Handle(pattern, handler)
}

func (m *middlewareMux) HandleFunc(pattern string, handler http.HandlerFunc) {
	m.Handle(pattern, handler)
}

func (m *middlewareMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

// Middlewares

func (s *server) withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")

		origin := r.Header.Get("Origin")
		if origin == "" {
			h.ServeHTTP(w, r)
		}

		for _, o := range s.corsOrigins {
			if origin == o {
				w.Header().Set("Access-Control-Allow-Origin", o)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				break
			}
		}

		h.ServeHTTP(w, r)
	})
}

func withLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		h.ServeHTTP(w, r)
	})
}

func (s *server) withAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		id, ok := session.Values["userID"]
		if !ok {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ctxKeyUserID, id)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
