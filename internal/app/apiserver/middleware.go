package apiserver

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Middleware func(http.Handler) http.Handler

func (s *server) withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		// Used by preflight cross-origin requests
		w.Header().Add("Vary", "Access-Control-Request-Method")

		origin := r.Header.Get("Origin")
		if origin == "" {
			h.ServeHTTP(w, r)
			return
		}

		for _, trustedOrigin := range s.CORSOrigins {
			if origin == trustedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", trustedOrigin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")

				// Check if the request is a preflight request
				if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
					w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, DELETE")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

					w.WriteHeader(http.StatusOK)
					return
				}

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

func (s *server) withGlobalRateLimit(h http.Handler) http.Handler {
	if s.globalLimiter == nil {
		return h
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.globalLimiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// Get the rate limiter based on the client IP
func (s *server) rateLimiterFromRequest(r *http.Request) *rate.Limiter {
	if s.userLimiter == nil {

		return nil
	}
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("error: ", err)
		return nil
	}
	c, ok := s.clients[ip]
	if !ok {
		log.Println("Creating new client. IP:", ip, ", rate limit:", s.userLimiter.Limit())
		c = &client{
			limiter:  rate.NewLimiter(s.userLimiter.Limit(), s.userLimiter.Burst()),
			lastSeen: time.Now(),
		}
		s.clients[ip] = c
	}

	return c.limiter
}

func (s *server) withUserRateLimit(h http.Handler) http.Handler {
	if s.userLimiter == nil {
		// log.Println("User rate limiting is disabled")
		return h
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := s.rateLimiterFromRequest(r)
		if !limiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}

		h.ServeHTTP(w, r)
	})
}
