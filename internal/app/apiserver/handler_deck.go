package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

func (s *server) handleDeckCreate() APIFunc {
	type request struct {
		DeckName string `json:"deckname"`
	}
	type response struct {
		DeckID int `json:"deck_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return err
		}

		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return InvalidJSON()
		}

		d := model.Deck{
			Name:   req.DeckName,
			UserID: u.ID,
		}

		if err := d.Validate(); err != nil {
			return ValidationErrors(err)
		}

		if err = s.store.Deck().Create(&d); err != nil {
			return NewAPIError(http.StatusBadRequest, "failed to create deck")
		}

		res := response{
			DeckID: d.ID,
		}
		return WriteJSON(w, http.StatusOK, res)
	}
}

func (s *server) handleDeckDelete() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return err
		}

		idStr := r.PathValue("id")
		if idStr == "" {
			return InvalidRequestData(map[string]string{"empty id": ""})
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return InvalidRequestData(map[string]string{"invalid deck id": idStr})
		}

		d, err := s.store.Deck().Find(id)
		if err != nil {
			return Unauthorized()
		}
		if d.UserID != u.ID {
			return Unauthorized()
		}

		err = s.store.Deck().Delete(id)
		if err != nil {
			return err
		}

		return WriteOK(w)
	}
}

func (s *server) handleDeckListCards() APIFunc {
	type response struct {
		Cards []model.Card `json:"cards"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return err
		}

		idStr := r.PathValue("id")
		if idStr == "" {
			return InvalidRequestData(map[string]string{"id": "empty id"})
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return InvalidRequestData(map[string]string{"id": fmt.Sprintf("invalid deck id: %v", idStr)})
		}

		d, err := s.store.Deck().Find(id)
		if err != nil {
			return Unauthorized()
		}
		if d.UserID != u.ID {
			return Unauthorized()
		}

		cards, err := s.store.Card().FindAllByDeckID(id)
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, response{Cards: cards})
	}
}

// Get the list of decks that belong to the current user
func (s *server) handleDecksList() APIFunc {
	type response struct {
		Decks []model.Deck `json:"decks"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return err
		}

		decks, err := s.store.Deck().FindAllByUserID(u.ID)
		if err != nil {
			return err
		}

		res := &response{
			Decks: decks,
		}
		return WriteJSON(w, http.StatusOK, res)
	}
}

func (s *server) handleDeckRename() APIFunc {
	type request struct {
		DeckName string `json:"deckname"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := s.userFromRequest(r)
		if err != nil {
			return err
		}
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return InvalidJSON()
		}

		// Extract the deck id from the request path
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return InvalidRequestData(map[string]string{"id": fmt.Sprintf("invalid deck id: %v", idStr)})
		}

		d, err := s.store.Deck().Find(id)
		if err != nil {
			return err
		}
		if d.UserID != u.ID {
			return Unauthorized()
		}

		// Validate the new deck name
		d.Name = req.DeckName
		if err := d.Validate(); err != nil {
			return ValidationErrors(err)
		}

		// Rename the deck
		err = s.store.Deck().Update(d)
		if err != nil {
			return err
		}
		return WriteOK(w)
	}
}
