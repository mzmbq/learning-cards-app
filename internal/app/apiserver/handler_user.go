package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

func (s *server) handleUserCreate() APIFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return InvalidJSON()
		}

		u := model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := u.Validate(); err != nil {
			return err
		}

		uFound, err := s.store.User().FindByEmail(req.Email)
		if err != nil && err != store.ErrRecordNotFound {
			return err
		}
		if uFound != nil {
			return NewAPIError(http.StatusConflict, "a user with this email already exists")
		}

		if err := s.store.User().Create(&u); err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, u.ID)
	}
}

func (s *server) handleUserAuth() APIFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return InvalidJSON()
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.CheckPassword(req.Password) {
			return NewAPIError(http.StatusUnauthorized, "email or password incorrect")
		}

		// Get a session. Creates a new session if the sessions doesn't exist
		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			return err
		}

		session.Values["userID"] = u.ID
		session.Values["email"] = u.Email

		err = session.Save(r, w)
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, map[string]any{"statucCode": http.StatusOK, "id": u.ID})
	}
}

func (s *server) handlerUserSignOut() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			return err
		}

		delete(session.Values, "userID")
		delete(session.Values, "email")
		if err = session.Save(r, w); err != nil {
			return err
		}
		return WriteOK(w)
	}

}

func (s *server) handleUserWhoami() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.Context().Value(ctxKeyUserID)
		if id == nil {
			panic("withAuth middleware required")
		}

		idInt, ok := id.(int)
		if !ok {
			return fmt.Errorf("failed to convert id \"%v\" to string", id)
		}

		u, err := s.store.User().Find(idInt)
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, u)
	}
}
