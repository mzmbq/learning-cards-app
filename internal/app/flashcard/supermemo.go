package flashcard

import (
	"fmt"
	"time"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
)

// Anki's legacy algorithm: https://faqs.ankiweb.net/what-spaced-repetition-algorithm.html
// based on SM-2: https://en.wikipedia.org/wiki/SuperMemo#Description_of_SM-2_algorithm

const (
	Again = iota
	Hard
	Good
	Easy
)

var stepDelays = map[int]time.Duration{
	0: 1 * time.Minute,
	1: 10 * time.Minute,
	2: 24 * time.Hour,
}

func SuperMemoAnki(fc *model.Flashcard, grade int) error {
	if fc.State != model.StateReview && fc.State != model.StateLearning {
		return fmt.Errorf("invalid flashcard state: %d", fc.State)
	}

	switch grade {
	case Again:
		answerAgain(fc)
	case Hard:
		answerHard(fc)
	case Good:
		answerGood(fc)
	case Easy:
		answerEasy(fc)
	default:
		return fmt.Errorf("unknown grade: %d", grade)
	}

	return nil
}

func answerAgain(fc *model.Flashcard) {
	switch fc.State {
	case model.StateLearning:
		fc.Step = 0
		fc.Due = time.Now().Add(stepDelays[fc.Step])

	case model.StateReview:
		toLearning(fc)
	}
}

func answerHard(fc *model.Flashcard) {
	switch fc.State {
	case model.StateLearning:
		fc.Due = time.Now().Add(stepDelays[fc.Step])

	case model.StateReview:
		fc.Ease -= 0.15
		if fc.Ease < 1.3 {
			fc.Ease = 1.3
		}
		fc.Interval = time.Duration(float64(fc.Interval) * 1.2)
	}
}

func answerGood(fc *model.Flashcard) {
	switch fc.State {
	case model.StateLearning:
		if fc.Step == 3 {
			toReview(fc)
		} else {
			fc.Step += 1
			fc.Due = time.Now().Add(stepDelays[fc.Step])
		}

	case model.StateReview:
		// Ease unchanged
		fc.Interval = time.Duration(float64(fc.Interval) * fc.Ease)

	}
}

func answerEasy(fc *model.Flashcard) {
	switch fc.State {
	case model.StateLearning:
		toReview(fc)

	case model.StateReview:
		fc.Ease += 0.15
		fc.Interval = time.Duration(float64(fc.Interval) * fc.Ease)
	}
}

func toReview(fc *model.Flashcard) {
	fc.State = model.StateReview
	fc.Ease = 2.5
	fc.Interval = 24 * time.Hour // ?
	fc.Due = time.Now().Add(fc.Interval)
}

func toLearning(fc *model.Flashcard) {
	fc.State = model.StateLearning
	fc.Step = 0
	fc.Due = time.Now().Add(stepDelays[fc.Step])
}

func New() *model.Flashcard {
	fc := &model.Flashcard{}
	toLearning(fc)
	return fc
}
