package store

import "github.com/mzmbq/learning-cards-app/backend/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

type DeckRepository interface {
	Create(*model.Deck) error
	Find(int) (*model.Deck, error)
}
