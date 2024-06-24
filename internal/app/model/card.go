package model

import "time"

type Card struct {
	ID     int    `json:"id"`
	Front  string `json:"front"`
	Back   string `json:"back"`
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
