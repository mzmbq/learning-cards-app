package apiserver

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"golang.org/x/time/rate"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

const sessionName = "session"

type contextKey string

const ctxKeyUserID contextKey = "userID"

type server struct {
	mux           *chi.Mux
	store         store.Store
	sessionsStore sessions.Store

	// allowed origins for CORS
	CORSOrigins []string
	// rate limiting
	globalLimiter *rate.Limiter
	userLimiter   *rate.Limiter
	clients       map[string]*client
	clientsMutex  sync.Mutex
}

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func newServer(store store.Store, sessionsStore sessions.Store, CORSOrigins []string, globalLimiter *rate.Limiter, userLimiter *rate.Limiter) *server {
	s := &server{
		mux:           chi.NewRouter(),
		store:         store,
		sessionsStore: sessionsStore,
		CORSOrigins:   CORSOrigins,
		globalLimiter: globalLimiter,
		userLimiter:   userLimiter,
		clients:       make(map[string]*client),
	}

	s.routes()

	go s.cleanupOldClients()

	return s
}

// Deletes rate limiters for clients that haven't been seen for 5 minutes
func (s *server) cleanupOldClients() {

	for {
		time.Sleep(1 * time.Minute)

		s.clientsMutex.Lock()
		// Delete rate limiter if client hasn't been seen for 5 minutes
		for ip, c := range s.clients {
			if time.Since(c.lastSeen) > 5*time.Minute {
				delete(s.clients, ip)
			}
		}
		s.clientsMutex.Unlock()
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

func WriteOK(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusOK, map[string]any{"statusCode": http.StatusOK})
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
