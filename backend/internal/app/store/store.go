package store

type Store interface {
	User() UserRepository
	// Deck() DeckRepository
}
