package model

type Card struct {
	ID    int    `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`

	DeckID int `json:"deck_id"`
}
