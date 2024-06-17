package apiserver

func (s *server) routes() {
	s.mux.AddMiddleware(s.withCORS)
	s.mux.AddMiddleware(withLogging)

	s.mux.Handle("POST /api/user/create", s.handleUserCreate())
	s.mux.Handle("POST /api/user/auth", s.handleUserAuth())
	s.mux.Handle("GET /api/user/whoami", s.withAuth(s.handleUserWhoami()))
	s.mux.Handle("GET /api/user/signout", s.handlerUserSignOut())

	s.mux.Handle("GET /api/decks/list", s.withAuth(s.handleDecksList()))
	s.mux.Handle("POST /api/deck/create", s.withAuth(s.handleDeckCreate()))
	s.mux.Handle("GET /api/deck/delete/{id}", s.withAuth(s.handleDeckDelete()))
	s.mux.Handle("GET /api/deck/list-cards/{id}", s.withAuth(s.handleDeckListCards()))

	s.mux.Handle("POST /api/card/create", s.withAuth(s.handleCardCreate()))
	s.mux.Handle("POST /api/card/update/{id}", s.withAuth(s.handleCardUpdate()))
	s.mux.Handle("GET /api/card/delete/{id}", s.withAuth(s.handleCardDelete()))

	s.mux.Handle("GET /api/study/get-card/{deck_id}", s.withAuth(s.handleStudyGetCard()))
	s.mux.Handle("POST /api/study/submit/{card_id}", s.withAuth(s.handleStudySubmit()))
}
