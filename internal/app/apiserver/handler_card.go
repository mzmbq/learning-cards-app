package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/flashcard"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

func (s *server) handleCardCreate() APIFunc {
	type request struct {
		Card model.Card `json:"card"`
	}
	type response struct {
		CardID int `json:"card_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return Unauthorized()
		}
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return InvalidJSON()
		}

		c := req.Card
		fc := flashcard.New()
		c.Flashcard = *fc

		if !s.store.Deck().BelongsToUser(c.DeckID, u.ID) {
			return Unauthorized()
		}

		if err := s.store.Card().Create(&c); err != nil {
			return err
		}

		res := response{
			CardID: c.ID,
		}
		return WriteJSON(w, http.StatusOK, res)
	}
}

func (s *server) handleCardUpdate() APIFunc {
	type request struct {
		Card model.Card `json:"card"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return Unauthorized()
		}
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return InvalidJSON()
		}

		cardIDStr := r.PathValue("id")
		cardID, err := strconv.Atoi(cardIDStr)
		if err != nil {
			return InvalidRequestData(map[string]string{"id": "invalid"})
		}

		c := req.Card
		c.ID = cardID
		if !s.store.Card().BelongsToUser(c.ID, u.ID) {
			return Unauthorized()
		}

		if !s.store.Deck().BelongsToUser(c.DeckID, u.ID) {
			return Unauthorized()
		}

		if err = s.store.Card().Update(&c); err != nil {
			return err
		}
		return WriteOK(w)
	}
}

func (s *server) handleCardDelete() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return err
		}

		cardIDStr := r.PathValue("id")
		cardID, err := strconv.Atoi(cardIDStr)
		if err != nil {
			return InvalidRequestData(map[string]string{"id": "invalid"})
		}

		if !s.store.Card().BelongsToUser(cardID, u.ID) {
			return Unauthorized()
		}

		if err := s.store.Card().Delete(cardID); err != nil {
			return err
		}
		return WriteOK(w)
	}
}
