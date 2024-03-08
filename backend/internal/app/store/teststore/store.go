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

// TODO: implement

func (s *Store) Deck() store.DeckRepository {
	return nil
}

func (s *Store) Card() store.CardRepository {
	return nil
}
