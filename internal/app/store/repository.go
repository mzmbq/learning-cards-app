package store

import "github.com/mzmbq/learning-cards-app/backend/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(id int) (*model.User, error)
	FindByEmail(emal string) (*model.User, error)
}

type DeckRepository interface {
	Create(*model.Deck) error
	Update(*model.Deck) error
	Delete(id int) error
	Find(id int) (*model.Deck, error)
	FindAllByUserID(userID int) ([]model.Deck, error)
	BelongsToUser(deckID int, userID int) bool
}

type CardRepository interface {
	Create(*model.Card) error
	Update(*model.Card) error
	Delete(id int) error
	Find(id int) (*model.Card, error)
	FindAllByDeckID(int) ([]model.Card, error)
	BelongsToUser(cardID int, userID int) bool
}
