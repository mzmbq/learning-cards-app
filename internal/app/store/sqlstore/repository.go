package sqlstore

import (
	"database/sql"
	"log"

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

func (r *DeckRepository) Delete(id int) error {
	// delete deck
	stmt, err := r.store.db.Prepare("DELETE FROM decks WHERE id = $1")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Println("rowsAffected not supported by driver")
	} else {
		log.Printf("repository: removed %d deck(s)\n", count)
	}

	// delete cards
	stmt, err = r.store.db.Prepare("DELETE FROM cards WHERE deck_id = $1")
	if err != nil {
		return err
	}
	res, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	count, err = res.RowsAffected()
	if err != nil {
		log.Println("rowsAffected not supported by driver")
	} else {
		log.Printf("repository: removed %d card(s)\n", count)
	}

	return nil
}

func (r *DeckRepository) Update(deck *model.Deck) error {
	res, err := r.store.db.Exec("UPDATE decks SET name = $1, user_id = $2 WHERE id = $3",
		deck.Name,
		deck.UserID,
		deck.ID,
	)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Println("rowsAffected not supported by driver")
		return err
	}

	if count == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}

func (r *DeckRepository) BelongsToUser(deckID int, userID int) bool {
	d, err := r.store.Deck().Find(deckID)
	if err != nil {
		return false
	}
	if d.UserID != userID {
		return false
	}
	return true
}

// Card

type CardRepository struct {
	store *Store
}

func (r *CardRepository) Create(c *model.Card) error {
	return r.store.db.QueryRow(
		"INSERT INTO cards (front, back, deck_id, ease, interval, state, step, due) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		c.Front,
		c.Back,
		c.DeckID,
		c.Flashcard.Ease,
		c.Flashcard.Interval,
		c.Flashcard.State,
		c.Flashcard.Step,
		c.Flashcard.Due,
	).Scan(&c.ID)
}

func (r *CardRepository) Find(id int) (*model.Card, error) {
	c := &model.Card{}
	err := r.store.db.QueryRow(
		"SELECT id, front, back, deck_id, ease, interval, state, step, due FROM cards WHERE id = $1",
		id,
	).Scan(
		&c.ID,
		&c.Front,
		&c.Back,
		&c.DeckID,
		&c.Flashcard.Ease,
		&c.Flashcard.Interval,
		&c.Flashcard.State,
		&c.Flashcard.Step,
		&c.Flashcard.Due,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return c, nil
}

func (r *CardRepository) FindAllByDeckID(deckID int) ([]model.Card, error) {
	rows, err := r.store.db.Query(
		`SELECT id, front, back, deck_id, ease, interval, state, step, due FROM cards WHERE deck_id = $1`,
		deckID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]model.Card, 0)
	for rows.Next() {
		c := model.Card{}
		err := rows.Scan(
			&c.ID,
			&c.Front,
			&c.Back,
			&c.DeckID,
			&c.Flashcard.Ease,
			&c.Flashcard.Interval,
			&c.Flashcard.State,
			&c.Flashcard.Step,
			&c.Flashcard.Due,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *CardRepository) Update(c *model.Card) error {
	_, err := r.store.db.Exec(
		"UPDATE cards SET front = $1, back = $2, deck_id = $3, ease = $4, interval = $5, state = $6, step = $7, due = $8 WHERE id = $9",
		c.Front,
		c.Back,
		c.DeckID,
		c.Flashcard.Ease,
		c.Flashcard.Interval,
		c.Flashcard.State,
		c.Flashcard.Step,
		c.Flashcard.Due,
		c.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *CardRepository) Delete(id int) error {
	err := r.store.db.QueryRow("DELETE FROM cards WHERE id = $1", id).Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}

	return nil
}

// ? use SQL query
func (r *CardRepository) BelongsToUser(cardID int, userID int) bool {
	decks, err := r.store.Deck().FindAllByUserID(userID)
	if err != nil || len(decks) == 0 {
		return false
	}

	log.Println(decks)

	card, err := r.store.Card().Find(cardID)
	if err != nil {
		return false
	}

	log.Println(card)

	for _, deck := range decks {
		if card.DeckID == deck.ID {
			return true
		}
	}

	return false
}
