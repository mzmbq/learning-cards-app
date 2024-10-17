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
	s := &Store{}
	s.userRepository = NewUserRepo(s)
	s.deckRepository = NewDeckRepo(s)
	s.cardRepository = NewCardRepo(s)
	return s
}

func (s *Store) User() store.UserRepository {
	return s.userRepository
}

func (s *Store) Deck() store.DeckRepository {
	return s.deckRepository
}

func (s *Store) Card() store.CardRepository {
	return s.cardRepository
}
