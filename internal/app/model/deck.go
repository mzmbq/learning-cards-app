package model

import "github.com/mzmbq/learning-cards-app/backend/internal/app/utils"

type Deck struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=100"`

	Background string `json:"background" validate:"required,hexcolor"`

	UserID int `json:"user_id"`
}

func (d *Deck) Validate() error {
	return validate.Struct(d)
}

func (d *Deck) BeforeCreate() {
	if d.Background == "" {
		d.Background = utils.RandomColor()
	}
}
