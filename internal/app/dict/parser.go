package dict

import "fmt"

type Entry struct {
	Word       string
	Definition string
	Examples   []string
}

type Parser interface {
	Search(word string) ([]string, error)
	Define(word string) ([]Entry, error)
}

func New(dict string) (Parser, error) {
	switch dict {
	case "wiktionary":
		return &Wiktionary{}, nil
	default:
		return nil, fmt.Errorf("unknown dictionary: %s", dict)
	}
}
