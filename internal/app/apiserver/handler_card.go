package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

func (s *server) handleCardCreate() http.HandlerFunc {
	type request struct {
		Card model.Card `json:"card"`
	}

	type response struct {
		CardID int `json:"card_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		c := req.Card

		if !s.store.Deck().BelongsToUser(c.DeckID, u.ID) {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		if err := s.store.Card().Create(&c); err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		res := response{
			CardID: c.ID,
		}

		s.WriteJSON(w, http.StatusOK, res)
	}
}

func (s *server) handleCardUpdate() http.HandlerFunc {
	type request struct {
		Card model.Card `json:"card"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		cardIDStr := r.PathValue("id")
		cardID, err := strconv.Atoi(cardIDStr)
		if err != nil {
			http.Error(w, "invalid card id", http.StatusBadRequest)
			return
		}

		c := req.Card
		c.ID = cardID
		if !s.store.Card().BelongsToUser(c.ID, u.ID) {
			http.Error(w, "", http.StatusUnauthorized)
			log.Printf("unauthorized: card %d, user %d\n", c.ID, u.ID)
			return
		}

		if !s.store.Deck().BelongsToUser(c.DeckID, u.ID) {
			http.Error(w, "", http.StatusUnauthorized)
			log.Printf("unauthorized: deck %d, user %d\n", c.DeckID, u.ID)
			return
		}

		if err = s.store.Card().Update(&c); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) handleCardDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		cardIDStr := r.PathValue("id")
		cardID, err := strconv.Atoi(cardIDStr)
		if err != nil {
			http.Error(w, "invalid card id", http.StatusBadRequest)
			return
		}

		if !s.store.Card().BelongsToUser(cardID, u.ID) {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		if err := s.store.Card().Delete(cardID); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
