package apiserver

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/flashcard"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

func (s *server) handleStudyGetCard() APIFunc {
	type response struct {
		Card model.Card `json:"card"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return err
		}

		deckIDStr := r.PathValue("deck_id")
		deckID, err := strconv.Atoi(deckIDStr)
		if err != nil {
			return InvalidRequestData(map[string]string{"deck_id": "invalid"})

		}

		if !s.store.Deck().BelongsToUser(deckID, u.ID) {
			return Unauthorized()
		}

		// temporary solution: get the card with the earliest due date
		cards, err := s.store.Card().FindAllByDeckID(deckID)
		if err != nil {
			return err
		}
		if len(cards) == 0 {
			return NewAPIError(http.StatusNoContent, "no cards")
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Flashcard.Due.Before(cards[j].Flashcard.Due)
		})

		if cards[0].Flashcard.Due.After(time.Now().Add(5 * time.Minute)) {
			return NewAPIError(http.StatusNoContent, "no cards for now")
		}

		return WriteJSON(w, http.StatusOK, response{Card: cards[0]})
	}
}

func (s *server) handleStudySubmit() APIFunc {
	type request struct {
		CardID int `json:"card_id"`
		Status int `json:"status"` // flashcard.Again ... flashcard.Easy
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		u, err := s.userFromRequest(r)
		if err != nil {
			return Unauthorized()
		}
		// validate request
		if !s.store.Card().BelongsToUser(req.CardID, u.ID) {
			return err
		}
		if req.Status < 0 || req.Status > 3 {
			return InvalidRequestData(map[string]string{"status": "must be in [1,2,3]"})
		}
		// update flashcard
		c, err := s.store.Card().Find(req.CardID)
		if err != nil {
			return err
		}
		if err := flashcard.SuperMemoAnki(&c.Flashcard, req.Status); err != nil {
			return err
		}
		if err := s.store.Card().Update(c); err != nil {
			return err
		}

		return WriteOK(w)
	}
}
