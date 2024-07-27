package apiserver

func (s *server) routes() {
	s.Use(s.withLogging, s.withCORS, s.withGlobalRateLimit, s.withAuth, s.withUserRateLimit)

	s.HandleNoMiddleware("GET /api/health", s.withLogging(s.withCORS(s.withGlobalRateLimit(s.handleHealthcheck()))))

	s.HandleNoMiddleware("POST /api/user/create", s.withLogging(s.withCORS(s.handleUserCreate())))
	s.HandleNoMiddleware("POST /api/user/auth", s.withLogging(s.withCORS(s.handleUserAuth())))
	s.Handle("GET /api/user/whoami", s.handleUserWhoami())
	s.Handle("GET /api/user/signout", s.handlerUserSignOut())

	s.Handle("GET /api/decks/list", s.handleDecksList())
	s.Handle("POST /api/deck/create", s.handleDeckCreate())
	s.Handle("GET /api/deck/delete/{id}", s.handleDeckDelete())
	s.Handle("GET /api/deck/list-cards/{id}", s.handleDeckListCards())

	s.Handle("POST /api/card/create", s.handleCardCreate())
	s.Handle("POST /api/card/update/{id}", s.handleCardUpdate())
	s.Handle("GET /api/card/delete/{id}", s.handleCardDelete())

	s.Handle("GET /api/study/get-card/{deck_id}", s.handleStudyGetCard())
	s.Handle("POST /api/study/submit/{card_id}", s.handleStudySubmit())

	s.Handle("GET /api/define/{dict}/{word}", s.handleDefine())
	s.Handle("GET /api/search/{dict}/{word}", s.handleSearch())
}
