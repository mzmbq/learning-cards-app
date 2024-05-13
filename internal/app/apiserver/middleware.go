package apiserver

import (
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

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
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
