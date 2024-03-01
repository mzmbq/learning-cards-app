package teststore

import (
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = newUserRepo(s)
	}

	return s.userRepository
}
