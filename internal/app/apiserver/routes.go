package apiserver

func (s *server) routes() {
	s.mux.AddMiddleware(s.withCORS)
	s.mux.AddMiddleware(withLogging)

	s.mux.HandleFunc("GET /", handleRoot)

	s.mux.Handle("POST /api/user/create", s.handleUserCreate())
	s.mux.Handle("POST /api/user/auth", s.handleUserAuth())
	s.mux.Handle("GET /api/user/whoami", s.withAuth(s.handleUserWhoami()))
	s.mux.Handle("GET /api/user/signout", s.handlerUserSignOut())

	s.mux.Handle("GET /api/decks/list", s.withAuth(s.handleDecksList()))
	s.mux.Handle("POST /api/deck/create", s.withAuth(s.handleDeckCreate()))
	s.mux.Handle("GET /api/deck/delete/{id}", s.withAuth(s.handleDeckDelete()))

	s.mux.Handle("GET /api/deck/list/{id}", s.withAuth(s.handleDeckList()))

	// TODO: implement
	// s.mux.Handle("GET /api/user/{email}", s.handleUserFind())

	// s.mux.Handle("GET /api/word", s.handleWordDefine())

	// s.mux.Handle("GET /api/decks", s.handleDecksList())

	// s.mux.Handle("POST /api/card", s.handleCardCreate())
	// s.mux.Handle("GET /api/deck/{deckId}/card", s.handleCardLearn())
	// s.mux.Handle("POST /api/card/{id}", s.handleCardUpdate())

}
