package sqlstore

import (
	"database/sql"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

// User

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	u.BeforeCreate()

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

//Deck

type DeckRepository struct {
	store *Store
}

func (r *DeckRepository) Create(d *model.Deck) error {
	return r.store.db.QueryRow(
		"INSERT INTO decks (name, user_id) VALUES ($1, $2) RETURNING id",
		d.Name,
		d.UserID,
	).Scan(&d.ID)
}

func (r *DeckRepository) Find(id int) (*model.Deck, error) {
	d := &model.Deck{}
	err := r.store.db.QueryRow(
		"SELECT id, name, user_id FROM decks WHERE id = $1",
		id,
	).Scan(
		&d.ID,
		&d.Name,
		&d.UserID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return d, nil
}

func (r *DeckRepository) FindAllByUserID(id int) ([]model.Deck, error) {
	rows, err := r.store.db.Query("SELECT id, name, user_id FROM decks WHERE user_id = $1",
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	decks := make([]model.Deck, 0)
	for rows.Next() {
		d := model.Deck{}
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.UserID,
		)
		if err != nil {
			return nil, err
		}
		decks = append(decks, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return decks, nil
}

// Card

type CardRepository struct {
	store *Store
}

func (r *CardRepository) Create(c *model.Card) error {
	return r.store.db.QueryRow(
		"INSERT INTO cards (front, back, deck_id) VALUES ($1, $2, $3) RETURNING id",
		c.Front,
		c.Back,
		c.DeckID,
	).Scan(&c.ID)
}

func (r *CardRepository) Find(id int) (*model.Card, error) {
	c := &model.Card{}
	err := r.store.db.QueryRow(
		"SELECT id, front, back, deck_id FROM cards WHERE id = $1",
		id,
	).Scan(
		&c.ID,
		&c.Front,
		&c.Back,
		&c.DeckID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return c, nil
}
