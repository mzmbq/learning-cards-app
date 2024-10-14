package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

func (s *server) handleDeckCreate() http.HandlerFunc {
	type request struct {
		DeckName string `json:"deckname"`
	}
	type response struct {
		DeckID int `json:"deck_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		d := model.Deck{
			Name:   req.DeckName,
			UserID: u.ID,
		}

		if err := d.Validate(); err != nil {
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			log.Println("deck validation failed for deck: ", d, " error: ", err)
			return
		}

		if err = s.store.Deck().Create(&d); err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		res := response{
			DeckID: d.ID,
		}
		WriteJSON(w, http.StatusOK, res)
	}
}

func (s *server) handleDeckDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		idStr := r.PathValue("id")
		if idStr == "" {
			http.Error(w, "no deck with id: "+idStr, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid deck id", http.StatusBadRequest)
			log.Println(err)
			return
		}

		d, err := s.store.Deck().Find(id)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			log.Println(err)
			return
		}
		if d.UserID != u.ID {
			http.Error(w, "", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		err = s.store.Deck().Delete(id)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) handleDeckListCards() http.HandlerFunc {
	type response struct {
		Cards []model.Card `json:"cards"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		idStr := r.PathValue("id")
		if idStr == "" {
			http.Error(w, "no deck with id: "+idStr, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid deck id", http.StatusBadRequest)
			log.Println(err)
			return
		}

		d, err := s.store.Deck().Find(id)
		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			log.Println(err)
			return
		}
		if d.UserID != u.ID {
			http.Error(w, "", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		cards, err := s.store.Card().FindAllByDeckID(id)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		WriteJSON(w, http.StatusOK, response{Cards: cards})
	}
}

func (s *server) handleDecksList() http.HandlerFunc {
	type response struct {
		Decks []model.Deck `json:"decks"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		decks, err := s.store.Deck().FindAllByUserID(u.ID)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		res := &response{
			Decks: decks,
		}
		WriteJSON(w, http.StatusOK, res)
	}

}

func (s *server) handleDeckRename() http.HandlerFunc {
	type request struct {
		DeckName string `json:"deckname"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.userFromRequest(r)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		// Extract the deck id from the request path
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid deck id", http.StatusBadRequest)
			log.Println(err)
			return
		}

		d, err := s.store.Deck().Find(id)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if d.UserID != u.ID {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		// Validate the new deck name
		d.Name = req.DeckName
		if err := d.Validate(); err != nil {
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			log.Println("deck validation failed for deck: ", d, " error: ", err)
			return
		}

		// Rename the deck
		err = s.store.Deck().Update(d)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
