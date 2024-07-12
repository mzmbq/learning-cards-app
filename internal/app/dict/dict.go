package dict

import (
	"errors"
	"fmt"
)

var (
	ErrTooManyRequests = errors.New("too many requests to the dictionary service")
)

type Entry struct {
	Word       string   `json:"word"`
	Definition string   `json:"definition"`
	Examples   []string `json:"examples"`
}

type Dict interface {
	Search(word string) ([]string, error)
	Define(word string) ([]Entry, error)
}

func New(dict string) (Dict, error) {
	switch dict {
	case "wiktionary":
		return &Wiktionary{}, nil
	default:
		return nil, fmt.Errorf("unknown dictionary: %s", dict)
	}
}
