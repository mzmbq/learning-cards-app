package teststore

import (
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
	deckRepository *DeckRepository
	cardRepository *CardRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = NewUserRepo(s)
	}

	return s.userRepository
}

func (s *Store) Deck() store.DeckRepository {
	if s.deckRepository == nil {
		s.deckRepository = NewDeckRepo(s)
	}

	return s.deckRepository
}

func (s *Store) Card() store.CardRepository {
	if s.cardRepository == nil {
		s.cardRepository = NewCardRepo(s)
	}

	return s.cardRepository
}
