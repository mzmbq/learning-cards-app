package model

import "time"

type Card struct {
	ID     int    `json:"id"`
	Front  string `json:"front" validate:"required,min=3,max=1000"`
	Back   string `json:"back" validate:"required,min=3,max=1000"`
	DeckID int    `json:"deck_id"`

	Flashcard Flashcard `json:"flashcard"`
}

type Flashcard struct {
	Ease float64 `json:"ease"`
	// Reps     int           `json:"reps"`
	Interval time.Duration `json:"interval"`
	State    int           `json:"state"`
	Step     int           `json:"step"` // current learning step
	Due      time.Time     `json:"due"`
}

// Flashcard state
const (
	StateLearning = iota
	StateReview
)
