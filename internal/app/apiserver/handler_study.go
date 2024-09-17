package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/flashcard"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

func (s *server) handleStudyGetCard() http.HandlerFunc {
	type response struct {
		Card model.Card `json:"card"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		deckIDStr := r.PathValue("deck_id")
		deckID, err := strconv.Atoi(deckIDStr)
		if err != nil {
			http.Error(w, "invalid deck id", http.StatusBadRequest)
			return
		}

		if !s.store.Deck().BelongsToUser(deckID, u.ID) {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		// temporary solution: get the card with the earliest due date
		cards, err := s.store.Card().FindAllByDeckID(deckID)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if len(cards) == 0 {
			http.Error(w, "", http.StatusNoContent)
			return
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Flashcard.Due.Before(cards[j].Flashcard.Due)
		})

		if cards[0].Flashcard.Due.After(time.Now().Add(5 * time.Minute)) {
			http.Error(w, "", http.StatusNoContent)
			log.Print("Card ", cards[0].Front, " is due ", cards[0].Flashcard.Due, " now is ", time.Now())
			return
		}

		s.WriteJSON(w, http.StatusOK, response{Card: cards[0]})
	}
}

func (s *server) handleStudySubmit() http.HandlerFunc {
	type request struct {
		CardID int `json:"card_id"`
		Status int `json:"status"` // flashcard.Again ... flashcard.Easy
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		// validate request
		if !s.store.Card().BelongsToUser(req.CardID, u.ID) {
			http.Error(w, "", http.StatusUnauthorized)
			log.Print("card does not belong to user")
			return
		}
		if req.Status < 0 || req.Status > 3 {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		// update flashcard
		c, err := s.store.Card().Find(req.CardID)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if err := flashcard.SuperMemoAnki(&c.Flashcard, req.Status); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if err := s.store.Card().Update(c); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
