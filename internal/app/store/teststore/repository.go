package teststore

import (
	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func NewUserRepo(s *Store) *UserRepository {
	return &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.Email] = u

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

type DeckRepository struct {
	store *Store
	decks map[int]*model.Deck
}

func NewDeckRepo(s *Store) *DeckRepository {
	return &DeckRepository{
		store: s,
		decks: make(map[int]*model.Deck),
	}
}

func (r *DeckRepository) Create(d *model.Deck) error {
	d.ID = len(r.decks) + 1
	r.decks[d.ID] = d
	return nil
}

func (r *DeckRepository) Update(d *model.Deck) error {
	r.decks[d.ID] = d
	return nil
}

func (r *DeckRepository) Delete(id int) error {
	delete(r.decks, id)

	cards, err := r.store.Card().FindAllByDeckID(id)
	if err != nil {
		return err
	}

	for _, c := range cards {
		if err := r.store.Card().Delete(c.ID); err != nil {
			return err
		}
	}

	return nil
}

func (r *DeckRepository) Find(id int) (*model.Deck, error) {
	if d, ok := r.decks[id]; ok {
		return d, nil
	}
	return nil, store.ErrRecordNotFound
}

func (r *DeckRepository) FindAllByUserID(userID int) ([]model.Deck, error) {
	decks := make([]model.Deck, 0)
	for _, d := range r.decks {
		if d.UserID == userID {
			decks = append(decks, *d)
		}
	}

	return decks, nil
}

func (r *DeckRepository) BelongsToUser(deckID int, userID int) bool {
	d, ok := r.decks[deckID]
	if !ok {
		return false
	}
	return d.UserID == userID
}

type CardRepository struct {
	store *Store
	cards map[int]*model.Card
}

func NewCardRepo(s *Store) *CardRepository {
	return &CardRepository{
		store: s,
		cards: make(map[int]*model.Card),
	}
}

func (r *CardRepository) Create(c *model.Card) error {
	c.ID = len(r.cards) + 1
	r.cards[c.ID] = c
	return nil
}

func (r *CardRepository) Update(c *model.Card) error {
	r.cards[c.ID] = c
	return nil
}

func (r *CardRepository) Delete(id int) error {
	if _, ok := r.cards[id]; !ok {
		return store.ErrRecordNotFound
	}
	delete(r.cards, id)
	return nil
}

func (r *CardRepository) Find(id int) (*model.Card, error) {
	if c, ok := r.cards[id]; ok {
		return c, nil
	}
	return nil, store.ErrRecordNotFound
}

func (r *CardRepository) FindAllByDeckID(deckID int) ([]model.Card, error) {
	cards := make([]model.Card, 0)
	for _, c := range r.cards {
		if c.DeckID == deckID {
			cards = append(cards, *c)
		}
	}

	return cards, nil
}

func (r *CardRepository) BelongsToUser(cardID int, userID int) bool {
	decks, err := r.store.Deck().FindAllByUserID(userID)
	if err != nil || len(decks) == 0 {
		return false
	}

	c, err := r.Find(cardID)
	if err != nil {
		return false
	}

	for _, d := range decks {
		if c.DeckID == d.ID {
			return true
		}
	}

	return false
}
