package utils

import (
	"fmt"
	"math/rand"
)

func RandomColor() string {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)

	hexColor := fmt.Sprintf("#%02x%02x%02x", r, g, b)
	return hexColor
}
