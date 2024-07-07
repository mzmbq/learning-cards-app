package dict_test

import (
	"testing"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/dict"
	"github.com/stretchr/testify/assert"
)

func TestWiktionary_Search(t *testing.T) {
	d := dict.Wiktionary{}
	hits, err := d.Search("hello")

	assert.Nil(t, err)
	assert.NotEmpty(t, hits, "no hits")
}

func TestWiktionary_Define(t *testing.T) {
	d := dict.Wiktionary{}

	defs, err := d.Define("hello")
	assert.Nil(t, err)
	assert.NotEmpty(t, defs, "no definitions")

	for d := range defs {
		assert.NotEmpty(t, defs[d].Definition, "empty definition")
	}

	defs, err = d.Define("nonexistentword123")
	assert.Error(t, err)
	assert.Empty(t, defs)
}
