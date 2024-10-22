package utils

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestRandomColor(t *testing.T) {
	for range 100 {
		color := RandomColor()
		assert.NotEmpty(t, color)

		validate := validator.New(validator.WithRequiredStructEnabled())
		err := validate.Var(color, "hexcolor")
		assert.Nil(t, err)
	}
}
