package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

		s.WriteJSON(w, http.StatusOK, u.ID)
	}
}

func (s *server) handleUserFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.PathValue("email")

		u, err := s.store.User().FindByEmail(email)
		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		s.WriteJSON(w, http.StatusOK, u)

	}
}

func (s *server) handleUserAuth() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.CheckPassword(req.Password) {
			http.Error(w, "email or password incorrect", http.StatusUnauthorized)
			return
		}

		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		session.Values["userID"] = u.ID
		session.Values["email"] = u.Email

		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		s.WriteJSON(w, http.StatusOK, u.ID)
	}
}

func (s *server) handlerUserSignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		delete(session.Values, "userID")
		delete(session.Values, "email")
		session.Save(r, w)
	}

}

func (s *server) handleUserWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(ctxKeyUserID)
		if id == nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		idInt, ok := id.(int)
		if !ok {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		u, err := s.store.User().Find(idInt)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		s.WriteJSON(w, http.StatusOK, u)
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
