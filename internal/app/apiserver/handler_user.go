package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			log.Println(err)
			return
		}

		u := model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := u.Validate(); err != nil {
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			log.Println("user validation failed for user: ", u, " error: ", err)
			return
		}

		uFound, err := s.store.User().FindByEmail(req.Email)
		if err != nil && err != store.ErrRecordNotFound {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if uFound != nil {
			http.Error(w, "a user with this email already exists", http.StatusConflict)
			return
		}

		if err := s.store.User().Create(&u); err != nil {
			http.Error(w, "create user failed", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		s.WriteJSON(w, http.StatusOK, u.ID)
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
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			log.Println(err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.CheckPassword(req.Password) {
			http.Error(w, "email or password incorrect", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		session.Values["userID"] = u.ID
		session.Values["email"] = u.Email

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
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
