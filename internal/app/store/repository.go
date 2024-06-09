package store

import "github.com/mzmbq/learning-cards-app/backend/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(id int) (*model.User, error)
	FindByEmail(emal string) (*model.User, error)
}

type DeckRepository interface {
	Create(*model.Deck) error
	Find(int) (*model.Deck, error)
	FindAllByUserID(int) ([]model.Deck, error)
	Delete(int) error
}

type CardRepository interface {
	Create(*model.Card) error
	Find(int) (*model.Card, error)
}
