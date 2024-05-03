package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type server struct {
	router *http.ServeMux
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: http.NewServeMux(),
		store:  store,
	}

	s.routes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			return
		}

		u := model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(&u); err != nil {
			http.Error(w, "create user failed", http.StatusInternalServerError)
			return
		}

		if err := s.WriteJSON(w, http.StatusOK, u.ID); err != nil {
			log.Println(err)
		}
	}
}

func (s *server) handleUserFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

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

		if err := s.WriteJSON(w, http.StatusOK, u); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *server) handleWordDefine() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}

func (s *server) handleDeckCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}

func (s *server) handleDecksList() http.HandlerFunc {

	type response struct {
		Decks []model.Deck `json:"decks"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)

		res := response{
			Decks: []model.Deck{},
		}

		for i := range 20 {
			res.Decks = append(res.Decks, model.Deck{
				ID:     i,
				Name:   fmt.Sprintf("Deck %d", i),
				UserID: 1,
			})
		}

		s.WriteJSON(w, 200, res)
	}

}

func (s *server) handleDeckGet() http.HandlerFunc {

	type response struct {
		Cards []model.Card `json:"cards"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)

		s.WriteJSON(w, 200, response{Cards: []model.Card{
			{ID: 1, DeckID: 1, Front: "Front 1", Back: "Back 1"},
			{ID: 2, DeckID: 1, Front: "Front 2", Back: "Back 2"},
			{ID: 3, DeckID: 1, Front: "Front 3", Back: "Back 3"},
			{ID: 4, DeckID: 1, Front: "Front 4", Back: "Back 4"},
			{ID: 5, DeckID: 1, Front: "Front 5", Back: "Back 5"},
		}})
	}
}

func (s *server) handleCardCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}

func (s *server) handleCardLearn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}

func (s *server) handleCardUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}

// Middleware

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
