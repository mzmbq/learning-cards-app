package apiserver

import (
	"encoding/json"
	"net/http"

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

		if !s.store.Card().BelongsToUser(&c, u.ID) {
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
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}

func (s *server) handleCardDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	}
}
