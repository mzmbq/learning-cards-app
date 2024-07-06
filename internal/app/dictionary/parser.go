package dictionary

type Entry struct {
	Word       string
	Definition string
	Examples   []string
}

type Parser interface {
	Search(word string) (string, error)
	Define(word string) ([]Entry, error)
}
