package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

func (s *server) handleStudyGetCard() http.HandlerFunc {
	type response struct {
		Card model.Card `json:"card"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, http.StatusOK, response{Card: model.Card{ID: 0, Front: "Mock", Back: "Data", DeckID: 0}})
	}
}

func (s *server) handleStudySubmit() http.HandlerFunc {
	type request struct {
		Status int `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		log.Println(req.Status)
		w.WriteHeader(http.StatusOK)
	}
}
