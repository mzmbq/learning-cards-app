package model

type Deck struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=100"`

	UserID int `json:"user_id"`
}

func (d *Deck) Validate() error {
	return validate.Struct(d)
}
