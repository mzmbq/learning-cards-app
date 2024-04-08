package sqlstore

import (
	"database/sql"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	deckRepository *DeckRepository
	cardRepository *CardRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Deck() store.DeckRepository {
	if s.deckRepository != nil {
		return s.deckRepository
	}

	s.deckRepository = &DeckRepository{
		store: s,
	}

	return s.deckRepository
}

func (s *Store) Card() store.CardRepository {
	if s.cardRepository != nil {
		return s.cardRepository
	}

	s.cardRepository = &CardRepository{
		store: s,
	}

	return s.cardRepository
}
