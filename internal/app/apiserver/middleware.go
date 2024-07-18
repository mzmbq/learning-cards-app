package apiserver

import (
	"context"
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func (s *server) withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")

		origin := r.Header.Get("Origin")
		if origin == "" {
			h.ServeHTTP(w, r)
			return
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

func (s *server) withLogging(h http.Handler) http.Handler {
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

func (s *server) withRateLimit(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.rateLimiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}

		h.ServeHTTP(w, r)
	})
}
